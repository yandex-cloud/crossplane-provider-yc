#!/usr/bin/env bash
set -aeuo pipefail

function join {
  local f=${1-}
  if shift; then
    printf %s "$f" "${@/#/,}"
  fi
}

## EXCLUDED ON PURPOSE:
# providerconfig gets created on initializing tests,
# testing folder requires cloud admin privileges, not feasible
# container/repository needs registry ID explicitly provided
# datatransfer/transfer actually runs upon creation

## CURRENTLY FAILING:
# securitygroup and securitygrouprule can be created, but not really altered (yet?), so don't pass tests
# iam/serviceaccountkey has a bug: https://nda.ya.ru/t/8HkjK_y074Vmc4
# storage/object needs investigation
# storage/bucket does get removed, but Crossplane receives Forbidden; needs investigation
# securitygroup and securitygrouprule can be created, but not really altered (yet?), so don't pass tests
all=$(find ${1} -name "*.yaml" \
-not -path "*/alb/*" \
-not -path "*/container/repository.yaml" \
-not -path "*/datatransfer/transfer.yaml" \
-not -path "*/dns/*" \
-not -path "*/kms/*" \
-not -path "*/kubernetes/*" \
-not -path "*/mdb/*" \
-not -path "*/message/*" \
-not -path "*/vpc/securitygroup*.yaml" \
-not -path "*/iam/serviceaccountkey.yaml" \
-not -path "*/storage/object.yaml" \
-not -path "*/storage/bucket.yaml" \
-not -path "*/providerconfig/providerconfig.yaml" \
-not -path "*/resourcemanager/folder.yaml")

join $all

