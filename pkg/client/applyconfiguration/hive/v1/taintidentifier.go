// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// TaintIdentifierApplyConfiguration represents an declarative configuration of the TaintIdentifier type for use
// with apply.
type TaintIdentifierApplyConfiguration struct {
	Key    *string         `json:"key,omitempty"`
	Effect *v1.TaintEffect `json:"effect,omitempty"`
}

// TaintIdentifierApplyConfiguration constructs an declarative configuration of the TaintIdentifier type for use with
// apply.
func TaintIdentifier() *TaintIdentifierApplyConfiguration {
	return &TaintIdentifierApplyConfiguration{}
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *TaintIdentifierApplyConfiguration) WithKey(value string) *TaintIdentifierApplyConfiguration {
	b.Key = &value
	return b
}

// WithEffect sets the Effect field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Effect field is set to the value of the last call.
func (b *TaintIdentifierApplyConfiguration) WithEffect(value v1.TaintEffect) *TaintIdentifierApplyConfiguration {
	b.Effect = &value
	return b
}