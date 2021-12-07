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

	backendgroup "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/alb/backendgroup"
	httprouter "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/alb/httprouter"
	loadbalancer "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/alb/loadbalancer"
	targetgroup "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/alb/targetgroup"
	virtualhost "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/alb/virtualhost"
	gateway "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/api/gateway"
	disk "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/compute/disk"
	diskplacementgroup "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/compute/diskplacementgroup"
	image "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/compute/image"
	instance "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/compute/instance"
	instancegroup "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/compute/instancegroup"
	placementgroup "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/compute/placementgroup"
	snapshot "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/compute/snapshot"
	registry "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/container/registry"
	registryiambinding "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/container/registryiambinding"
	repository "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/container/repository"
	repositoryiambinding "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/container/repositoryiambinding"
	cluster "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/dataproc/cluster"
	recordset "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/dns/recordset"
	zone "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/dns/zone"
	iambinding "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/function/iambinding"
	scalingpolicy "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/function/scalingpolicy"
	trigger "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/function/trigger"
	serviceaccount "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/iam/serviceaccount"
	serviceaccountapikey "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/iam/serviceaccountapikey"
	serviceaccountiambinding "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/iam/serviceaccountiambinding"
	serviceaccountiammember "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/iam/serviceaccountiammember"
	serviceaccountiampolicy "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/iam/serviceaccountiampolicy"
	serviceaccountkey "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/iam/serviceaccountkey"
	serviceaccountstaticaccesskey "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/iam/serviceaccountstaticaccesskey"
	secretciphertext "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/kms/secretciphertext"
	symmetrickey "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/kms/symmetrickey"
	symmetrickeyiambinding "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/kms/symmetrickeyiambinding"
	clusterkubernetes "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/kubernetes/cluster"
	nodegroup "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/kubernetes/nodegroup"
	networkloadbalancer "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/lb/networkloadbalancer"
	targetgrouplb "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/lb/targetgroup"
	group "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/logging/group"
	clickhousecluster "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/mdb/clickhousecluster"
	elasticsearchcluster "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/mdb/elasticsearchcluster"
	greenplumcluster "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/mdb/greenplumcluster"
	kafkacluster "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/mdb/kafkacluster"
	kafkatopic "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/mdb/kafkatopic"
	mongodbcluster "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/mdb/mongodbcluster"
	mysqlcluster "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/mdb/mysqlcluster"
	postgresqlcluster "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/mdb/postgresqlcluster"
	rediscluster "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/mdb/rediscluster"
	sqlservercluster "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/mdb/sqlservercluster"
	queue "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/message/queue"
	organizationiambinding "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/organizationmanager/organizationiambinding"
	organizationiammember "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/organizationmanager/organizationiammember"
	samlfederation "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/organizationmanager/samlfederation"
	providerconfig "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/providerconfig"
	cloudiambinding "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/resourcemanager/cloudiambinding"
	cloudiammember "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/resourcemanager/cloudiammember"
	folder "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/resourcemanager/folder"
	folderiambinding "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/resourcemanager/folderiambinding"
	folderiammember "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/resourcemanager/folderiammember"
	folderiampolicy "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/resourcemanager/folderiampolicy"
	bucket "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/storage/bucket"
	object "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/storage/object"
	address "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/vpc/address"
	defaultsecuritygroup "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/vpc/defaultsecuritygroup"
	network "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/vpc/network"
	routetable "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/vpc/routetable"
	securitygroup "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/vpc/securitygroup"
	securitygrouprule "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/vpc/securitygrouprule"
	subnet "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/vpc/subnet"
	function "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/yandex/function"
	databasededicated "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/ydb/databasededicated"
	databaseserverless "bb.yandex-team.ru/crossplane/provider-jet-yandex-cloud/internal/controller/ydb/databaseserverless"
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
