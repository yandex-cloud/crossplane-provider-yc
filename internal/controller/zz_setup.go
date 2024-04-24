/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	backendgroup "github.com/yandex-cloud/provider-jet-yc/internal/controller/alb/backendgroup"
	httprouter "github.com/yandex-cloud/provider-jet-yc/internal/controller/alb/httprouter"
	loadbalancer "github.com/yandex-cloud/provider-jet-yc/internal/controller/alb/loadbalancer"
	targetgroup "github.com/yandex-cloud/provider-jet-yc/internal/controller/alb/targetgroup"
	virtualhost "github.com/yandex-cloud/provider-jet-yc/internal/controller/alb/virtualhost"
	instance "github.com/yandex-cloud/provider-jet-yc/internal/controller/compute/instance"
	registry "github.com/yandex-cloud/provider-jet-yc/internal/controller/container/registry"
	repository "github.com/yandex-cloud/provider-jet-yc/internal/controller/container/repository"
	endpoint "github.com/yandex-cloud/provider-jet-yc/internal/controller/datatransfer/endpoint"
	transfer "github.com/yandex-cloud/provider-jet-yc/internal/controller/datatransfer/transfer"
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
	elasticsearchcluster "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/elasticsearchcluster"
	kafkacluster "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/kafkacluster"
	kafkaconnector "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/kafkaconnector"
	kafkatopic "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/kafkatopic"
	kafkauser "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/kafkauser"
	mongodbcluster "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/mongodbcluster"
	mongodbdatabase "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/mongodbdatabase"
	mongodbuser "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/mongodbuser"
	mysqlcluster "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/mysqlcluster"
	mysqldatabase "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/mysqldatabase"
	mysqluser "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/mysqluser"
	postgresqlcluster "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/postgresqlcluster"
	postgresqldatabase "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/postgresqldatabase"
	postgresqluser "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/postgresqluser"
	rediscluster "github.com/yandex-cloud/provider-jet-yc/internal/controller/mdb/rediscluster"
	queue "github.com/yandex-cloud/provider-jet-yc/internal/controller/message/queue"
	providerconfig "github.com/yandex-cloud/provider-jet-yc/internal/controller/providerconfig"
	folder "github.com/yandex-cloud/provider-jet-yc/internal/controller/resourcemanager/folder"
	bucket "github.com/yandex-cloud/provider-jet-yc/internal/controller/storage/bucket"
	object "github.com/yandex-cloud/provider-jet-yc/internal/controller/storage/object"
	address "github.com/yandex-cloud/provider-jet-yc/internal/controller/vpc/address"
	defaultsecuritygroup "github.com/yandex-cloud/provider-jet-yc/internal/controller/vpc/defaultsecuritygroup"
	network "github.com/yandex-cloud/provider-jet-yc/internal/controller/vpc/network"
	securitygroup "github.com/yandex-cloud/provider-jet-yc/internal/controller/vpc/securitygroup"
	securitygrouprule "github.com/yandex-cloud/provider-jet-yc/internal/controller/vpc/securitygrouprule"
	subnet "github.com/yandex-cloud/provider-jet-yc/internal/controller/vpc/subnet"
	databasededicated "github.com/yandex-cloud/provider-jet-yc/internal/controller/ydb/databasededicated"
	databaseserverless "github.com/yandex-cloud/provider-jet-yc/internal/controller/ydb/databaseserverless"
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
		instance.Setup,
		registry.Setup,
		repository.Setup,
		endpoint.Setup,
		transfer.Setup,
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
		providerconfig.Setup,
		folder.Setup,
		bucket.Setup,
		object.Setup,
		address.Setup,
		defaultsecuritygroup.Setup,
		network.Setup,
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
