/*
Copyright 2022 YANDEX LLC
This is modified version of the software, made by the Crossplane Authors
and available at: https://github.com/crossplane-contrib/provider-jet-template

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

	"github.com/crossplane/upjet/v2/pkg/controller"

	backendgroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/alb/backendgroup"
	httprouter "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/alb/httprouter"
	loadbalancer "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/alb/loadbalancer"
	targetgroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/alb/targetgroup"
	virtualhost "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/alb/virtualhost"
	origingroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/cdn/origingroup"
	disk "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/compute/disk"
	diskplacementgroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/compute/diskplacementgroup"
	filesystem "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/compute/filesystem"
	gpucluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/compute/gpucluster"
	image "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/compute/image"
	instance "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/compute/instance"
	instancegroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/compute/instancegroup"
	placementgroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/compute/placementgroup"
	snapshot "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/compute/snapshot"
	snapshotschedule "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/compute/snapshotschedule"
	registry "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/container/registry"
	repository "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/container/repository"
	endpoint "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/datatransfer/endpoint"
	transfer "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/datatransfer/transfer"
	recordset "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/dns/recordset"
	zone "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/dns/zone"
	cloudiambinding "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/iam/cloudiambinding"
	cloudiammember "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/iam/cloudiammember"
	folderiambinding "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/iam/folderiambinding"
	folderiammember "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/iam/folderiammember"
	groupiammember "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/iam/groupiammember"
	organizationiambinding "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/iam/organizationiambinding"
	serviceaccount "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/iam/serviceaccount"
	serviceaccountapikey "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/iam/serviceaccountapikey"
	serviceaccountiambinding "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/iam/serviceaccountiambinding"
	serviceaccountiammember "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/iam/serviceaccountiammember"
	serviceaccountkey "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/iam/serviceaccountkey"
	serviceaccountstaticaccesskey "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/iam/serviceaccountstaticaccesskey"
	symmetrickey "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/kms/symmetrickey"
	symmetrickeyiambinding "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/kms/symmetrickeyiambinding"
	cluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/kubernetes/cluster"
	nodegroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/kubernetes/nodegroup"
	networkloadbalancer "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/lb/networkloadbalancer"
	targetgrouplb "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/lb/targetgroup"
	clickhousecluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/clickhousecluster"
	elasticsearchcluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/elasticsearchcluster"
	kafkacluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/kafkacluster"
	kafkaconnector "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/kafkaconnector"
	kafkatopic "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/kafkatopic"
	kafkauser "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/kafkauser"
	mongodbcluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/mongodbcluster"
	mongodbdatabase "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/mongodbdatabase"
	mongodbuser "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/mongodbuser"
	mysqlcluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/mysqlcluster"
	mysqldatabase "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/mysqldatabase"
	mysqluser "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/mysqluser"
	postgresqlcluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/postgresqlcluster"
	postgresqldatabase "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/postgresqldatabase"
	postgresqluser "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/postgresqluser"
	rediscluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/mdb/rediscluster"
	queue "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/message/queue"
	group "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/organizationmanager/group"
	samlfederation "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/organizationmanager/samlfederation"
	samlfederationuseraccount "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/organizationmanager/samlfederationuseraccount"
	providerconfig "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/providerconfig"
	cloud "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/resourcemanager/cloud"
	folder "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/resourcemanager/folder"
	bucket "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/storage/bucket"
	object "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/storage/object"
	address "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/vpc/address"
	defaultsecuritygroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/vpc/defaultsecuritygroup"
	gateway "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/vpc/gateway"
	network "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/vpc/network"
	routetable "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/vpc/routetable"
	securitygroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/vpc/securitygroup"
	securitygrouprule "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/vpc/securitygrouprule"
	subnet "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/vpc/subnet"
	databasededicated "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/ydb/databasededicated"
	databaseserverless "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/namespaced/ydb/databaseserverless"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backendgroup.Setup,
		httprouter.Setup,
		loadbalancer.Setup,
		targetgroup.Setup,
		virtualhost.Setup,
		origingroup.Setup,
		disk.Setup,
		diskplacementgroup.Setup,
		filesystem.Setup,
		gpucluster.Setup,
		image.Setup,
		instance.Setup,
		instancegroup.Setup,
		placementgroup.Setup,
		snapshot.Setup,
		snapshotschedule.Setup,
		registry.Setup,
		repository.Setup,
		endpoint.Setup,
		transfer.Setup,
		recordset.Setup,
		zone.Setup,
		cloudiambinding.Setup,
		cloudiammember.Setup,
		folderiambinding.Setup,
		folderiammember.Setup,
		groupiammember.Setup,
		organizationiambinding.Setup,
		serviceaccount.Setup,
		serviceaccountapikey.Setup,
		serviceaccountiambinding.Setup,
		serviceaccountiammember.Setup,
		serviceaccountkey.Setup,
		serviceaccountstaticaccesskey.Setup,
		symmetrickey.Setup,
		symmetrickeyiambinding.Setup,
		cluster.Setup,
		nodegroup.Setup,
		networkloadbalancer.Setup,
		targetgrouplb.Setup,
		clickhousecluster.Setup,
		elasticsearchcluster.Setup,
		kafkacluster.Setup,
		kafkaconnector.Setup,
		kafkatopic.Setup,
		kafkauser.Setup,
		mongodbcluster.Setup,
		mongodbdatabase.Setup,
		mongodbuser.Setup,
		mysqlcluster.Setup,
		mysqldatabase.Setup,
		mysqluser.Setup,
		postgresqlcluster.Setup,
		postgresqldatabase.Setup,
		postgresqluser.Setup,
		rediscluster.Setup,
		queue.Setup,
		group.Setup,
		samlfederation.Setup,
		samlfederationuseraccount.Setup,
		providerconfig.Setup,
		cloud.Setup,
		folder.Setup,
		bucket.Setup,
		object.Setup,
		address.Setup,
		defaultsecuritygroup.Setup,
		gateway.Setup,
		network.Setup,
		routetable.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		subnet.Setup,
		databasededicated.Setup,
		databaseserverless.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backendgroup.SetupGated,
		httprouter.SetupGated,
		loadbalancer.SetupGated,
		targetgroup.SetupGated,
		virtualhost.SetupGated,
		origingroup.SetupGated,
		disk.SetupGated,
		diskplacementgroup.SetupGated,
		filesystem.SetupGated,
		gpucluster.SetupGated,
		image.SetupGated,
		instance.SetupGated,
		instancegroup.SetupGated,
		placementgroup.SetupGated,
		snapshot.SetupGated,
		snapshotschedule.SetupGated,
		registry.SetupGated,
		repository.SetupGated,
		endpoint.SetupGated,
		transfer.SetupGated,
		recordset.SetupGated,
		zone.SetupGated,
		cloudiambinding.SetupGated,
		cloudiammember.SetupGated,
		folderiambinding.SetupGated,
		folderiammember.SetupGated,
		groupiammember.SetupGated,
		organizationiambinding.SetupGated,
		serviceaccount.SetupGated,
		serviceaccountapikey.SetupGated,
		serviceaccountiambinding.SetupGated,
		serviceaccountiammember.SetupGated,
		serviceaccountkey.SetupGated,
		serviceaccountstaticaccesskey.SetupGated,
		symmetrickey.SetupGated,
		symmetrickeyiambinding.SetupGated,
		cluster.SetupGated,
		nodegroup.SetupGated,
		networkloadbalancer.SetupGated,
		targetgrouplb.SetupGated,
		clickhousecluster.SetupGated,
		elasticsearchcluster.SetupGated,
		kafkacluster.SetupGated,
		kafkaconnector.SetupGated,
		kafkatopic.SetupGated,
		kafkauser.SetupGated,
		mongodbcluster.SetupGated,
		mongodbdatabase.SetupGated,
		mongodbuser.SetupGated,
		mysqlcluster.SetupGated,
		mysqldatabase.SetupGated,
		mysqluser.SetupGated,
		postgresqlcluster.SetupGated,
		postgresqldatabase.SetupGated,
		postgresqluser.SetupGated,
		rediscluster.SetupGated,
		queue.SetupGated,
		group.SetupGated,
		samlfederation.SetupGated,
		samlfederationuseraccount.SetupGated,
		providerconfig.SetupGated,
		cloud.SetupGated,
		folder.SetupGated,
		bucket.SetupGated,
		object.SetupGated,
		address.SetupGated,
		defaultsecuritygroup.SetupGated,
		gateway.SetupGated,
		network.SetupGated,
		routetable.SetupGated,
		securitygroup.SetupGated,
		securitygrouprule.SetupGated,
		subnet.SetupGated,
		databasededicated.SetupGated,
		databaseserverless.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
