#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
echo "Creating cloud credential secret..."
${KUBECTL} -n crossplane-system create secret generic provider-secret --from-literal=credentials="${CREDENTIALS}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating a default legacy provider config for cluster-scoped MRs..."
cat <<EOF | sed -e "s@<<CLOUD_ID>>@${CLOUD_ID}@g" -e "s@<<FOLDER_ID>>@${FOLDER_ID}@g" | ${KUBECTL} apply -f -
apiVersion: yandex-cloud.jet.crossplane.io/v1alpha1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    cloudId: <<CLOUD_ID>>
    folderId: <<FOLDER_ID>>
    source: Secret
    secretRef:
      name: provider-secret
      namespace: crossplane-system
      key: credentials
EOF

echo "Creating a default cluster provider config for namespaced MRs..."
cat <<EOF | sed -e "s@<<CLOUD_ID>>@${CLOUD_ID}@g" -e "s@<<FOLDER_ID>>@${FOLDER_ID}@g" | ${KUBECTL} apply -f -
apiVersion: yandex-cloud.m.jet.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    cloudId: <<CLOUD_ID>>
    folderId: <<FOLDER_ID>>
    source: Secret
    secretRef:
      name: provider-secret
      namespace: crossplane-system
      key: credentials
EOF

${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m
${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m

for crd in $(${KUBECTL} get crds -o jsonpath='{.items[*].metadata.name}'); do
  ${KUBECTL} wait --for=condition=established --timeout=5m "crd/$crd"
done
