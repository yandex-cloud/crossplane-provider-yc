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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConnectionLimitsObservation struct {
}

type ConnectionLimitsParameters struct {

	// +kubebuilder:validation:Optional
	// Max connections per hour.
	MaxConnectionsPerHour *float64 `json:"maxConnectionsPerHour,omitempty" tf:"max_connections_per_hour,omitempty"`

	// +kubebuilder:validation:Optional
	// Max questions per hour.
	MaxQuestionsPerHour *float64 `json:"maxQuestionsPerHour,omitempty" tf:"max_questions_per_hour,omitempty"`

	// +kubebuilder:validation:Optional
	// Max updates per hour.
	MaxUpdatesPerHour *float64 `json:"maxUpdatesPerHour,omitempty" tf:"max_updates_per_hour,omitempty"`

	// +kubebuilder:validation:Optional
	// Max user connections.
	MaxUserConnections *float64 `json:"maxUserConnections,omitempty" tf:"max_user_connections,omitempty"`
}

type MySQLClusterAccessObservation struct {
}

type MySQLClusterAccessParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Allow access for [Yandex DataLens](https://cloud.yandex.com/services/datalens).
	DataLens *bool `json:"dataLens,omitempty" tf:"data_lens,omitempty"`

	// +kubebuilder:validation:Optional
	// Allows access for [SQL queries in the management console](https://cloud.yandex.com/docs/managed-mysql/operations/web-sql-query).
	WebSQL *bool `json:"webSql,omitempty" tf:"web_sql,omitempty"`
}

type MySQLClusterBackupWindowStartObservation struct {
}

type MySQLClusterBackupWindowStartParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) The hour at which backup will be started.
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The minute at which backup will be started.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type MySQLClusterDatabaseObservation struct {
}

type MySQLClusterDatabaseParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Name of the MySQL cluster. Provided by the client when the cluster is created.
	Name *string `json:"name" tf:"name,omitempty"`
}

type MySQLClusterHostObservation struct {
	// (Computed) The fully qualified domain name of the host.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// (Computed) Host replication source (fqdn), when replication_source is empty then host is in HA group.
	ReplicationSource *string `json:"replicationSource,omitempty" tf:"replication_source,omitempty"`
}

type MySQLClusterHostParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Sets whether the host should get a public IP address. It can be changed on the fly only when `name` is set.
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Host backup priority. Value is between 0 and 100, default is 0.
	BackupPriority *float64 `json:"backupPriority,omitempty" tf:"backup_priority,omitempty"`

	// +kubebuilder:validation:Optional
	// (Required) Name of the MySQL cluster. Provided by the client when the cluster is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Host master promotion priority. Value is between 0 and 100, default is 0.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Host replication source name points to host's `name` from which this host should replicate. When not set then host in HA group. It works only when `name` is set.
	ReplicationSourceName *string `json:"replicationSourceName,omitempty" tf:"replication_source_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	// (Optional) The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	// (Required) The availability zone where the MySQL host will be created.
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type MySQLClusterMaintenanceWindowObservation struct {
}

type MySQLClusterMaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Day of the week (in `DDD` format). Allowed values: "MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Hour of the day in UTC (in `HH` format). Allowed value is between 0 and 23.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Type of maintenance window. Can be either `ANYTIME` or `WEEKLY`. A day and hour of window need to be specified with weekly window.
	Type *string `json:"type" tf:"type,omitempty"`
}

type MySQLClusterObservation struct {
	// Creation timestamp of the cluster.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Aggregated health of the cluster.
	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) A host of the MySQL cluster. The structure is documented below.
	Host []MySQLClusterHostObservation `json:"host,omitempty" tf:"host,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Status of the cluster.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type MySQLClusterParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Access policy to the MySQL cluster. The structure is documented below.
	Access []MySQLClusterAccessParameters `json:"access,omitempty" tf:"access,omitempty"`

	// +kubebuilder:validation:Optional
	AllowRegenerationHost *bool `json:"allowRegenerationHost,omitempty" tf:"allow_regeneration_host,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Time to start the daily backup, in the UTC. The structure is documented below.
	BackupWindowStart []MySQLClusterBackupWindowStartParameters `json:"backupWindowStart,omitempty" tf:"backup_window_start,omitempty"`

	// +kubebuilder:validation:Optional
	// (Deprecated) To manage databases, please switch to using a separate resource type `yandex_mdb_mysql_databases`.
	Database []MySQLClusterDatabaseParameters `json:"database,omitempty" tf:"database,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Description of the MySQL cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Deployment environment of the MySQL cluster.
	Environment *string `json:"environment" tf:"environment,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The ID of the folder that the resource belongs to. If it
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) A host of the MySQL cluster. The structure is documented below.
	Host []MySQLClusterHostParameters `json:"host" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	HostGroupIds []*string `json:"hostGroupIds,omitempty" tf:"host_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A set of key/value label pairs to assign to the MySQL cluster.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Maintenance policy of the MySQL cluster. The structure is documented below.
	MaintenanceWindow []MySQLClusterMaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) MySQL cluster config. Detail info in "MySQL config" section (documented below).
	MySQLConfig map[string]*string `json:"mysqlConfig,omitempty" tf:"mysql_config,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	// (Required) ID of the network, to which the MySQL cluster uses.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional) Cluster performance diagnostics settings. The structure is documented below. [YC Documentation](https://cloud.yandex.com/en-ru/docs/managed-mysql/api-ref/grpc/cluster_service#PerformanceDiagnostics)
	PerformanceDiagnostics []PerformanceDiagnosticsParameters `json:"performanceDiagnostics,omitempty" tf:"performance_diagnostics,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Resources allocated to hosts of the MySQL cluster. The structure is documented below.
	Resources []MySQLClusterResourcesParameters `json:"resources" tf:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional, ForceNew) The cluster will be created from the specified backup. The structure is documented below.
	Restore []RestoreParameters `json:"restore,omitempty" tf:"restore,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	// (Optional) A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Deprecated) To manage users, please switch to using a separate resource type `yandex_mdb_mysql_user`.
	User []MySQLClusterUserParameters `json:"user,omitempty" tf:"user,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Version of the MySQL cluster. (allowed versions are: 5.7, 8.0)
	Version *string `json:"version" tf:"version,omitempty"`
}

type MySQLClusterResourcesObservation struct {
}

type MySQLClusterResourcesParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Volume of the storage available to a MySQL host, in gigabytes.
	DiskSize *float64 `json:"diskSize" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Type of the storage of MySQL hosts.
	DiskTypeID *string `json:"diskTypeId" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

type MySQLClusterUserObservation struct {
}

type MySQLClusterUserParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Authentication plugin. Allowed values: `MYSQL_NATIVE_PASSWORD`, `CACHING_SHA2_PASSWORD`, `SHA256_PASSWORD` (for version 5.7 `MYSQL_NATIVE_PASSWORD`, `SHA256_PASSWORD`)
	AuthenticationPlugin *string `json:"authenticationPlugin,omitempty" tf:"authentication_plugin,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) User's connection limits. The structure is documented below.
	ConnectionLimits []ConnectionLimitsParameters `json:"connectionLimits,omitempty" tf:"connection_limits,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) List user's global permissions     
	GlobalPermissions []*string `json:"globalPermissions,omitempty" tf:"global_permissions,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Name of the MySQL cluster. Provided by the client when the cluster is created.
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional) Set of permissions granted to the user. The structure is documented below.
	Permission []UserPermissionParameters `json:"permission,omitempty" tf:"permission,omitempty"`
}

type PerformanceDiagnosticsObservation struct {
}

type PerformanceDiagnosticsParameters struct {

	// +kubebuilder:validation:Optional
	// Enable performance diagnostics
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	// Interval (in seconds) for my_stat_activity sampling Acceptable values are 1 to 86400, inclusive.
	SessionsSamplingInterval *float64 `json:"sessionsSamplingInterval" tf:"sessions_sampling_interval,omitempty"`

	// +kubebuilder:validation:Required
	// Interval (in seconds) for my_stat_statements sampling Acceptable values are 1 to 86400, inclusive.
	StatementsSamplingInterval *float64 `json:"statementsSamplingInterval" tf:"statements_sampling_interval,omitempty"`
}

type RestoreObservation struct {
}

type RestoreParameters struct {

	// +kubebuilder:validation:Required
	// (Required, ForceNew) Backup ID. The cluster will be created from the specified backup. [How to get a list of MySQL backups](https://cloud.yandex.com/docs/managed-mysql/operations/cluster-backups). 
	BackupID *string `json:"backupId" tf:"backup_id,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional, ForceNew) Timestamp of the moment to which the MySQL cluster should be restored. (Format: "2006-01-02T15:04:05" - UTC). When not set, current time is used.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type UserPermissionObservation struct {
}

type UserPermissionParameters struct {

	// +kubebuilder:validation:Required
	// (Required) The name of the database that the permission grants access to.
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) List user's roles in the database.
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

// MySQLClusterSpec defines the desired state of MySQLCluster
type MySQLClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MySQLClusterParameters `json:"forProvider"`
}

// MySQLClusterStatus defines the observed state of MySQLCluster.
type MySQLClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MySQLClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLCluster is the Schema for the MySQLClusters API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type MySQLCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MySQLClusterSpec   `json:"spec"`
	Status            MySQLClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLClusterList contains a list of MySQLClusters
type MySQLClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLCluster `json:"items"`
}

// Repository type metadata.
var (
	MySQLCluster_Kind             = "MySQLCluster"
	MySQLCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MySQLCluster_Kind}.String()
	MySQLCluster_KindAPIVersion   = MySQLCluster_Kind + "." + CRDGroupVersion.String()
	MySQLCluster_GroupVersionKind = CRDGroupVersion.WithKind(MySQLCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&MySQLCluster{}, &MySQLClusterList{})
}
