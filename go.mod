module github.com/yandex-cloud/provider-jet-yc

go 1.16

require (
	github.com/benagricola/crossplane-composition-generator v0.0.0-20210505120457-1b2497dd442b
	github.com/crossplane/crossplane-runtime v0.15.1-0.20220315141414-988c9ba9c255
	github.com/crossplane/crossplane-tools v0.0.0-20220310165030-1f43fc12793e
	github.com/crossplane/terrajet v0.4.0-rc.0.0.20220512072756-6e78c0471d62
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.10.0
	github.com/pkg/errors v0.9.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.23.0
	k8s.io/client-go v0.23.0
	sigs.k8s.io/controller-runtime v0.11.0
	sigs.k8s.io/controller-tools v0.8.0
)
