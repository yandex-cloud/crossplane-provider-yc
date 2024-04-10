#!/usr/bin/env bash
set -aeuo pipefail

function join {
  local f=${1-}
  if shift; then
    printf %s "$f" "${@/#/,}"
  fi
}

## EXCLUDED ON PURPOSE:
# alb/targetgroup needs instance IPs explicitly provided
# container/repository needs registry ID explicitly provided
# message/queue works, but cleaning it in case of tests failure is a problem
all=$(find ${1} -name "*.yaml" \
-not -path "*/alb/*" \
-not -path "*/container/repository.yaml" \
-not -path "*/message/queue.yaml" )

join $all

