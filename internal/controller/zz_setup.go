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

	backendgroup "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/alb/backendgroup"
	httprouter "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/alb/httprouter"
	loadbalancer "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/alb/loadbalancer"
	targetgroup "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/alb/targetgroup"
	virtualhost "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/alb/virtualhost"
	gateway "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/api/gateway"
	disk "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/compute/disk"
	diskplacementgroup "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/compute/diskplacementgroup"
	image "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/compute/image"
	instance "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/compute/instance"
	instancegroup "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/compute/instancegroup"
	placementgroup "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/compute/placementgroup"
	snapshot "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/compute/snapshot"
	registry "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/container/registry"
	registryiambinding "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/container/registryiambinding"
	repository "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/container/repository"
	repositoryiambinding "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/container/repositoryiambinding"
	cluster "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/dataproc/cluster"
	recordset "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/dns/recordset"
	zone "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/dns/zone"
	iambinding "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/function/iambinding"
	scalingpolicy "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/function/scalingpolicy"
	trigger "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/function/trigger"
	serviceaccount "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/iam/serviceaccount"
	serviceaccountapikey "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/iam/serviceaccountapikey"
	serviceaccountiambinding "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/iam/serviceaccountiambinding"
	serviceaccountiammember "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/iam/serviceaccountiammember"
	serviceaccountiampolicy "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/iam/serviceaccountiampolicy"
	serviceaccountkey "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/iam/serviceaccountkey"
	serviceaccountstaticaccesskey "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/iam/serviceaccountstaticaccesskey"
	secretciphertext "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/kms/secretciphertext"
	symmetrickey "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/kms/symmetrickey"
	symmetrickeyiambinding "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/kms/symmetrickeyiambinding"
	clusterkubernetes "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/kubernetes/cluster"
	nodegroup "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/kubernetes/nodegroup"
	networkloadbalancer "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/lb/networkloadbalancer"
	targetgrouplb "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/lb/targetgroup"
	group "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/logging/group"
	clickhousecluster "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/mdb/clickhousecluster"
	elasticsearchcluster "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/mdb/elasticsearchcluster"
	greenplumcluster "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/mdb/greenplumcluster"
	kafkacluster "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/mdb/kafkacluster"
	kafkatopic "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/mdb/kafkatopic"
	mongodbcluster "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/mdb/mongodbcluster"
	mysqlcluster "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/mdb/mysqlcluster"
	postgresqlcluster "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/mdb/postgresqlcluster"
	rediscluster "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/mdb/rediscluster"
	sqlservercluster "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/mdb/sqlservercluster"
	queue "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/message/queue"
	organizationiambinding "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/organizationmanager/organizationiambinding"
	organizationiammember "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/organizationmanager/organizationiammember"
	samlfederation "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/organizationmanager/samlfederation"
	providerconfig "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/providerconfig"
	cloudiambinding "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/resourcemanager/cloudiambinding"
	cloudiammember "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/resourcemanager/cloudiammember"
	folder "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/resourcemanager/folder"
	folderiambinding "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/resourcemanager/folderiambinding"
	folderiammember "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/resourcemanager/folderiammember"
	folderiampolicy "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/resourcemanager/folderiampolicy"
	bucket "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/storage/bucket"
	object "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/storage/object"
	address "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/vpc/address"
	defaultsecuritygroup "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/vpc/defaultsecuritygroup"
	network "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/vpc/network"
	routetable "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/vpc/routetable"
	securitygroup "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/vpc/securitygroup"
	securitygrouprule "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/vpc/securitygrouprule"
	subnet "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/vpc/subnet"
	function "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/yandex/function"
	databasededicated "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/ydb/databasededicated"
	databaseserverless "github.com/crossplane-contrib/provider-jet-yandex-cloud/internal/controller/ydb/databaseserverless"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		backendgroup.Setup,
		httprouter.Setup,
		loadbalancer.Setup,
		targetgroup.Setup,
		virtualhost.Setup,
		gateway.Setup,
		disk.Setup,
		diskplacementgroup.Setup,
		image.Setup,
		instance.Setup,
		instancegroup.Setup,
		placementgroup.Setup,
		snapshot.Setup,
		registry.Setup,
		registryiambinding.Setup,
		repository.Setup,
		repositoryiambinding.Setup,
		cluster.Setup,
		recordset.Setup,
		zone.Setup,
		iambinding.Setup,
		scalingpolicy.Setup,
		trigger.Setup,
		serviceaccount.Setup,
		serviceaccountapikey.Setup,
		serviceaccountiambinding.Setup,
		serviceaccountiammember.Setup,
		serviceaccountiampolicy.Setup,
		serviceaccountkey.Setup,
		serviceaccountstaticaccesskey.Setup,
		secretciphertext.Setup,
		symmetrickey.Setup,
		symmetrickeyiambinding.Setup,
		clusterkubernetes.Setup,
		nodegroup.Setup,
		networkloadbalancer.Setup,
		targetgrouplb.Setup,
		group.Setup,
		clickhousecluster.Setup,
		elasticsearchcluster.Setup,
		greenplumcluster.Setup,
		kafkacluster.Setup,
		kafkatopic.Setup,
		mongodbcluster.Setup,
		mysqlcluster.Setup,
		postgresqlcluster.Setup,
		rediscluster.Setup,
		sqlservercluster.Setup,
		queue.Setup,
		organizationiambinding.Setup,
		organizationiammember.Setup,
		samlfederation.Setup,
		providerconfig.Setup,
		cloudiambinding.Setup,
		cloudiammember.Setup,
		folder.Setup,
		folderiambinding.Setup,
		folderiammember.Setup,
		folderiampolicy.Setup,
		bucket.Setup,
		object.Setup,
		address.Setup,
		defaultsecuritygroup.Setup,
		network.Setup,
		routetable.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		function.Setup,
		databasededicated.Setup,
		databaseserverless.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
