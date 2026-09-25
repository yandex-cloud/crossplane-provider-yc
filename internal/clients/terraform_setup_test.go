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
	"net"
	"strings"
	"testing"
	"time"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/yandex-cloud/terraform-provider-yandex/yandex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	clustermdb "github.com/yandex-cloud/crossplane-provider-yc/apis/cluster/mdb/v1alpha1"
	clusterpc "github.com/yandex-cloud/crossplane-provider-yc/apis/cluster/v1beta1"
	namespacedmdb "github.com/yandex-cloud/crossplane-provider-yc/apis/namespaced/mdb/v1alpha1"
	namespacedpc "github.com/yandex-cloud/crossplane-provider-yc/apis/namespaced/v1beta1"
)

func TestTerraformSetupAsyncContext(t *testing.T) {
	// Configure the real Terraform provider with synthetic credentials and a
	// loopback endpoint. No cloud resources or credentials are needed.
	for _, key := range []string{
		"YC_SERVICE_ACCOUNT_KEY_FILE", "YC_STORAGE_ACCESS_KEY", "YC_STORAGE_SECRET_KEY", "TF_ENABLE_API_LOGGING",
		"HTTP_PROXY", "HTTPS_PROXY", "ALL_PROXY", "http_proxy", "https_proxy", "all_proxy",
	} {
		t.Setenv(key, "")
	}
	t.Setenv("YC_PLAINTEXT", "true")
	t.Setenv("YC_SHARED_CREDENTIALS_FILE", t.TempDir()+"/credentials")
	listener, err := (&net.ListenConfig{}).Listen(t.Context(), "tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	const reachedRPC = "Redis update reached the API"
	server := grpc.NewServer(grpc.UnknownServiceHandler(func(interface{}, grpc.ServerStream) error {
		return status.Error(codes.FailedPrecondition, reachedRPC)
	}))
	go func() { _ = server.Serve(listener) }()
	t.Cleanup(server.Stop)

	s := runtime.NewScheme()
	for _, add := range []func(*runtime.Scheme) error{
		clusterpc.SchemeBuilder.AddToScheme, clustermdb.SchemeBuilder.AddToScheme,
		namespacedpc.SchemeBuilder.AddToScheme, namespacedmdb.SchemeBuilder.AddToScheme,
	} {
		if err := add(s); err != nil {
			t.Fatal(err)
		}
	}
	legacy := &clustermdb.RedisCluster{ObjectMeta: metav1.ObjectMeta{Name: "test", UID: "legacy"}}
	legacy.SetProviderConfigReference(&xpv1.Reference{Name: "test"})
	modern := &namespacedmdb.RedisCluster{ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default", UID: "modern"}}
	modern.SetProviderConfigReference(&xpv1.ProviderConfigReference{Name: "test", Kind: "ProviderConfig"})
	for _, tc := range []struct {
		name string
		mg   resource.Managed
		pc   client.Object
	}{
		{
			name: "ClusterScoped", mg: legacy,
			pc: &clusterpc.ProviderConfig{
				ObjectMeta: metav1.ObjectMeta{Name: "test"},
				Spec: clusterpc.ProviderConfigSpec{Credentials: clusterpc.ProviderCredentials{
					Token: stringPtr("t1.synthetic.synthetic"), Endpoint: listener.Addr().String(), FolderID: "test-folder",
				}},
			},
		},
		{
			name: "Namespaced", mg: modern,
			pc: &namespacedpc.ProviderConfig{
				ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default"},
				Spec: namespacedpc.ProviderConfigSpec{Credentials: namespacedpc.ProviderCredentials{
					Token: stringPtr("t1.synthetic.synthetic"), Endpoint: listener.Addr().String(), FolderID: "test-folder",
				}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			providerCtx, stopProvider := context.WithCancel(t.Context())
			defer stopProvider()
			reconcileCtx, endReconcile := context.WithTimeout(providerCtx, time.Minute)
			defer endReconcile()
			p := yandex.NewSDKProvider()
			p.Schema["plaintext"].Default = true
			setup := TerraformSetupBuilder(providerCtx, "test", "yandex-cloud/yandex", "test", &config.Provider{TerraformProvider: p})
			kube := fake.NewClientBuilder().WithScheme(s).WithObjects(tc.pc).Build()
			ps, err := setup(reconcileCtx, kube, tc.mg)
			if err != nil {
				t.Fatal(err)
			}
			meta := ps.Meta.(*yandex.Config)
			endReconcile()
			if err := meta.Context().Err(); err != nil {
				t.Fatalf("saved provider context canceled after reconciliation: %v", err)
			}
			if deadline, ok := meta.Context().Deadline(); ok {
				if parentDeadline, parentOK := providerCtx.Deadline(); !parentOK || !deadline.Equal(parentDeadline) {
					t.Fatalf("saved provider context inherited the reconciliation deadline: %v", deadline)
				}
			}
			// Redis Update reads the cluster using the context saved in provider
			// metadata, even when Upjet supplies a fresh async operation context.
			r := p.ResourcesMap["yandex_mdb_redis_cluster"]
			d := schema.TestResourceDataRaw(t, r.Schema, map[string]interface{}{"name": "test", "folder_id": "test-folder"})
			d.SetId("test-cluster")
			if err := r.Update(d, meta); err == nil || !strings.Contains(err.Error(), reachedRPC) { //nolint:staticcheck // Redis uses this legacy SDK handler.
				t.Fatalf("Redis update did not reach the API after reconciliation ended: %v", err)
			}
			stopProvider()
			if err := meta.Context().Err(); err != context.Canceled {
				t.Fatalf("saved provider context did not cancel on shutdown: %v", err)
			}
			if err := r.Update(d, meta); err == nil || !strings.Contains(err.Error(), "context canceled") { //nolint:staticcheck // Redis uses this legacy SDK handler.
				t.Fatalf("Redis update did not stop on provider shutdown: %v", err)
			}
		})
	}
}
