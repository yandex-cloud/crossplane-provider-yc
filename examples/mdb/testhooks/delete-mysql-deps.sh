#!/usr/bin/env bash
set -aeuo pipefail

# Delete user before cluster
${KUBECTL} delete mysqluser.mdb.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=mysql
# Delete the database before the user
${KUBECTL} delete mysqldatabase.mdb.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=mysql