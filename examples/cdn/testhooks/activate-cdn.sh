#!/usr/bin/env bash
set -aeuo pipefail

yc provider activate --folder-id ${FOLDER_ID} --type gcore