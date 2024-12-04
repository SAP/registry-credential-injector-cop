/*
SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and registry-credential-injector-cop contributors
SPDX-License-Identifier: Apache-2.0
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	"github.com/sap/component-operator-runtime/pkg/component"
	componentoperatorruntimetypes "github.com/sap/component-operator-runtime/pkg/types"
)

// RegistryCredentialInjectorSpec defines the desired state of RegistryCredentialInjector.
type RegistryCredentialInjectorSpec struct {
	component.Spec `json:",inline"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=1
	ReplicaCount int `json:"replicaCount,omitempty"`
	// +optional
	Image                          component.ImageSpec `json:"image"`
	component.KubernetesProperties `json:",inline"`
	ObjectSelector                 *metav1.LabelSelector `json:"objectSelector,omitempty"`
	NamespaceSelector              *metav1.LabelSelector `json:"namespaceSelector,omitempty"`
	DefaultPullSecret              string                `json:"defaultPullSecret,omitempty"`
	LogLevel                       int                   `json:"logLevel,omitempty"`
}

// RegistryCredentialInjectorStatus defines the observed state of RegistryCredentialInjector.
type RegistryCredentialInjectorStatus struct {
	component.Status `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +genclient

// RegistryCredentialInjector is the Schema for the registrycredentialinjectors API.
type RegistryCredentialInjector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RegistryCredentialInjectorSpec `json:"spec,omitempty"`
	// +kubebuilder:default={"observedGeneration":-1}
	Status RegistryCredentialInjectorStatus `json:"status,omitempty"`
}

var _ component.Component = &RegistryCredentialInjector{}

// +kubebuilder:object:root=true

// RegistryCredentialInjectorList contains a list of RegistryCredentialInjector.
type RegistryCredentialInjectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryCredentialInjector `json:"items"`
}

func (s *RegistryCredentialInjectorSpec) ToUnstructured() map[string]any {
	result, err := runtime.DefaultUnstructuredConverter.ToUnstructured(s)
	if err != nil {
		panic(err)
	}
	return result
}

func (c *RegistryCredentialInjector) GetDeploymentNamespace() string {
	if c.Spec.Namespace != "" {
		return c.Spec.Namespace
	}
	return c.Namespace
}

func (c *RegistryCredentialInjector) GetDeploymentName() string {
	if c.Spec.Name != "" {
		return c.Spec.Name
	}
	return c.Name
}

func (c *RegistryCredentialInjector) GetSpec() componentoperatorruntimetypes.Unstructurable {
	return &c.Spec
}

func (c *RegistryCredentialInjector) GetStatus() *component.Status {
	return &c.Status.Status
}

func init() {
	SchemeBuilder.Register(&RegistryCredentialInjector{}, &RegistryCredentialInjectorList{})
}
