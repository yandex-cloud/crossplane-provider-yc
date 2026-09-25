# Redis lifecycle regression

`make uptest-redis` runs create, update, import, and delete against a running
provider for both cluster-scoped and namespaced Redis resources. Supply the same
credentials, folder, cloud, and kubeconfig as for `make uptest`. Both `make e2e`
and `make e2e-cloud` run this target after the general example suite.

Each Redis example is a separate UpTest run. UpTest 2.2's native update step
breaks JSON quoting before sending the patch. These examples instead execute
`cluster/test/redis-update.sh` as a post-assert hook after creation. The target
sets `UPTEST_REDIS_RESOURCE` to the resource under test. The native update phase
is skipped; the hook performs the real update and assertions before import.
Keep these cases separate from the general list, which excludes MDB resources.

The hook changes the description after creation, waits for the new value in
`status.atProvider`, then waits for Ready, Synced, LastAsyncOperation and Test.
This exercises the asynchronous Terraform SDK update after the original
reconcile context has ended (CLOUDMPDEV-5975). Keep the password in
`forProvider` and retain the default management policies, including late
initialization.

The hook explicitly selects the namespaced example's `default` namespace.
Each example uses one host; sharding is not required to trigger this regression.
