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

type InboundRuleInitParameters struct {

	// The action to take when rule match. Possible values are: accept or drop.
	// Action when rule match request (drop or accept)
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The ip this rule apply to. If no ip nor ip_range are specified, rule will apply to all ip. Only one of ip and ip_range should be specified.
	// Ip address for this rule (e.g: 1.1.1.1). Only one of ip or ip_range should be provided
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The ip range (e.g 192.168.1.0/24) this rule applies to. If no ip nor ip_range are specified, rule will apply to all ip. Only one of ip and ip_range should be specified.
	// Ip range for this rule (e.g: 192.168.1.0/24). Only one of ip or ip_range should be provided
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// The port this rule applies to. If no port nor port_range are specified, the rule will apply to all port. Only one of port and port_range should be specified.
	// Network port for this rule
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// 13.0  The port range (e.g 22-23) this rule applies to.
	// Port range MUST comply the Scaleway-notation: interval between ports must be a power of 2 2^X-1 number (e.g 2^13-1=8191 in port_range = "10000-18191").
	// If no port nor port_range are specified, rule will apply to all port.
	// Only one of port and port_range should be specified.
	// Computed port range for this rule (e.g: 1-1024, 22-22)
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// (Defaults to TCP) The protocol this rule apply to. Possible values are: TCP, UDP, ICMP or ANY.
	// Protocol for this rule (TCP, UDP, ICMP or ANY)
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type InboundRuleObservation struct {

	// The action to take when rule match. Possible values are: accept or drop.
	// Action when rule match request (drop or accept)
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The ip this rule apply to. If no ip nor ip_range are specified, rule will apply to all ip. Only one of ip and ip_range should be specified.
	// Ip address for this rule (e.g: 1.1.1.1). Only one of ip or ip_range should be provided
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The ip range (e.g 192.168.1.0/24) this rule applies to. If no ip nor ip_range are specified, rule will apply to all ip. Only one of ip and ip_range should be specified.
	// Ip range for this rule (e.g: 192.168.1.0/24). Only one of ip or ip_range should be provided
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// The port this rule applies to. If no port nor port_range are specified, the rule will apply to all port. Only one of port and port_range should be specified.
	// Network port for this rule
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// 13.0  The port range (e.g 22-23) this rule applies to.
	// Port range MUST comply the Scaleway-notation: interval between ports must be a power of 2 2^X-1 number (e.g 2^13-1=8191 in port_range = "10000-18191").
	// If no port nor port_range are specified, rule will apply to all port.
	// Only one of port and port_range should be specified.
	// Computed port range for this rule (e.g: 1-1024, 22-22)
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// (Defaults to TCP) The protocol this rule apply to. Possible values are: TCP, UDP, ICMP or ANY.
	// Protocol for this rule (TCP, UDP, ICMP or ANY)
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type InboundRuleParameters struct {

	// The action to take when rule match. Possible values are: accept or drop.
	// Action when rule match request (drop or accept)
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// The ip this rule apply to. If no ip nor ip_range are specified, rule will apply to all ip. Only one of ip and ip_range should be specified.
	// Ip address for this rule (e.g: 1.1.1.1). Only one of ip or ip_range should be provided
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The ip range (e.g 192.168.1.0/24) this rule applies to. If no ip nor ip_range are specified, rule will apply to all ip. Only one of ip and ip_range should be specified.
	// Ip range for this rule (e.g: 192.168.1.0/24). Only one of ip or ip_range should be provided
	// +kubebuilder:validation:Optional
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// The port this rule applies to. If no port nor port_range are specified, the rule will apply to all port. Only one of port and port_range should be specified.
	// Network port for this rule
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// 13.0  The port range (e.g 22-23) this rule applies to.
	// Port range MUST comply the Scaleway-notation: interval between ports must be a power of 2 2^X-1 number (e.g 2^13-1=8191 in port_range = "10000-18191").
	// If no port nor port_range are specified, rule will apply to all port.
	// Only one of port and port_range should be specified.
	// Computed port range for this rule (e.g: 1-1024, 22-22)
	// +kubebuilder:validation:Optional
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// (Defaults to TCP) The protocol this rule apply to. Possible values are: TCP, UDP, ICMP or ANY.
	// Protocol for this rule (TCP, UDP, ICMP or ANY)
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type OutboundRuleInitParameters struct {

	// The action to take when rule match. Possible values are: accept or drop.
	// Action when rule match request (drop or accept)
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The ip this rule apply to. If no ip nor ip_range are specified, rule will apply to all ip. Only one of ip and ip_range should be specified.
	// Ip address for this rule (e.g: 1.1.1.1). Only one of ip or ip_range should be provided
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The ip range (e.g 192.168.1.0/24) this rule applies to. If no ip nor ip_range are specified, rule will apply to all ip. Only one of ip and ip_range should be specified.
	// Ip range for this rule (e.g: 192.168.1.0/24). Only one of ip or ip_range should be provided
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// The port this rule applies to. If no port nor port_range are specified, the rule will apply to all port. Only one of port and port_range should be specified.
	// Network port for this rule
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// 13.0  The port range (e.g 22-23) this rule applies to.
	// Port range MUST comply the Scaleway-notation: interval between ports must be a power of 2 2^X-1 number (e.g 2^13-1=8191 in port_range = "10000-18191").
	// If no port nor port_range are specified, rule will apply to all port.
	// Only one of port and port_range should be specified.
	// Computed port range for this rule (e.g: 1-1024, 22-22)
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// (Defaults to TCP) The protocol this rule apply to. Possible values are: TCP, UDP, ICMP or ANY.
	// Protocol for this rule (TCP, UDP, ICMP or ANY)
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type OutboundRuleObservation struct {

	// The action to take when rule match. Possible values are: accept or drop.
	// Action when rule match request (drop or accept)
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The ip this rule apply to. If no ip nor ip_range are specified, rule will apply to all ip. Only one of ip and ip_range should be specified.
	// Ip address for this rule (e.g: 1.1.1.1). Only one of ip or ip_range should be provided
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The ip range (e.g 192.168.1.0/24) this rule applies to. If no ip nor ip_range are specified, rule will apply to all ip. Only one of ip and ip_range should be specified.
	// Ip range for this rule (e.g: 192.168.1.0/24). Only one of ip or ip_range should be provided
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// The port this rule applies to. If no port nor port_range are specified, the rule will apply to all port. Only one of port and port_range should be specified.
	// Network port for this rule
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// 13.0  The port range (e.g 22-23) this rule applies to.
	// Port range MUST comply the Scaleway-notation: interval between ports must be a power of 2 2^X-1 number (e.g 2^13-1=8191 in port_range = "10000-18191").
	// If no port nor port_range are specified, rule will apply to all port.
	// Only one of port and port_range should be specified.
	// Computed port range for this rule (e.g: 1-1024, 22-22)
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// (Defaults to TCP) The protocol this rule apply to. Possible values are: TCP, UDP, ICMP or ANY.
	// Protocol for this rule (TCP, UDP, ICMP or ANY)
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type OutboundRuleParameters struct {

	// The action to take when rule match. Possible values are: accept or drop.
	// Action when rule match request (drop or accept)
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// The ip this rule apply to. If no ip nor ip_range are specified, rule will apply to all ip. Only one of ip and ip_range should be specified.
	// Ip address for this rule (e.g: 1.1.1.1). Only one of ip or ip_range should be provided
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The ip range (e.g 192.168.1.0/24) this rule applies to. If no ip nor ip_range are specified, rule will apply to all ip. Only one of ip and ip_range should be specified.
	// Ip range for this rule (e.g: 192.168.1.0/24). Only one of ip or ip_range should be provided
	// +kubebuilder:validation:Optional
	IPRange *string `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// The port this rule applies to. If no port nor port_range are specified, the rule will apply to all port. Only one of port and port_range should be specified.
	// Network port for this rule
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// 13.0  The port range (e.g 22-23) this rule applies to.
	// Port range MUST comply the Scaleway-notation: interval between ports must be a power of 2 2^X-1 number (e.g 2^13-1=8191 in port_range = "10000-18191").
	// If no port nor port_range are specified, rule will apply to all port.
	// Only one of port and port_range should be specified.
	// Computed port range for this rule (e.g: 1-1024, 22-22)
	// +kubebuilder:validation:Optional
	PortRange *string `json:"portRange,omitempty" tf:"port_range,omitempty"`

	// (Defaults to TCP) The protocol this rule apply to. Possible values are: TCP, UDP, ICMP or ANY.
	// Protocol for this rule (TCP, UDP, ICMP or ANY)
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type SecurityGroupInitParameters struct {

	// The description of the security group.
	// The description of the security group
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, open a support ticket.
	// Enable blocking of SMTP on IPv4 and IPv6
	EnableDefaultSecurity *bool `json:"enableDefaultSecurity,omitempty" tf:"enable_default_security,omitempty"`

	// (Defaults to false) A boolean to specify whether to use instance_security_group_rules.
	// If external_rules is set to true, inbound_rule and outbound_rule can not be set directly in the security group.
	ExternalRules *bool `json:"externalRules,omitempty" tf:"external_rules,omitempty"`

	// (Defaults to accept) The default policy on incoming traffic. Possible values are: accept or drop.
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy *string `json:"inboundDefaultPolicy,omitempty" tf:"inbound_default_policy,omitempty"`

	// A list of inbound rule to add to the security group. (Structure is documented below.)
	// Inbound rules for this security group
	InboundRule []InboundRuleInitParameters `json:"inboundRule,omitempty" tf:"inbound_rule,omitempty"`

	// The name of the security group.
	// The name of the security group
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to accept) The default policy on outgoing traffic. Possible values are: accept or drop.
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy *string `json:"outboundDefaultPolicy,omitempty" tf:"outbound_default_policy,omitempty"`

	// A list of outbound rule to add to the security group. (Structure is documented below.)
	// Outbound rules for this security group
	OutboundRule []OutboundRuleInitParameters `json:"outboundRule,omitempty" tf:"outbound_rule,omitempty"`

	// (Defaults to provider project_id) The ID of the project the security group is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to true) A boolean to specify whether the security group should be stateful or not.
	// The stateful value of the security group
	Stateful *bool `json:"stateful,omitempty" tf:"stateful,omitempty"`

	// The tags of the security group.
	// The tags associated with the security group
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to provider zone) The zone in which the security group should be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type SecurityGroupObservation struct {

	// The description of the security group.
	// The description of the security group
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, open a support ticket.
	// Enable blocking of SMTP on IPv4 and IPv6
	EnableDefaultSecurity *bool `json:"enableDefaultSecurity,omitempty" tf:"enable_default_security,omitempty"`

	// (Defaults to false) A boolean to specify whether to use instance_security_group_rules.
	// If external_rules is set to true, inbound_rule and outbound_rule can not be set directly in the security group.
	ExternalRules *bool `json:"externalRules,omitempty" tf:"external_rules,omitempty"`

	// The ID of the security group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Defaults to accept) The default policy on incoming traffic. Possible values are: accept or drop.
	// Default inbound traffic policy for this security group
	InboundDefaultPolicy *string `json:"inboundDefaultPolicy,omitempty" tf:"inbound_default_policy,omitempty"`

	// A list of inbound rule to add to the security group. (Structure is documented below.)
	// Inbound rules for this security group
	InboundRule []InboundRuleObservation `json:"inboundRule,omitempty" tf:"inbound_rule,omitempty"`

	// The name of the security group.
	// The name of the security group
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The organization ID the security group is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// (Defaults to accept) The default policy on outgoing traffic. Possible values are: accept or drop.
	// Default outbound traffic policy for this security group
	OutboundDefaultPolicy *string `json:"outboundDefaultPolicy,omitempty" tf:"outbound_default_policy,omitempty"`

	// A list of outbound rule to add to the security group. (Structure is documented below.)
	// Outbound rules for this security group
	OutboundRule []OutboundRuleObservation `json:"outboundRule,omitempty" tf:"outbound_rule,omitempty"`

	// (Defaults to provider project_id) The ID of the project the security group is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to true) A boolean to specify whether the security group should be stateful or not.
	// The stateful value of the security group
	Stateful *bool `json:"stateful,omitempty" tf:"stateful,omitempty"`

	// The tags of the security group.
	// The tags associated with the security group
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to provider zone) The zone in which the security group should be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type SecurityGroupParameters struct {

	// The description of the security group.
	// The description of the security group
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to block SMTP on IPv4/IPv6 (Port 25, 465, 587). Set to false will unblock SMTP if your account is authorized to. If your organization is not yet authorized to send SMTP traffic, open a support ticket.
	// Enable blocking of SMTP on IPv4 and IPv6
	// +kubebuilder:validation:Optional
	EnableDefaultSecurity *bool `json:"enableDefaultSecurity,omitempty" tf:"enable_default_security,omitempty"`

	// (Defaults to false) A boolean to specify whether to use instance_security_group_rules.
	// If external_rules is set to true, inbound_rule and outbound_rule can not be set directly in the security group.
	// +kubebuilder:validation:Optional
	ExternalRules *bool `json:"externalRules,omitempty" tf:"external_rules,omitempty"`

	// (Defaults to accept) The default policy on incoming traffic. Possible values are: accept or drop.
	// Default inbound traffic policy for this security group
	// +kubebuilder:validation:Optional
	InboundDefaultPolicy *string `json:"inboundDefaultPolicy,omitempty" tf:"inbound_default_policy,omitempty"`

	// A list of inbound rule to add to the security group. (Structure is documented below.)
	// Inbound rules for this security group
	// +kubebuilder:validation:Optional
	InboundRule []InboundRuleParameters `json:"inboundRule,omitempty" tf:"inbound_rule,omitempty"`

	// The name of the security group.
	// The name of the security group
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to accept) The default policy on outgoing traffic. Possible values are: accept or drop.
	// Default outbound traffic policy for this security group
	// +kubebuilder:validation:Optional
	OutboundDefaultPolicy *string `json:"outboundDefaultPolicy,omitempty" tf:"outbound_default_policy,omitempty"`

	// A list of outbound rule to add to the security group. (Structure is documented below.)
	// Outbound rules for this security group
	// +kubebuilder:validation:Optional
	OutboundRule []OutboundRuleParameters `json:"outboundRule,omitempty" tf:"outbound_rule,omitempty"`

	// (Defaults to provider project_id) The ID of the project the security group is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to true) A boolean to specify whether the security group should be stateful or not.
	// The stateful value of the security group
	// +kubebuilder:validation:Optional
	Stateful *bool `json:"stateful,omitempty" tf:"stateful,omitempty"`

	// The tags of the security group.
	// The tags associated with the security group
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Defaults to provider zone) The zone in which the security group should be created.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// SecurityGroupSpec defines the desired state of SecurityGroup
type SecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupParameters `json:"forProvider"`
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
	InitProvider SecurityGroupInitParameters `json:"initProvider,omitempty"`
}

// SecurityGroupStatus defines the observed state of SecurityGroup.
type SecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroup is the Schema for the SecurityGroups API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupSpec   `json:"spec"`
	Status            SecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupList contains a list of SecurityGroups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroup_Kind             = "SecurityGroup"
	SecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroup_Kind}.String()
	SecurityGroup_KindAPIVersion   = SecurityGroup_Kind + "." + CRDGroupVersion.String()
	SecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}
