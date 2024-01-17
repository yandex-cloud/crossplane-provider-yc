#!/usr/bin/env bash
set -aeuo pipefail

cat <<EOF | sed -e "s@<<FOLDER_ID>>@${FOLDER_ID}@g" >_output/tests/data.yaml
folderId: <<FOLDER_ID>>
EOF

echo _output/tests/data.yaml