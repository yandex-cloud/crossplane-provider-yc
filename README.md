# Yandex Cloud Crossplane Provider

`crossplane-provider-yc` is a [Crossplane](https://crossplane.io/) provider that is built
using [Upjet](https://github.com/crossplane/upjet/) code generation tools and exposes XRM-conformant
managed resources for [Yandex Cloud](https://cloud.yandex.com/).

## Getting Started

### Install crossplane-provider-yc

Install crossplane:

```
kubectl create namespace crossplane-system

helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm install crossplane --namespace crossplane-system crossplane-stable/crossplane 
```

Check crossplane status:

```
helm list -n crossplane-system

kubectl get all -n crossplane-system
```

Install crossplane CLI:

```shell
curl -sL https://raw.githubusercontent.com/crossplane/crossplane/master/install.sh | sh
sudo mv kubectl-crossplane $(dirname $(which kubectl))
```

Install the provider by using the following command after changing the image name and tag to the latest release
(image name and tag can be found in the [current release page](https://yandex.cloud/en-ru/marketplace/products/yc/crossplane#docker-images))

```
kubectl crossplane install provider cr.yandex/yc-marketplace/yandex-cloud/crossplane/provider-jet-yc1715238233618419317890636198513831921783193441180:v0.4.0
```

### Setup ProviderConfig

Create service account:

```
yc iam service-account create --name <service-account>
```

Add roles to this service account:

```shell
yc resource-manager folder add-access-binding <folder-id> --role <role>
```

Request key:

```shell
yc iam key create --service-account-id <service-account-id> --output key.json
```

Create k8s secret:

```shell
kubectl create secret generic yc-creds -n "crossplane-system" --from-file=credentials=./key.json
```

Apply example ProviderConfig:

```
kubectl apply -f examples/providerconfig/providerconfig.yaml
```

## Useful things

### Reconcile existing resources:

Add existing resource id `metadata.annotations["crossplane.io/external-name"]`

```yaml
metadata:
  annotations:
    crossplane.io/external-name: <cloud-resource-id>
```

To avoid making changes to an existing resouce that needs to be references (folder, subnet, etc.),
add `spec.managementPolicy: ObserveOnly`

```yaml
spec:
  managementPolicy: ObserveOnly
```

### Do not delete external resource with `kubectl delete`

Add `spec.deletionPolicy: Orphan`

```yaml
spec:
  deletionPolicy: Orphan
```

```shell
❯ k explain Folder.spec.deletionPolicy
KIND:     Folder
VERSION:  resourcemanager.yandex-cloud.jet.crossplane.io/v1alpha1

FIELD:    deletionPolicy <string>

DESCRIPTION:
     DeletionPolicy specifies what will happen to the underlying external when
     this managed resource is deleted - either "Delete" or "Orphan" the external
     resource.
```

You can enforce `deletionPolicy: Orphan` with gatekeeper if you need.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open
an [issue](https://github.com/yandex-cloud/crossplane-provider-yc/issues).
