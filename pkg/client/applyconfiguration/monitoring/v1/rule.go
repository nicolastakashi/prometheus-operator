// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// RuleApplyConfiguration represents a declarative configuration of the Rule type for use
// with apply.
type RuleApplyConfiguration struct {
	Record        *string                        `json:"record,omitempty"`
	Alert         *string                        `json:"alert,omitempty"`
	Expr          *intstr.IntOrString            `json:"expr,omitempty"`
	For           *monitoringv1.Duration         `json:"for,omitempty"`
	KeepFiringFor *monitoringv1.NonEmptyDuration `json:"keep_firing_for,omitempty"`
	Labels        map[string]string              `json:"labels,omitempty"`
	Annotations   map[string]string              `json:"annotations,omitempty"`
}

// RuleApplyConfiguration constructs a declarative configuration of the Rule type for use with
// apply.
func Rule() *RuleApplyConfiguration {
	return &RuleApplyConfiguration{}
}

// WithRecord sets the Record field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Record field is set to the value of the last call.
func (b *RuleApplyConfiguration) WithRecord(value string) *RuleApplyConfiguration {
	b.Record = &value
	return b
}

// WithAlert sets the Alert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Alert field is set to the value of the last call.
func (b *RuleApplyConfiguration) WithAlert(value string) *RuleApplyConfiguration {
	b.Alert = &value
	return b
}

// WithExpr sets the Expr field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Expr field is set to the value of the last call.
func (b *RuleApplyConfiguration) WithExpr(value intstr.IntOrString) *RuleApplyConfiguration {
	b.Expr = &value
	return b
}

// WithFor sets the For field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the For field is set to the value of the last call.
func (b *RuleApplyConfiguration) WithFor(value monitoringv1.Duration) *RuleApplyConfiguration {
	b.For = &value
	return b
}

// WithKeepFiringFor sets the KeepFiringFor field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeepFiringFor field is set to the value of the last call.
func (b *RuleApplyConfiguration) WithKeepFiringFor(value monitoringv1.NonEmptyDuration) *RuleApplyConfiguration {
	b.KeepFiringFor = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *RuleApplyConfiguration) WithLabels(entries map[string]string) *RuleApplyConfiguration {
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *RuleApplyConfiguration) WithAnnotations(entries map[string]string) *RuleApplyConfiguration {
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}
