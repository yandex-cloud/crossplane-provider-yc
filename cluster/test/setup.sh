#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
echo "Creating cloud credential secret..."
${KUBECTL} -n upbound-system create secret generic provider-secret --from-literal=credentials="${CREDENTIALS}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating a default provider config..."
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
      namespace: upbound-system
      key: credentials
EOF
