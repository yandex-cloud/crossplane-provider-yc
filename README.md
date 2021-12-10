# Terrajet Yandex.Cloud Provider

`provider-jet-yc` is a [Crossplane](https://crossplane.io/) provider that is built
using [Terrajet](https://github.com/crossplane-contrib/terrajet) code generation tools and exposes XRM-conformant
managed resources for
[Yandex.Cloud](https://cloud.yandex.com/).

## Getting Started

Install the provider by using the following command after changing the image tag to
the [latest release](https://github.com/crossplane-contrib/provider-jet-yc/releases):

```
kubectl crossplane install provider cr.yandex/crp0kch415f0lke009ft/crossplane/provider-jet-yc-amd64:v0.0.1
```

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-jet-yc).

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please open
an [issue](https://github.com/crossplane/provider-jet-yc/issues).
