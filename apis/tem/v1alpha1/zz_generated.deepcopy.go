//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Domain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainInitParameters) DeepCopyInto(out *DomainInitParameters) {
	*out = *in
	if in.AcceptTos != nil {
		in, out := &in.AcceptTos, &out.AcceptTos
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainInitParameters.
func (in *DomainInitParameters) DeepCopy() *DomainInitParameters {
	if in == nil {
		return nil
	}
	out := new(DomainInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainList) DeepCopyInto(out *DomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Domain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainList.
func (in *DomainList) DeepCopy() *DomainList {
	if in == nil {
		return nil
	}
	out := new(DomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainObservation) DeepCopyInto(out *DomainObservation) {
	*out = *in
	if in.AcceptTos != nil {
		in, out := &in.AcceptTos, &out.AcceptTos
		*out = new(bool)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.DKIMConfig != nil {
		in, out := &in.DKIMConfig, &out.DKIMConfig
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LastError != nil {
		in, out := &in.LastError, &out.LastError
		*out = new(string)
		**out = **in
	}
	if in.LastValidAt != nil {
		in, out := &in.LastValidAt, &out.LastValidAt
		*out = new(string)
		**out = **in
	}
	if in.MxBlackhole != nil {
		in, out := &in.MxBlackhole, &out.MxBlackhole
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NextCheckAt != nil {
		in, out := &in.NextCheckAt, &out.NextCheckAt
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Reputation != nil {
		in, out := &in.Reputation, &out.Reputation
		*out = make([]ReputationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RevokedAt != nil {
		in, out := &in.RevokedAt, &out.RevokedAt
		*out = new(string)
		**out = **in
	}
	if in.SMTPHost != nil {
		in, out := &in.SMTPHost, &out.SMTPHost
		*out = new(string)
		**out = **in
	}
	if in.SMTPPort != nil {
		in, out := &in.SMTPPort, &out.SMTPPort
		*out = new(float64)
		**out = **in
	}
	if in.SMTPPortAlternative != nil {
		in, out := &in.SMTPPortAlternative, &out.SMTPPortAlternative
		*out = new(float64)
		**out = **in
	}
	if in.SMTPPortUnsecure != nil {
		in, out := &in.SMTPPortUnsecure, &out.SMTPPortUnsecure
		*out = new(float64)
		**out = **in
	}
	if in.SmtpsPort != nil {
		in, out := &in.SmtpsPort, &out.SmtpsPort
		*out = new(float64)
		**out = **in
	}
	if in.SmtpsPortAlternative != nil {
		in, out := &in.SmtpsPortAlternative, &out.SmtpsPortAlternative
		*out = new(float64)
		**out = **in
	}
	if in.SpfConfig != nil {
		in, out := &in.SpfConfig, &out.SpfConfig
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainObservation.
func (in *DomainObservation) DeepCopy() *DomainObservation {
	if in == nil {
		return nil
	}
	out := new(DomainObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainParameters) DeepCopyInto(out *DomainParameters) {
	*out = *in
	if in.AcceptTos != nil {
		in, out := &in.AcceptTos, &out.AcceptTos
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainParameters.
func (in *DomainParameters) DeepCopy() *DomainParameters {
	if in == nil {
		return nil
	}
	out := new(DomainParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpec) DeepCopyInto(out *DomainSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpec.
func (in *DomainSpec) DeepCopy() *DomainSpec {
	if in == nil {
		return nil
	}
	out := new(DomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainStatus) DeepCopyInto(out *DomainStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainStatus.
func (in *DomainStatus) DeepCopy() *DomainStatus {
	if in == nil {
		return nil
	}
	out := new(DomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReputationInitParameters) DeepCopyInto(out *ReputationInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReputationInitParameters.
func (in *ReputationInitParameters) DeepCopy() *ReputationInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReputationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReputationObservation) DeepCopyInto(out *ReputationObservation) {
	*out = *in
	if in.PreviousScore != nil {
		in, out := &in.PreviousScore, &out.PreviousScore
		*out = new(float64)
		**out = **in
	}
	if in.PreviousScoredAt != nil {
		in, out := &in.PreviousScoredAt, &out.PreviousScoredAt
		*out = new(string)
		**out = **in
	}
	if in.Score != nil {
		in, out := &in.Score, &out.Score
		*out = new(float64)
		**out = **in
	}
	if in.ScoredAt != nil {
		in, out := &in.ScoredAt, &out.ScoredAt
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReputationObservation.
func (in *ReputationObservation) DeepCopy() *ReputationObservation {
	if in == nil {
		return nil
	}
	out := new(ReputationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReputationParameters) DeepCopyInto(out *ReputationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReputationParameters.
func (in *ReputationParameters) DeepCopy() *ReputationParameters {
	if in == nil {
		return nil
	}
	out := new(ReputationParameters)
	in.DeepCopyInto(out)
	return out
}
