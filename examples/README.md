# Crossplane Provider YC Examples

This directory contains example manifests for using the Yandex Cloud Crossplane Provider.

## Crossplane v2 Support

Starting with provider version v0.14.0, this provider supports both Crossplane v1 (cluster-scoped) and Crossplane v2 (namespaced) resources.

## Example Types

### 1. Cluster-Scoped Resources (Legacy - Crossplane v1)

These examples use the original API group `yandex-cloud.jet.crossplane.io` and are cluster-scoped.

**ProviderConfig:**
- `providerconfig/providerconfig.yaml` - Basic cluster-scoped ProviderConfig
- `providerconfig/providerconfig-separate-secrets.yaml` - Using separate secrets
- `providerconfig/providerconfig-mixed-secret.yaml` - Mixed authentication
- `providerconfig/providerconfig-token-secret.yaml` - Token-based authentication

**Managed Resources:**
- All examples in subdirectories (e.g., `storage/bucket.yaml`, `compute/instance.yaml`) are cluster-scoped

### 2. Namespaced Resources (Modern - Crossplane v2)

These examples use the new API group `yandex-cloud.m.jet.crossplane.io` and are namespace-scoped.

**ProviderConfig:**
- `providerconfig/namespaced-providerconfig.yaml` - Namespace-scoped ProviderConfig
- `providerconfig/cluster-providerconfig.yaml` - Cluster-scoped ProviderConfig in new API group

**Managed Resources:**
- `storage/namespaced-bucket.yaml` - Example namespaced MR

## Key Differences

### Cluster-Scoped (v1)
```yaml
apiVersion: storage.yandex-cloud.jet.crossplane.io/v1alpha1
kind: Bucket
metadata:
  name: my-bucket
spec:
  providerConfigRef:
    name: default
  forProvider:
    secretKeySecretRef:
      name: my-secret
      namespace: crossplane-system  # namespace required
      key: secret-key
```

### Namespaced (v2)
```yaml
apiVersion: storage.yandex-cloud.m.jet.crossplane.io/v1alpha1
kind: Bucket
metadata:
  name: my-bucket
  namespace: default  # namespace required
spec:
  providerConfigRef:
    kind: ClusterProviderConfig  # kind required
    name: default
  forProvider:
    secretKeySecretRef:
      name: my-secret  # local to MR namespace
      key: secret-key
```

## ProviderConfig Reference Types

Namespaced MRs can reference either:

1. **ClusterProviderConfig** (cluster-scoped, new API group):
   ```yaml
   providerConfigRef:
     kind: ClusterProviderConfig
     name: default
   ```

2. **ProviderConfig** (namespace-scoped, new API group):
   ```yaml
   providerConfigRef:
     kind: ProviderConfig
     name: default
   ```

If `providerConfigRef` is omitted, it defaults to `kind: ClusterProviderConfig, name: default`.

## Migration Guide

To migrate from cluster-scoped to namespaced resources:

1. Update the API group from `yandex-cloud.jet.crossplane.io` to `yandex-cloud.m.jet.crossplane.io`
2. Add `metadata.namespace` to your MRs
3. Add `spec.providerConfigRef.kind` (ClusterProviderConfig or ProviderConfig)
4. Remove `namespace` field from all secret references in `spec.forProvider`
5. Remove `namespace` field from `spec.writeConnectionSecretToRef`

## Backward Compatibility

The provider maintains full backward compatibility:
- Existing cluster-scoped MRs continue to work without changes
- Both cluster-scoped and namespaced MRs can coexist
- Cluster-scoped MRs only reference legacy ProviderConfig
- Namespaced MRs only reference modern ProviderConfig types