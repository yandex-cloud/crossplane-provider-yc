#!/usr/bin/env bash
set -aeuo pipefail

# Delete the database before the user
${KUBECTL} delete postgresqldatabase.mdb.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=postgres
# Delete user before cluster
${KUBECTL} delete postgresqluser.mdb.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=postgres