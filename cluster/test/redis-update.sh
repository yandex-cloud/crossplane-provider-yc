#!/usr/bin/env bash
set -euo pipefail

# UpTest 2.2's update template interpolates unescaped JSON into a quoted shell
# command. Exercise the update in a post-create hook until that template is fixed.
resource=${UPTEST_REDIS_RESOURCE:?Set by make uptest-redis}
timeout=${UPTEST_DEFAULT_TIMEOUT:-3600s}

echo "Testing post-create Redis update: ${resource}"
# Also require the initial value, so rerunning against an existing fixture
# cannot turn the update into a no-op.
"${KUBECTL}" --namespace default wait "${resource}" \
  --for='jsonpath={.status.atProvider.description}=Redis update regression' \
  --timeout="${timeout}"
"${KUBECTL}" --namespace default patch "${resource}" --type=merge \
  -p '{"spec":{"forProvider":{"description":"Redis update regression passed"}}}'

# Checking the new observed value prevents stale pre-update conditions from
# passing the regression without an actual cloud update.
"${KUBECTL}" --namespace default wait "${resource}" \
  --for='jsonpath={.status.atProvider.description}=Redis update regression passed' \
  --timeout="${timeout}"
for condition in Ready Synced LastAsyncOperation Test; do
  "${KUBECTL}" --namespace default wait "${resource}" \
    --for="condition=${condition}=True" --timeout="${timeout}"
done
echo "Post-create Redis update passed: ${resource}"
