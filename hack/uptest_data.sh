#!/usr/bin/env bash
set -aeuo pipefail

mkdir -p _output/tests
cat <<EOF | sed -e "s@<<FOLDER_ID>>@${FOLDER_ID}@g" -e "s@<<CLOUD_ID>>@${CLOUD_ID}@g" >_output/tests/data.yaml
folderId: <<FOLDER_ID>>
cloudId: <<CLOUD_ID>>
EOF

echo _output/tests/data.yaml