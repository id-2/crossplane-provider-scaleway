/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceInitParameters struct {

	// Boolean to store logical backups in the same region as the database instance.
	// Boolean to store logical backups in the same region as the database instance
	BackupSameRegion *bool `json:"backupSameRegion,omitempty" tf:"backup_same_region,omitempty"`

	// Backup schedule frequency in hours.
	// Backup schedule frequency in hours
	BackupScheduleFrequency *float64 `json:"backupScheduleFrequency,omitempty" tf:"backup_schedule_frequency,omitempty"`

	// Backup schedule retention in days.
	// Backup schedule retention in days
	BackupScheduleRetention *float64 `json:"backupScheduleRetention,omitempty" tf:"backup_schedule_retention,omitempty"`

	// Disable automated backup for the database instance.
	// Disable automated backup for the database instance
	DisableBackup *bool `json:"disableBackup,omitempty" tf:"disable_backup,omitempty"`

	// Database Instance's engine version (e.g. PostgreSQL-11).
	// Database's engine version id
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Map of engine settings to be set at database initialisation.
	// Map of engine settings to be set at database initialisation.
	InitSettings map[string]*string `json:"initSettings,omitempty" tf:"init_settings,omitempty"`

	// Enable or disable high availability for the database instance.
	// Enable or disable high availability for the database instance
	IsHaCluster *bool `json:"isHaCluster,omitempty" tf:"is_ha_cluster,omitempty"`

	// The name of the Database Instance.
	// Name of the database instance
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of database instance you want to create (e.g. db-dev-s).
	// The type of database instance you want to create
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// List of private networks endpoints of the database instance.
	// List of private network to expose your database instance
	PrivateNetwork []PrivateNetworkInitParameters `json:"privateNetwork,omitempty" tf:"private_network,omitempty"`

	// (Defaults to provider project_id) The ID of the project the Database
	// Instance is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region) The region
	// in which the Database Instance should be created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Map of engine settings to be set. Using this option will override default config.
	// Map of engine settings to be set on a running instance.
	Settings map[string]*string `json:"settings,omitempty" tf:"settings,omitempty"`

	// The tags associated with the Database Instance.
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Identifier for the first user of the database instance.
	// Identifier for the first user of the database instance
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Volume size (in GB) when volume_type is set to bssd.
	// Volume size (in GB) when volume_type is not lssd
	VolumeSizeInGb *float64 `json:"volumeSizeInGb,omitempty" tf:"volume_size_in_gb,omitempty"`

	// Type of volume where data are stored (bssd or lssd).
	// Type of volume where data are stored
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type InstanceObservation struct {

	// Boolean to store logical backups in the same region as the database instance.
	// Boolean to store logical backups in the same region as the database instance
	BackupSameRegion *bool `json:"backupSameRegion,omitempty" tf:"backup_same_region,omitempty"`

	// Backup schedule frequency in hours.
	// Backup schedule frequency in hours
	BackupScheduleFrequency *float64 `json:"backupScheduleFrequency,omitempty" tf:"backup_schedule_frequency,omitempty"`

	// Backup schedule retention in days.
	// Backup schedule retention in days
	BackupScheduleRetention *float64 `json:"backupScheduleRetention,omitempty" tf:"backup_schedule_retention,omitempty"`

	// Certificate of the database instance.
	// Certificate of the database instance
	Certificate *string `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// Disable automated backup for the database instance.
	// Disable automated backup for the database instance
	DisableBackup *bool `json:"disableBackup,omitempty" tf:"disable_backup,omitempty"`

	// (Deprecated) The IP of the Database Instance.
	// Endpoint IP of the database instance
	EndpointIP *string `json:"endpointIp,omitempty" tf:"endpoint_ip,omitempty"`

	// (Deprecated) The port of the Database Instance.
	// Endpoint port of the database instance
	EndpointPort *float64 `json:"endpointPort,omitempty" tf:"endpoint_port,omitempty"`

	// Database Instance's engine version (e.g. PostgreSQL-11).
	// Database's engine version id
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// The ID of the Database Instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Map of engine settings to be set at database initialisation.
	// Map of engine settings to be set at database initialisation.
	InitSettings map[string]*string `json:"initSettings,omitempty" tf:"init_settings,omitempty"`

	// Enable or disable high availability for the database instance.
	// Enable or disable high availability for the database instance
	IsHaCluster *bool `json:"isHaCluster,omitempty" tf:"is_ha_cluster,omitempty"`

	// List of load balancer endpoints of the database instance.
	// Load balancer of the database instance
	LoadBalancer []LoadBalancerObservation `json:"loadBalancer,omitempty" tf:"load_balancer,omitempty"`

	// The name of the Database Instance.
	// Name of the database instance
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of database instance you want to create (e.g. db-dev-s).
	// The type of database instance you want to create
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// The organization ID the Database Instance is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// List of private networks endpoints of the database instance.
	// List of private network to expose your database instance
	PrivateNetwork []PrivateNetworkObservation `json:"privateNetwork,omitempty" tf:"private_network,omitempty"`

	// (Defaults to provider project_id) The ID of the project the Database
	// Instance is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// List of read replicas of the database instance.
	// Read replicas of the database instance
	ReadReplicas []ReadReplicasObservation `json:"readReplicas,omitempty" tf:"read_replicas,omitempty"`

	// (Defaults to provider region) The region
	// in which the Database Instance should be created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Map of engine settings to be set. Using this option will override default config.
	// Map of engine settings to be set on a running instance.
	Settings map[string]*string `json:"settings,omitempty" tf:"settings,omitempty"`

	// The tags associated with the Database Instance.
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Identifier for the first user of the database instance.
	// Identifier for the first user of the database instance
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Volume size (in GB) when volume_type is set to bssd.
	// Volume size (in GB) when volume_type is not lssd
	VolumeSizeInGb *float64 `json:"volumeSizeInGb,omitempty" tf:"volume_size_in_gb,omitempty"`

	// Type of volume where data are stored (bssd or lssd).
	// Type of volume where data are stored
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type InstanceParameters struct {

	// Boolean to store logical backups in the same region as the database instance.
	// Boolean to store logical backups in the same region as the database instance
	// +kubebuilder:validation:Optional
	BackupSameRegion *bool `json:"backupSameRegion,omitempty" tf:"backup_same_region,omitempty"`

	// Backup schedule frequency in hours.
	// Backup schedule frequency in hours
	// +kubebuilder:validation:Optional
	BackupScheduleFrequency *float64 `json:"backupScheduleFrequency,omitempty" tf:"backup_schedule_frequency,omitempty"`

	// Backup schedule retention in days.
	// Backup schedule retention in days
	// +kubebuilder:validation:Optional
	BackupScheduleRetention *float64 `json:"backupScheduleRetention,omitempty" tf:"backup_schedule_retention,omitempty"`

	// Disable automated backup for the database instance.
	// Disable automated backup for the database instance
	// +kubebuilder:validation:Optional
	DisableBackup *bool `json:"disableBackup,omitempty" tf:"disable_backup,omitempty"`

	// Database Instance's engine version (e.g. PostgreSQL-11).
	// Database's engine version id
	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// Map of engine settings to be set at database initialisation.
	// Map of engine settings to be set at database initialisation.
	// +kubebuilder:validation:Optional
	InitSettings map[string]*string `json:"initSettings,omitempty" tf:"init_settings,omitempty"`

	// Enable or disable high availability for the database instance.
	// Enable or disable high availability for the database instance
	// +kubebuilder:validation:Optional
	IsHaCluster *bool `json:"isHaCluster,omitempty" tf:"is_ha_cluster,omitempty"`

	// The name of the Database Instance.
	// Name of the database instance
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type of database instance you want to create (e.g. db-dev-s).
	// The type of database instance you want to create
	// +kubebuilder:validation:Optional
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// Password for the first user of the database instance.
	// Password for the first user of the database instance
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// List of private networks endpoints of the database instance.
	// List of private network to expose your database instance
	// +kubebuilder:validation:Optional
	PrivateNetwork []PrivateNetworkParameters `json:"privateNetwork,omitempty" tf:"private_network,omitempty"`

	// (Defaults to provider project_id) The ID of the project the Database
	// Instance is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region) The region
	// in which the Database Instance should be created.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Map of engine settings to be set. Using this option will override default config.
	// Map of engine settings to be set on a running instance.
	// +kubebuilder:validation:Optional
	Settings map[string]*string `json:"settings,omitempty" tf:"settings,omitempty"`

	// The tags associated with the Database Instance.
	// List of tags ["tag1", "tag2", ...] attached to a database instance
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Identifier for the first user of the database instance.
	// Identifier for the first user of the database instance
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Volume size (in GB) when volume_type is set to bssd.
	// Volume size (in GB) when volume_type is not lssd
	// +kubebuilder:validation:Optional
	VolumeSizeInGb *float64 `json:"volumeSizeInGb,omitempty" tf:"volume_size_in_gb,omitempty"`

	// Type of volume where data are stored (bssd or lssd).
	// Type of volume where data are stored
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type LoadBalancerInitParameters struct {
}

type LoadBalancerObservation struct {

	// The ID of the endpoint of the load balancer.
	EndpointID *string `json:"endpointId,omitempty" tf:"endpoint_id,omitempty"`

	// Name of the endpoint.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// IP of the replica.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The name of the Database Instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Port of the replica.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type LoadBalancerParameters struct {
}

type PrivateNetworkInitParameters struct {

	// The IP network address within the private subnet. This must be an IPv4 address with a
	// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
	// service if not set.
	// The IP with the given mask within the private subnet
	IPNet *string `json:"ipNet,omitempty" tf:"ip_net,omitempty"`

	// The ID of the private network.
	// The private network ID
	PnID *string `json:"pnId,omitempty" tf:"pn_id,omitempty"`

	// Port of the replica.
	// The port of your private service
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PrivateNetworkObservation struct {

	// The ID of the endpoint of the load balancer.
	// The endpoint ID
	EndpointID *string `json:"endpointId,omitempty" tf:"endpoint_id,omitempty"`

	// Name of the endpoint.
	// The hostname of your endpoint
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// IP of the replica.
	// The IP of your Instance within the private service
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The IP network address within the private subnet. This must be an IPv4 address with a
	// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
	// service if not set.
	// The IP with the given mask within the private subnet
	IPNet *string `json:"ipNet,omitempty" tf:"ip_net,omitempty"`

	// The name of the Database Instance.
	// The name of your private service
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the private network.
	// The private network ID
	PnID *string `json:"pnId,omitempty" tf:"pn_id,omitempty"`

	// Port of the replica.
	// The port of your private service
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PrivateNetworkParameters struct {

	// The IP network address within the private subnet. This must be an IPv4 address with a
	// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
	// service if not set.
	// The IP with the given mask within the private subnet
	// +kubebuilder:validation:Optional
	IPNet *string `json:"ipNet,omitempty" tf:"ip_net,omitempty"`

	// The ID of the private network.
	// The private network ID
	// +kubebuilder:validation:Optional
	PnID *string `json:"pnId" tf:"pn_id,omitempty"`

	// Port of the replica.
	// The port of your private service
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ReadReplicasInitParameters struct {
}

type ReadReplicasObservation struct {

	// IP of the replica.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The name of the Database Instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Port of the replica.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type ReadReplicasParameters struct {
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider InstanceInitParameters `json:"initProvider,omitempty"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.engine) || (has(self.initProvider) && has(self.initProvider.engine))",message="spec.forProvider.engine is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodeType) || (has(self.initProvider) && has(self.initProvider.nodeType))",message="spec.forProvider.nodeType is a required parameter"
	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
