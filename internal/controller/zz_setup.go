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

	"github.com/crossplane/upjet/pkg/controller"

	backendgroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/alb/backendgroup"
	httprouter "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/alb/httprouter"
	loadbalancer "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/alb/loadbalancer"
	targetgroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/alb/targetgroup"
	virtualhost "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/alb/virtualhost"
	origingroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/cdn/origingroup"
	resource "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/cdn/resource"
	instance "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/compute/instance"
	registry "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/container/registry"
	repository "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/container/repository"
	endpoint "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/datatransfer/endpoint"
	transfer "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/datatransfer/transfer"
	recordset "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/dns/recordset"
	zone "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/dns/zone"
	cloudiambinding "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/iam/cloudiambinding"
	cloudiammember "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/iam/cloudiammember"
	folderiambinding "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/iam/folderiambinding"
	folderiammember "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/iam/folderiammember"
	groupiammember "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/iam/groupiammember"
	organizationiambinding "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/iam/organizationiambinding"
	serviceaccount "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/iam/serviceaccount"
	serviceaccountapikey "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/iam/serviceaccountapikey"
	serviceaccountiambinding "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/iam/serviceaccountiambinding"
	serviceaccountiammember "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/iam/serviceaccountiammember"
	serviceaccountkey "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/iam/serviceaccountkey"
	serviceaccountstaticaccesskey "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/iam/serviceaccountstaticaccesskey"
	symmetrickey "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/kms/symmetrickey"
	symmetrickeyiambinding "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/kms/symmetrickeyiambinding"
	cluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/kubernetes/cluster"
	nodegroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/kubernetes/nodegroup"
	networkloadbalancer "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/lb/networkloadbalancer"
	targetgrouplb "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/lb/targetgroup"
	elasticsearchcluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/elasticsearchcluster"
	kafkacluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/kafkacluster"
	kafkaconnector "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/kafkaconnector"
	kafkatopic "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/kafkatopic"
	kafkauser "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/kafkauser"
	mongodbcluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/mongodbcluster"
	mongodbdatabase "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/mongodbdatabase"
	mongodbuser "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/mongodbuser"
	mysqlcluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/mysqlcluster"
	mysqldatabase "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/mysqldatabase"
	mysqluser "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/mysqluser"
	opensearchcluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/opensearchcluster"
	postgresqlcluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/postgresqlcluster"
	postgresqldatabase "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/postgresqldatabase"
	postgresqluser "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/postgresqluser"
	rediscluster "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/mdb/rediscluster"
	queue "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/message/queue"
	group "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/organizationmanager/group"
	samlfederation "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/organizationmanager/samlfederation"
	samlfederationuseraccount "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/organizationmanager/samlfederationuseraccount"
	providerconfig "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/providerconfig"
	cloud "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/resourcemanager/cloud"
	folder "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/resourcemanager/folder"
	bucket "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/storage/bucket"
	object "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/storage/object"
	address "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/vpc/address"
	defaultsecuritygroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/vpc/defaultsecuritygroup"
	gateway "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/vpc/gateway"
	network "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/vpc/network"
	routetable "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/vpc/routetable"
	securitygroup "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/vpc/securitygroup"
	securitygrouprule "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/vpc/securitygrouprule"
	subnet "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/vpc/subnet"
	databasededicated "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/ydb/databasededicated"
	databaseserverless "github.com/yandex-cloud/crossplane-provider-yc/internal/controller/ydb/databaseserverless"
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
		resource.Setup,
		instance.Setup,
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
		opensearchcluster.Setup,
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
