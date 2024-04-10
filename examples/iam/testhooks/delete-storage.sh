#!/usr/bin/env bash
set -aeuo pipefail

# Delete storage bucket so that is does not get sync failure with 403.
${KUBECTL} delete buckets.storage.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=storage
