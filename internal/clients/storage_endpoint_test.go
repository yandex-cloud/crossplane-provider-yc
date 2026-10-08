/*
Copyright 2026 YANDEX LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clients

import (
	"context"
	"testing"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/yandex-cloud/terraform-provider-yandex/yandex"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	clusterstorage "github.com/yandex-cloud/crossplane-provider-yc/apis/cluster/storage/v1alpha1"
	clusterpc "github.com/yandex-cloud/crossplane-provider-yc/apis/cluster/v1beta1"
	namespacedstorage "github.com/yandex-cloud/crossplane-provider-yc/apis/namespaced/storage/v1alpha1"
	namespacedpc "github.com/yandex-cloud/crossplane-provider-yc/apis/namespaced/v1beta1"
)

func TestTerraformSetupStorageEndpoint(t *testing.T) {
	for _, key := range []string{
		"YC_SERVICE_ACCOUNT_KEY_FILE", "YC_STORAGE_ACCESS_KEY", "YC_STORAGE_SECRET_KEY", "TF_ENABLE_API_LOGGING",
	} {
		t.Setenv(key, "")
	}
	t.Setenv("YC_SHARED_CREDENTIALS_FILE", t.TempDir()+"/credentials")

	s := runtime.NewScheme()
	for _, add := range []func(*runtime.Scheme) error{
		clusterpc.SchemeBuilder.AddToScheme, clusterstorage.SchemeBuilder.AddToScheme,
		namespacedpc.SchemeBuilder.AddToScheme, namespacedstorage.SchemeBuilder.AddToScheme,
	} {
		if err := add(s); err != nil {
			t.Fatal(err)
		}
	}

	for _, endpointCase := range []struct {
		name, api, storage, wantStorage string
	}{
		{name: "Default", api: "api.cloud.yandex.net:443", wantStorage: "storage.yandexcloud.net"},
		{name: "Kazakhstan", api: "api.yandexcloud.kz:443", storage: "s3.yandexcloud.kz", wantStorage: "s3.yandexcloud.kz"},
		{name: "Custom", api: "api.example.test:443", storage: "s3.example.test", wantStorage: "s3.example.test"},
	} {
		t.Run(endpointCase.name, func(t *testing.T) {
			credentials := namespacedpc.ProviderCredentials{
				Token: stringPtr("t1.synthetic.synthetic"), Endpoint: endpointCase.api,
				StorageEndpoint: endpointCase.storage,
			}
			legacy := &clusterstorage.Bucket{ObjectMeta: metav1.ObjectMeta{Name: "test", UID: "legacy"}}
			legacy.SetProviderConfigReference(&xpv1.Reference{Name: "test"})
			namespaced := &namespacedstorage.Bucket{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default", UID: "namespaced"}}
			namespaced.SetProviderConfigReference(&xpv1.ProviderConfigReference{Name: "test", Kind: "ProviderConfig"})
			cluster := &namespacedstorage.Bucket{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default", UID: "cluster"}}
			cluster.SetProviderConfigReference(&xpv1.ProviderConfigReference{Name: "test", Kind: "ClusterProviderConfig"})

			for _, tc := range []struct {
				name string
				mg   resource.Managed
				pc   client.Object
			}{
				{
					name: "LegacyProviderConfig", mg: legacy,
					pc: &clusterpc.ProviderConfig{
						ObjectMeta: metav1.ObjectMeta{Name: "test"},
						Spec: clusterpc.ProviderConfigSpec{Credentials: clusterpc.ProviderCredentials{
							Token: credentials.Token, Endpoint: credentials.Endpoint, StorageEndpoint: credentials.StorageEndpoint,
						}},
					},
				},
				{
					name: "NamespacedProviderConfig", mg: namespaced,
					pc: &namespacedpc.ProviderConfig{
						ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default"},
						Spec:       namespacedpc.ProviderConfigSpec{Credentials: credentials},
					},
				},
				{
					name: "ClusterProviderConfig", mg: cluster,
					pc: &namespacedpc.ClusterProviderConfig{
						ObjectMeta: metav1.ObjectMeta{Name: "test"},
						Spec:       namespacedpc.ProviderConfigSpec{Credentials: credentials},
					},
				},
			} {
				t.Run(tc.name, func(t *testing.T) {
					p := yandex.NewSDKProvider()
					setup := TerraformSetupBuilder(t.Context(), "test", "yandex-cloud/yandex", "test", &config.Provider{TerraformProvider: p})
					kube := fake.NewClientBuilder().WithScheme(s).WithObjects(tc.pc).
						WithInterceptorFuncs(interceptor.Funcs{
							Get: func(ctx context.Context, c client.WithWatch, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
								// The real client ignores namespaces for cluster-scoped kinds.
								if _, ok := obj.(*namespacedpc.ClusterProviderConfig); ok {
									key.Namespace = ""
								}
								return c.Get(ctx, key, obj, opts...)
							},
						}).Build()
					ps, err := setup(t.Context(), kube, tc.mg)
					if err != nil {
						t.Fatal(err)
					}
					meta := ps.Meta.(*yandex.Config)
					if meta.StorageEndpoint != endpointCase.wantStorage {
						t.Errorf("configured storage endpoint = %q, want %q", meta.StorageEndpoint, endpointCase.wantStorage)
					}
					if ps.Configuration[endpoint] != endpointCase.api {
						t.Errorf("API endpoint = %v, want %q", ps.Configuration[endpoint], endpointCase.api)
					}
					got, present := ps.Configuration[storageEndpoint]
					if endpointCase.storage == "" {
						if present {
							t.Errorf("omitted storage endpoint overrides Terraform default with %v", got)
						}
					} else if got != endpointCase.storage {
						t.Errorf("Terraform storage endpoint = %v, want %q", got, endpointCase.storage)
					}
				})
			}
		})
	}
}
