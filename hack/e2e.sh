#!/bin/bash

[ -n "${OAUTH_TOKEN}" ] || { echo OAUTH_TOKEN env var not set, can not proceed; exit 1; }
[ -n "${SECRET_ID}" ] || { echo SECRET_ID env var not set, can not proceed; exit 1; }
[ -n "${CLOUD_ID}" ] || { echo CLOUD_ID env var not set, can not proceed; exit 1; }
[ -n "${FOLDER_ID}" ] || { echo FOLDER_ID env var not set, can not proceed; exit 1; }
[ -n "${NETWORK_ID}" ] || { echo NETWORK_ID env var not set, can not proceed; exit 1; }
[ -n "${SUBNET_ID}" ] || { echo SUBNET_ID env var not set, can not proceed; exit 1; }
[ -n "${CLUSTER_IP}" ] || { echo CLUSTER_IP env var not set, can not proceed; exit 1; }


echo "##teamcity[blockOpened name='keys' description='set up YC keys']"
yc config profile create robot
yc config set token ${OAUTH_TOKEN}

yc lockbox payload get --id ${SECRET_ID} --key key > key.json
export SA_ID=$(jq -r .service_account_id key.json)

yc config profile create sa-profile
yc config set service-account-key key.json
yc config set folder-id ${FOLDER_ID}
yc config set cloud-id ${CLOUD_ID}

echo "##teamcity[blockClosed name='keys']"

WORKDIR=${DOCKER_WORKDIR:-"$(cd .. && pwd)"}
git config --global --add safe.directory ${WORKDIR}

echo "##teamcity[blockOpened name='cleanup' description='clean up test folder']"
./hack/folder_cleanup.sh || exit 1
echo "##teamcity[blockClosed name='cleanup']"

echo "##teamcity[blockOpened name='provision' description='set up cluster and CR']"
./hack/provision_e2e.sh
echo "##teamcity[blockClosed name='provision']"

echo "##teamcity[blockOpened name='up' description='a temporaty measure to mitigate https://github.com/upbound/up/issues/416']"
yc lockbox payload get ${SECRET_ID} --key access-key >> awskey
mkdir ~/.aws && echo [default] > ~/.aws/credentials && echo '  'aws_access_key_id = $(jq -r .access_key.key_id awskey) >> ~/.aws/credentials && echo '  'aws_secret_access_key = $(jq -r .secret awskey) >> ~/.aws/credentials
mkdir -p .cache/tools/linux_x86_64
aws s3 --region=ru-central1 --endpoint-url=https://storage.yandexcloud.net cp s3://patched-for-temp-use/up .cache/tools/linux_x86_64/up-v0.21.0
chmod +x .cache/tools/linux_x86_64/up-v0.21.0
echo "##teamcity[blockClosed name='up']"

export KUBECONFIG=kubeconfig
echo "##teamcity[blockOpened name='make e2e-cloud']"
make e2e-cloud
exitcode=$?
echo "##teamcity[blockClosed name='make e2e-cloud']"

echo "##teamcity[blockOpened name='dump' description='dump cluster info']"
make controlplane.dump
echo "##teamcity[blockClosed name='dump']"

if [ $exitcode = 0 ]; then
  echo "##teamcity[blockOpened name='cleanup' description='clean up test folder']"
  ./hack/folder_cleanup.sh
  echo "##teamcity[blockClosed name='cleanup']"
fi

exit $exitcode
