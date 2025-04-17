#!/bin/bash

CLUSTER_IPV4_RANGE=${CLUSTER_IPV4_RANGE:-"10.120.0.0/16"}
CLUSTER_IPV6_RANGE=${CLUSTER_IPV6_RANGE:-"fc1e::/112"}
SERVICE_IPV4_RANGE=${SERVICE_IPV4_RANGE:-"10.121.0.0/16"}
SERVICE_IPV6_RANGE=${SERVICE_IPV6_RANGE:-"fc1f::/112"}
NODEGROUP_ZONE=${COMPUTE_DEFAULT_ZONE:-"ru-central1-d"}

echo "Provisioning e2e infrastructure..."

echo "Creating cluster..."
yc managed-kubernetes cluster create \
    --name e2e-controlplane --network-id ${NETWORK_ID} \
    --dual-stack \
    --cluster-ipv4-range ${CLUSTER_IPV4_RANGE} \
    --cluster-ipv6-range ${CLUSTER_IPV6_RANGE} \
    --service-ipv4-range ${SERVICE_IPV4_RANGE} \
    --service-ipv6-range ${SERVICE_IPV6_RANGE} \
    --service-account-id ${SA_ID} --node-service-account-id ${SA_ID} \
    --regional \
	--master-location zone=ru-central1-a \
	--master-location zone=ru-central1-b \
	--master-location zone=ru-central1-d \
    --public-ipv6 ${CLUSTER_IP}

echo "Creating node group..."
yc managed-kubernetes node-group create --name e2e-ng \
    --cluster-name e2e-controlplane --platform-id standard-v2 \
    --cores 2 --memory 4 --core-fraction 50 \
    --fixed-size 1 \
    --location zone=${NODEGROUP_ZONE} \
    --network-interface "ipv4-address=nat,ipv6-address=auto,subnets=${SUBNET_ID}"

echo "Exporting credentials..."
yc managed-kubernetes cluster get-credentials --external-ipv6 --name e2e-controlplane --kubeconfig kubeconfig --force

echo "Creating CR..."
yc container registry list | grep crossplane-e2e-cr || yc container registry create crossplane-e2e-cr
yc container registry add-access-binding crossplane-e2e-cr --subject system:allUsers --role viewer
yc container registry configure-docker

echo "Done!"
