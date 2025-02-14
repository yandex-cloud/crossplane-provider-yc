#!/usr/bin/env bash
set -aeuo pipefail

${KUBECTL} delete instancegroup.compute.yandex-cloud.jet.crossplane.io -l testing.upbound.io/example-name=ig
sleep 10
