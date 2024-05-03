#!/usr/bin/env bash
set -aeuo pipefail

# Delete topic before cluster
${KUBECTL} delete mongodbuser.mdb.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=mongo
# Delete the useringroup resource before the user pool
${KUBECTL} delete mongodbdatabase.mdb.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=mongo