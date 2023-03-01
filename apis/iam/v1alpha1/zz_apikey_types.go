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

type ApiKeyObservation struct {

	// The access key of the iam api key.
	// The access key of the iam api key
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// The date and time of the creation of the iam api key.
	// The date and time of the creation of the iam api key
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The IP Address of the device which created the API key.
	// The IPv4 Address of the device which created the API key
	CreationIP *string `json:"creationIp,omitempty" tf:"creation_ip,omitempty"`

	// Whether the iam api key is editable.
	// Whether or not the iam api key is editable
	Editable *bool `json:"editable,omitempty" tf:"editable,omitempty"`

	// The ID of the API key, which is the access key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date and time of the last update of the iam api key.
	// The date and time of the last update of the iam api key
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ApiKeyParameters struct {

	// :  ID of the application attached to the api key.
	// Only one of the application_id and user_id should be specified.
	// ID of the application attached to the api key
	// +crossplane:generate:reference:type=Application
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Reference to a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDRef *v1.Reference `json:"applicationIdRef,omitempty" tf:"-"`

	// Selector for a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// The default project ID to use with object storage.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	DefaultProjectID *string `json:"defaultProjectId,omitempty" tf:"default_project_id,omitempty"`

	// :  The description of the iam api key.
	// The description of the iam api key
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The date and time of the expiration of the iam api key. Please note that in case of change,
	// the resource will be recreated.
	// The date and time of the expiration of the iam api key. Cannot be changed afterwards
	// +kubebuilder:validation:Optional
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// ID of the user attached to the api key.
	// Only one of the application_id and user_id should be specified.
	// ID of the user attached to the api key
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// ApiKeySpec defines the desired state of ApiKey
type ApiKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApiKeyParameters `json:"forProvider"`
}

// ApiKeyStatus defines the observed state of ApiKey.
type ApiKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApiKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiKey is the Schema for the ApiKeys API. Manages Scaleway IAM API Keys.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type ApiKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiKeySpec   `json:"spec"`
	Status            ApiKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiKeyList contains a list of ApiKeys
type ApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiKey `json:"items"`
}

// Repository type metadata.
var (
	ApiKey_Kind             = "ApiKey"
	ApiKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApiKey_Kind}.String()
	ApiKey_KindAPIVersion   = ApiKey_Kind + "." + CRDGroupVersion.String()
	ApiKey_GroupVersionKind = CRDGroupVersion.WithKind(ApiKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ApiKey{}, &ApiKeyList{})
}
