#!/bin/bash

[ -n "${SECRET_ID}" ] || { echo SECRET_ID env var not set, can not proceed; exit 1; }
[ -n "${NETWORK_ID}" ] || { echo NETWORK_ID env var not set, can not proceed; exit 1; }
[ -n "${SUBNET_ID}" ] || { echo SUBNET_ID env var not set, can not proceed; exit 1; }
[ -n "${CLUSTER_IP}" ] || { echo CLUSTER_IP env var not set, can not proceed; exit 1; }


echo "##teamcity[blockOpened name='keys' description='set up YC keys']"

#yc config profile create robot
#yc config set token ${OAUTH_TOKEN}
#
#yc lockbox payload get --id ${SECRET_ID} --key key > key.json
#
#yc config profile create sa-profile
#yc config set service-account-key key.json
#yc config set folder-id ${FOLDER_ID}
#yc config set cloud-id ${CLOUD_ID}

export SA_ID=$(jq -r .service_account_id ${SA_KEY_FILE})
yc lockbox payload get --id ${SECRET_ID} --key access-key >> awskey
mkdir ~/.aws && echo [default] > ~/.aws/credentials && echo '  'aws_access_key_id = $(jq -r .access_key.key_id awskey) >> ~/.aws/credentials && echo '  'aws_secret_access_key = $(jq -r .secret awskey) >> ~/.aws/credentials
echo "##teamcity[blockClosed name='keys']"

#WORKDIR=${DOCKER_WORKDIR:-"$(cd .. && pwd)"}
#git config --global --add safe.directory ${WORKDIR}

echo "##teamcity[blockOpened name='cleanup' description='clean up test folder']"
if ! ./hack/folder_cleanup.sh; then
  echo "##teamcity[buildStatus text='Failed to clean up test folder']"
  exit 1
fi
echo "##teamcity[blockClosed name='cleanup']"

echo "##teamcity[blockOpened name='certificate' description='create Certificate Manager IAM prerequisite']"
certificate_dir=$(mktemp -d)
trap 'rm -f "$certificate_dir/certificate.pem" "$certificate_dir/private-key.pem"; rmdir "$certificate_dir"' EXIT
openssl req -x509 -nodes -newkey rsa:2048 -days 1 \
  -subj "/CN=crossplane-provider-yc-e2e.invalid" \
  -keyout "$certificate_dir/private-key.pem" \
  -out "$certificate_dir/certificate.pem" >/dev/null 2>&1
CERTIFICATE_ID=$(yc certificate-manager certificate create \
  --folder-id "${FOLDER_ID}" \
  --name "crossplane-provider-yc-e2e-${RANDOM}" \
  --chain "$certificate_dir/certificate.pem" \
  --key "$certificate_dir/private-key.pem" \
  --format json | jq -r .id)
[ -n "${CERTIFICATE_ID}" ] || { echo Failed to create Certificate Manager prerequisite; exit 1; }
export CERTIFICATE_ID
echo "##teamcity[blockClosed name='certificate']"

echo "##teamcity[blockOpened name='provision' description='set up cluster and CR']"
./hack/provision_e2e.sh
echo "##teamcity[blockClosed name='provision']"

export KUBECONFIG=kubeconfig
export DOCKER_CLI_EXPERIMENTAL=enabled
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
