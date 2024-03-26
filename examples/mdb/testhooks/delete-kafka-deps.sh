#!/usr/bin/env bash
set -aeuo pipefail

# Delete topic before cluster
${KUBECTL} delete kafkatopic.mdb.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=kafka
# Delete the useringroup resource before the user pool
${KUBECTL} delete kafkauser.mdb.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=kafka