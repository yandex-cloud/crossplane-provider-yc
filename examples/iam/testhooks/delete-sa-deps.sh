#!/usr/bin/env bash
set -aeuo pipefail

# Delete storage bucket controlled by this SA.
${KUBECTL} delete buckets.storage.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=storage
# Delete folder IAM binding depending on this SA.
${KUBECTL} delete folderiambindings.iam.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=storage
# Delete static key depending on this SA.
${KUBECTL} delete serviceaccountstaticaccesskeys.iam.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=storage