#!/bin/bash

echo "Provisioning e2e infrastructure..."

echo "Creating cluster..."
yc managed-kubernetes cluster create \
\--name e2e-controlplane --network-id ${NETWORK_ID} \
\--dual-stack \
\--cluster-ipv4-range "10.120.0.0/16" \
\--cluster-ipv6-range "fc0e::/112" \
\--service-ipv4-range "10.121.0.0/16" \
\--service-ipv6-range "fc0f::/112" \
\--service-account-id ${SA_ID} --node-service-account-id ${SA_ID} \
\--regional \
\--public-ipv6 ${CLUSTER_IP}

echo "Creating node group..."
yc managed-kubernetes node-group create --name e2e-ng \
\--cluster-name e2e-controlplane --platform-id standard-v2 \
\--cores 2 --memory 4 --core-fraction 50 \
\--fixed-size 1 \
\--location zone=ru-central1-c \
\--network-interface "ipv4-address=nat,ipv6-address=auto,subnets=${SUBNET_ID}"

echo "Exporting credentials..."
yc managed-kubernetes cluster get-credentials --external-ipv6 --name e2e-controlplane --kubeconfig kubeconfig --force

echo "Creating CR..."
yc container registry list | grep crossplane-e2e-cr || yc container registry create crossplane-e2e-cr
yc container registry add-access-binding crossplane-e2e-cr --subject system:allUsers --role viewer
yc container registry configure-docker

echo "Done!"
