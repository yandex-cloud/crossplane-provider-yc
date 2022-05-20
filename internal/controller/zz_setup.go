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
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	instance "github.com/yandex-cloud/provider-jet-yc/internal/controller/compute/instance"
	registry "github.com/yandex-cloud/provider-jet-yc/internal/controller/container/registry"
	repository "github.com/yandex-cloud/provider-jet-yc/internal/controller/container/repository"
	recordset "github.com/yandex-cloud/provider-jet-yc/internal/controller/dns/recordset"
	zone "github.com/yandex-cloud/provider-jet-yc/internal/controller/dns/zone"
	folderiambinding "github.com/yandex-cloud/provider-jet-yc/internal/controller/iam/folderiambinding"
	folderiammember "github.com/yandex-cloud/provider-jet-yc/internal/controller/iam/folderiammember"
	serviceaccount "github.com/yandex-cloud/provider-jet-yc/internal/controller/iam/serviceaccount"
	serviceaccountiammember "github.com/yandex-cloud/provider-jet-yc/internal/controller/iam/serviceaccountiammember"
	serviceaccountkey "github.com/yandex-cloud/provider-jet-yc/internal/controller/iam/serviceaccountkey"
	serviceaccountstaticaccesskey "github.com/yandex-cloud/provider-jet-yc/internal/controller/iam/serviceaccountstaticaccesskey"
	symmetrickey "github.com/yandex-cloud/provider-jet-yc/internal/controller/kms/symmetrickey"
	cluster "github.com/yandex-cloud/provider-jet-yc/internal/controller/kubernetes/cluster"
	nodegroup "github.com/yandex-cloud/provider-jet-yc/internal/controller/kubernetes/nodegroup"
	mongodbcluster "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/mongodbcluster"
	postgresqlcluster "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/postgresqlcluster"
	rediscluster "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/rediscluster"
	providerconfig "github.com/yandex-cloud/provider-jet-yc/internal/controller/providerconfig"
	folder "github.com/yandex-cloud/provider-jet-yc/internal/controller/resourcemanager/folder"
	bucket "github.com/yandex-cloud/provider-jet-yc/internal/controller/storage/bucket"
	object "github.com/yandex-cloud/provider-jet-yc/internal/controller/storage/object"
	defaultsecuritygroup "github.com/yandex-cloud/provider-jet-yc/internal/controller/vpc/defaultsecuritygroup"
	network "github.com/yandex-cloud/provider-jet-yc/internal/controller/vpc/network"
	securitygroup "github.com/yandex-cloud/provider-jet-yc/internal/controller/vpc/securitygroup"
	securitygrouprule "github.com/yandex-cloud/provider-jet-yc/internal/controller/vpc/securitygrouprule"
	subnet "github.com/yandex-cloud/provider-jet-yc/internal/controller/vpc/subnet"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		instance.Setup,
		registry.Setup,
		repository.Setup,
		recordset.Setup,
		zone.Setup,
		folderiambinding.Setup,
		folderiammember.Setup,
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
		defaultsecuritygroup.Setup,
		network.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
