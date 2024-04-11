#!/usr/bin/env bash
set -aeuo pipefail

function join {
  local f=${1-}
  if shift; then
    printf %s "$f" "${@/#/,}"
  fi
}

## EXCLUDED ON PURPOSE:
# message/queue works, but cleaning it in case of tests failure is a problem
all=$(find ${1} -name "*.yaml" \
-not -path "*/message/queue.yaml" )

join $all

