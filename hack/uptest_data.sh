#!/usr/bin/env bash
set -aeuo pipefail

mkdir -p _output/tests

cat <<EOF | sed -e "s@<<FOLDER_ID>>@${FOLDER_ID}@g" -e "s@<<CLOUD_ID>>@${CLOUD_ID}@g" -e "s@<<RANDOM_ID>>@$RANDOM@g" >_output/tests/data.yaml
folderId: <<FOLDER_ID>>
cloudId: <<CLOUD_ID>>
randomId: <<RANDOM_ID>>
ip1: 10.0.0.40
ip2: 10.0.0.41
EOF


echo _output/tests/data.yaml