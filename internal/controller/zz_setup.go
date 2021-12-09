/*
Copyright 2021 The Crossplane Authors.

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

package controller

import (
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/crossplane-contrib/terrajet/pkg/terraform"

	instance "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/compute/instance"
	registry "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/container/registry"
	repository "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/container/repository"
	recordset "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/dns/recordset"
	zone "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/dns/zone"
	serviceaccount "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/iam/serviceaccount"
	serviceaccountiammember "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/iam/serviceaccountiammember"
	serviceaccountkey "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/iam/serviceaccountkey"
	serviceaccountstaticaccesskey "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/iam/serviceaccountstaticaccesskey"
	symmetrickey "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/kms/symmetrickey"
	cluster "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/kubernetes/cluster"
	nodegroup "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/kubernetes/nodegroup"
	mongodbcluster "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/mdb/mongodbcluster"
	postgresqlcluster "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/mdb/postgresqlcluster"
	rediscluster "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/mdb/rediscluster"
	providerconfig "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/providerconfig"
	folder "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/resourcemanager/folder"
	bucket "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/storage/bucket"
	object "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/storage/object"
	network "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/vpc/network"
	subnet "bb.yandex-team.ru/crossplane/provider-jet-yc/internal/controller/vpc/subnet"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		instance.Setup,
		registry.Setup,
		repository.Setup,
		recordset.Setup,
		zone.Setup,
		serviceaccount.Setup,
		serviceaccountiammember.Setup,
		serviceaccountkey.Setup,
		serviceaccountstaticaccesskey.Setup,
		symmetrickey.Setup,
		cluster.Setup,
		nodegroup.Setup,
		mongodbcluster.Setup,
		postgresqlcluster.Setup,
		rediscluster.Setup,
		providerconfig.Setup,
		folder.Setup,
		bucket.Setup,
		object.Setup,
		network.Setup,
		subnet.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
