#!/bin/bash

NODEGROUP_ZONE=${COMPUTE_DEFAULT_ZONE:-"ru-central1-d"}

# Find a free 10.X.0.0/16 IPv4 CIDR not overlapping any in the used list.
# Usage: find_free_ipv4 <used_cidrs_newline_separated>
find_free_ipv4() {
    python3 -c '
import sys, random, ipaddress
used = [ipaddress.ip_network(c) for c in sys.argv[1].split()]
for _ in range(100):
    cand = ipaddress.ip_network(f"10.{random.randint(0, 255)}.0.0/16")
    if not any(cand.overlaps(u) for u in used):
        print(cand)
        sys.exit(0)
sys.exit("ERROR: could not find a free IPv4 CIDR after 100 attempts")' "$1"
}

# Find a free fcXY::/PREFIX IPv6 CIDR not overlapping any in the used list.
# Usage: find_free_ipv6 <used_cidrs_newline_separated> <prefix_len>
find_free_ipv6() {
    python3 -c '
import sys, random, ipaddress
used = [ipaddress.ip_network(c) for c in sys.argv[1].split()]
prefix = sys.argv[2]
for _ in range(100):
    cand = ipaddress.ip_network(f"fc{random.randint(0, 255):02x}::/{prefix}")
    if not any(cand.overlaps(u) for u in used):
        print(cand)
        sys.exit(0)
sys.exit("ERROR: could not find a free IPv6 CIDR after 100 attempts")' "$1" "$2"
}

echo "Provisioning e2e infrastructure..."

echo "Collecting existing subnet CIDRs from network ${NETWORK_ID}..."
SUBNET_JSON=$(yc vpc network list-subnets --id "${NETWORK_ID}" --format json)

USED_IPV4=$(echo "${SUBNET_JSON}" | jq -r '.[].v4_cidr_blocks[]')
USED_IPV6=$(echo "${SUBNET_JSON}" | jq -r '.[].v6_cidr_blocks[]')

echo "Finding free cluster IPv4 range..."
CLUSTER_IPV4_RANGE=$(find_free_ipv4 "${USED_IPV4}")
echo "  -> ${CLUSTER_IPV4_RANGE}"

echo "Finding free service IPv4 range..."
SERVICE_IPV4_RANGE=$(find_free_ipv4 "${USED_IPV4}
${CLUSTER_IPV4_RANGE}")
echo "  -> ${SERVICE_IPV4_RANGE}"

echo "Finding free cluster IPv6 range..."
CLUSTER_IPV6_RANGE=$(find_free_ipv6 "${USED_IPV6}" "112")
echo "  -> ${CLUSTER_IPV6_RANGE}"

echo "Finding free service IPv6 range..."
SERVICE_IPV6_RANGE=$(find_free_ipv6 "${USED_IPV6}
${CLUSTER_IPV6_RANGE}" "112")
echo "  -> ${SERVICE_IPV6_RANGE}"

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
