#!/bin/bash

[ -n "${SA_ID}" ] || { echo SA_ID env var not set, can not proceed; exit 1; }
[ -n "${CLOUD_ID}" ] || { echo CLOUD_ID env var not set, can not proceed; exit 1; }
[ -n "${FOLDER_ID}" ] || { echo FOLDER_ID env var not set, can not proceed; exit 1; }
[ -n "${NETWORK_ID}" ] || { echo NETWORK_ID env var not set, can not proceed; exit 1; }
[ -n "${SUBNET_ID}" ] || { echo SUBNET_ID env var not set, can not proceed; exit 1; }
[ -n "${CLUSTER_IP}" ] || { echo CLUSTER_IP env var not set, can not proceed; exit 1; }


echo "##teamcity[blockOpened name='keys' description='set up YC keys']"
if [[ -n "${E2E_TEST_OAUTH_TOKEN}" ]]; then
  yc config set token ${E2E_TEST_OAUTH_TOKEN}
fi

yc iam key create --service-account-id ${SA_ID} --output key.json
yc iam access-key create --service-account-id ${SA_ID} --format json > awskey
echo "##teamcity[blockClosed name='keys']"

mkdir -p .cache/tools/linux_x86_64

ROOT="$(cd .. && pwd)"

DOCKER_CONF_VOLUME=
DOCKER_CRED_HELPER_VOLUME_PARAM=
SSH_AUTH_SOCK_PARAMS=()
OS_NAME="$(uname -s)"
CONTAINER_HOME="/root"
DOCKER_PARAMS=${DOCKER_PARAMS:-}
DOCKER_WORKDIR=${DOCKER_WORKDIR:-"$(cd .. && pwd)"}
AW_TOOLS_COMMON_VERSION=${AW_TOOLS_COMMON_VERSION:-"5"}

if [[ -z "${TEAMCITY_VERSION}" ]]; then
  # enable tty
  DOCKER_PARAMS="${DOCKER_PARAMS} -t"
else
  # get functional git repo inside docker container
  DOCKER_PARAMS="${DOCKER_PARAMS} --volume /opt/buildagent/system/git:/opt/buildagent/system/git"
fi

if which go; then
  DOCKER_PARAMS="${DOCKER_PARAMS} --volume $(go env GOCACHE):/root/.cache/go-build  --volume $(go env GOPATH)/pkg:/gopath/pkg"
else
  DOCKER_PARAMS="${DOCKER_PARAMS} --volume /tmp/go-build-cache:/root/.cache/go-build  --volume /tmp/go-pkg-cache:/gopath/pkg"
fi

case "${OS_NAME}" in
Linux*)
  if [[ -z "${DOCKER_CONF}" ]]; then
    DOCKER_CONF_VOLUME="${HOME}/.docker:${CONTAINER_HOME}/.docker"
  else
    DOCKER_CONF_VOLUME="${DOCKER_CONF}:${CONTAINER_HOME}/.docker/config.json"
  fi
  DOCKER_CRED_HELPER_VOLUME_PARAM="--volume /usr/local/bin/docker-credential-ycr:/usr/local/bin/docker-credential-ycr:ro"

  # Our envoy uses mlock() on linux.
  DOCKER_PARAMS="${DOCKER_PARAMS} --cap-add CAP_IPC_LOCK"

  # get full path to docker socket
  # linux: /var/run/docker.sock -> /run/docker.sock
  DOCKERD_SOCK=$(readlink -f /var/run/docker.sock)

  if [[ -n "${SSH_AUTH_SOCK}" ]]; then
    SSH_AUTH_SOCK_PARAMS=(-v $(dirname "${SSH_AUTH_SOCK}"):$(dirname "${SSH_AUTH_SOCK}") -e "SSH_AUTH_SOCK=${SSH_AUTH_SOCK}")
  fi
  ;;
Darwin*)
  if [[ -z "${DOCKER_CONF}" ]]; then
    # Mac OS X docker creds most likely are stored in Keychain and cannot be
    # used inside Linux VM. Thus we use some static creds unless user provides
    # explicit config file to mount.
    DOCKER_CONF_VOLUME="${ROOT}/build/ci/default-docker-auth.json:${CONTAINER_HOME}/.docker/config.json"
  else
    DOCKER_CONF_VOLUME="${DOCKER_CONF}:${CONTAINER_HOME}/.docker/config.json"
  fi

  if [[ "$OSTYPE" == "darwin"* ]]; then
    # https://github.com/docker/for-mac/issues/4755#issuecomment-726351209
    DOCKERD_SOCK="/var/run/docker.sock.raw"
  else
    DOCKERD_SOCK="/var/run/docker.sock"
  fi

  if [[ -n "${SSH_AUTH_SOCK}" ]]; then
      # https://github.com/docker/for-mac/issues/410#issuecomment-577064671
      # https://github.com/docker/for-mac/issues/410#issuecomment-557613306
      # "volume parameter needs to be exactly as above, i.e. -v /run/host-services/ssh-auth.sock:/run/host-services/ssh-auth.sock
      # You're not actually pointing at a socket on your host system, you're pointing at some other special machinery
      # in Docker Desktop for Mac that enables specifically the SSH agent forwarding"
      SSH_AUTH_SOCK_PARAMS=(--env SSH_AUTH_SOCK=/run/host-services/ssh-auth.sock -v /run/host-services/ssh-auth.sock:/run/host-services/ssh-auth.sock)
  fi
  ;;
MINGW*)
  if [[ -z "${DOCKER_CONF}" ]]; then
    # Windows docker creds most likely are stored in Keychain and cannot be
    # used inside Linux VM. Thus we use some static creds unless user provides
    # explicit config file to mount.
    DOCKER_CONF_VOLUME="${ROOT}/build/ci/default-docker-auth.json:${CONTAINER_HOME}/.docker/config.json"
  else
    DOCKER_CONF_VOLUME="${DOCKER_CONF}:${CONTAINER_HOME}/.docker/config.json"
  fi
  DOCKERD_SOCK="//var/run/docker.sock"

  if [[ -n "${SSH_AUTH_SOCK}" ]]; then
    SSH_AUTH_SOCK_PARAMS=(-v $(dirname "${SSH_AUTH_SOCK}"):$(dirname "${SSH_AUTH_SOCK}") -e "SSH_AUTH_SOCK=${SSH_AUTH_SOCK}")
  fi
  ;;
*)
  echo "Unsupported OS platform: ", "${OS_NAME}"
  exit 1
  ;;
esac

export DOCKER_CLI_EXPERIMENTAL=enabled


env |
  cut -f1 -d= |
  grep -v -e '^LANG$' -e '^LANGUAGE$' -e '^LC_' -e '^PATH$' -e '^HOME$' -e '^HOSTNAME' -e '^DOCKER_PARAMS$' -e '^PWD$' -e '^SHELL$' -e '^TMPDIR$' -e '^GOBIN$' -e '^GOPATH$' -e ' ' |
  sort > ./env.list


AWS_KEY_ID='$(jq -r .access_key.key_id awskey)'
AWS_SECRET='$(jq -r .secret awskey)'
EXITCODE_LITERAL='$?'
EXITCODE='$exitcode'

cat <<EOF > e2e.sh
#!/bin/bash

git config --global --add safe.directory ${DOCKER_WORKDIR}

echo "##teamcity[blockOpened name='creds' description='set up YC credentials']"
yc config profile create sa-profile
yc config set service-account-key key.json
yc config set folder-id ${FOLDER_ID}
yc config set cloud-id ${CLOUD_ID}
echo "##teamcity[blockClosed name='creds']"

echo "##teamcity[blockOpened name='cleanup' description='clean up test folder']"
./hack/folder_cleanup.sh
echo "##teamcity[blockClosed name='cleanup']"

echo "##teamcity[blockOpened name='provision' description='set up cluster and CR']"
./hack/provision_e2e.sh
echo "##teamcity[blockClosed name='provision']"


echo "##teamcity[blockOpened name='up' description='fetch patched up']"
# a temporaty measure to mitigate https://github.com/upbound/up/issues/416
mkdir ~/.aws && echo [default] > ~/.aws/credentials && echo '  'aws_access_key_id = $AWS_KEY_ID >> ~/.aws/credentials && echo '  'aws_secret_access_key = $AWS_SECRET >> ~/.aws/credentials
aws s3 --region=ru-central1 --endpoint-url=https://storage.yandexcloud.net cp s3://patched-for-temp-use/up .cache/tools/linux_x86_64/up-v0.21.0
chmod +x .cache/tools/linux_x86_64/up-v0.21.0
echo "##teamcity[blockClosed name='up']"

export KUBECONFIG=kubeconfig
make e2e-cloud
exitcode=$EXITCODE_LITERAL


echo "##teamcity[blockOpened name='dump' description='dump cluster info']"
make controlplane.dump
echo "##teamcity[blockClosed name='up']"

yc iam access-key delete $AWS_KEY_ID
if [ $EXITCODE ]; then
  echo "##teamcity[blockOpened name='cleanup' description='clean up test folder']"
  ./hack/folder_cleanup.sh
  echo "##teamcity[blockClosed name='cleanup']"
fi

exit $EXITCODE
EOF
chmod +x e2e.sh

docker run --rm -i ${DOCKER_PARAMS} --env-file env.list \
  --cap-add SYS_PTRACE \
  "${SSH_AUTH_SOCK_PARAMS[@]}" \
  ${DOCKER_CRED_HELPER_VOLUME_PARAM} \
  --volume "${DOCKER_CONF_VOLUME}" \
  --volume "${DOCKERD_SOCK}:/var/run/docker.sock" \
  --volume "${ROOT}":"${DOCKER_WORKDIR}" \
  --workdir "${DOCKER_WORKDIR}/provider-jet-yc" \
  --net=host \
  cr.yandex/yc-internal/aw-tools-common:${AW_TOOLS_COMMON_VERSION} ./e2e.sh
