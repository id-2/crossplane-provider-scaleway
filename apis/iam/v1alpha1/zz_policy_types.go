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

type PolicyObservation struct {

	// The date and time of the creation of the policy.
	// The date and time of the creation of the policy
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Whether the policy is editable.
	// Whether or not the policy is editable.
	Editable *bool `json:"editable,omitempty" tf:"editable,omitempty"`

	// The ID of the policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time of the last update of the policy.
	// The date and time of the last update of the policy
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type PolicyParameters struct {

	// ID of the Application the policy will be linked to
	// Application id
	// +crossplane:generate:reference:type=Application
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Reference to a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDRef *v1.Reference `json:"applicationIdRef,omitempty" tf:"-"`

	// Selector for a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// The description of the iam policy.
	// The description of the iam policy
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the Group the policy will be linked to
	// Group id
	// +crossplane:generate:reference:type=Group
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// .The name of the iam policy.
	// The name of the iam policy
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If the policy doesn't apply to a principal.
	// Deactivate policy to a principal
	// +kubebuilder:validation:Optional
	NoPrincipal *bool `json:"noPrincipal,omitempty" tf:"no_principal,omitempty"`

	// (Defaults to provider organization_id) The ID of the organization the policy is associated with.
	// ID of organization the resource is associated to.
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// List of rules in the policy.
	// Rules of the policy to create
	// +kubebuilder:validation:Required
	Rule []RuleParameters `json:"rule" tf:"rule,omitempty"`

	// ID of the User the policy will be linked to
	// User id
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type RuleObservation struct {
}

type RuleParameters struct {

	// (Defaults to provider organization_id) The ID of the organization the policy is associated with.
	// ID of organization scoped to the rule. Only one of project_ids and organization_id may be set.
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Names of permission sets bound to the rule.
	// Names of permission sets bound to the rule.
	// +kubebuilder:validation:Required
	PermissionSetNames []*string `json:"permissionSetNames" tf:"permission_set_names,omitempty"`

	// List of project IDs scoped to the rule.
	// List of project IDs scoped to the rule. Only one of project_ids and organization_id may be set.
	// +kubebuilder:validation:Optional
	ProjectIds []*string `json:"projectIds,omitempty" tf:"project_ids,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API. Manages Scaleway IAM Policies.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec"`
	Status            PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
