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

package v1alpha1

import (
	monitoringv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
)

// KubernetesSDConfigApplyConfiguration represents a declarative configuration of the KubernetesSDConfig type for use
// with apply.
type KubernetesSDConfigApplyConfiguration struct {
	APIServer                        *string                                 `json:"apiServer,omitempty"`
	Role                             *monitoringv1alpha1.KubernetesRole      `json:"role,omitempty"`
	Namespaces                       *NamespaceDiscoveryApplyConfiguration   `json:"namespaces,omitempty"`
	AttachMetadata                   *AttachMetadataApplyConfiguration       `json:"attachMetadata,omitempty"`
	Selectors                        []K8SSelectorConfigApplyConfiguration   `json:"selectors,omitempty"`
	BasicAuth                        *v1.BasicAuthApplyConfiguration         `json:"basicAuth,omitempty"`
	Authorization                    *v1.SafeAuthorizationApplyConfiguration `json:"authorization,omitempty"`
	OAuth2                           *v1.OAuth2ApplyConfiguration            `json:"oauth2,omitempty"`
	v1.ProxyConfigApplyConfiguration `json:",inline"`
	FollowRedirects                  *bool                               `json:"followRedirects,omitempty"`
	EnableHTTP2                      *bool                               `json:"enableHTTP2,omitempty"`
	TLSConfig                        *v1.SafeTLSConfigApplyConfiguration `json:"tlsConfig,omitempty"`
}

// KubernetesSDConfigApplyConfiguration constructs a declarative configuration of the KubernetesSDConfig type for use with
// apply.
func KubernetesSDConfig() *KubernetesSDConfigApplyConfiguration {
	return &KubernetesSDConfigApplyConfiguration{}
}

// WithAPIServer sets the APIServer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIServer field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithAPIServer(value string) *KubernetesSDConfigApplyConfiguration {
	b.APIServer = &value
	return b
}

// WithRole sets the Role field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Role field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithRole(value monitoringv1alpha1.KubernetesRole) *KubernetesSDConfigApplyConfiguration {
	b.Role = &value
	return b
}

// WithNamespaces sets the Namespaces field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespaces field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithNamespaces(value *NamespaceDiscoveryApplyConfiguration) *KubernetesSDConfigApplyConfiguration {
	b.Namespaces = value
	return b
}

// WithAttachMetadata sets the AttachMetadata field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AttachMetadata field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithAttachMetadata(value *AttachMetadataApplyConfiguration) *KubernetesSDConfigApplyConfiguration {
	b.AttachMetadata = value
	return b
}

// WithSelectors adds the given value to the Selectors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Selectors field.
func (b *KubernetesSDConfigApplyConfiguration) WithSelectors(values ...*K8SSelectorConfigApplyConfiguration) *KubernetesSDConfigApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSelectors")
		}
		b.Selectors = append(b.Selectors, *values[i])
	}
	return b
}

// WithBasicAuth sets the BasicAuth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BasicAuth field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithBasicAuth(value *v1.BasicAuthApplyConfiguration) *KubernetesSDConfigApplyConfiguration {
	b.BasicAuth = value
	return b
}

// WithAuthorization sets the Authorization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Authorization field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithAuthorization(value *v1.SafeAuthorizationApplyConfiguration) *KubernetesSDConfigApplyConfiguration {
	b.Authorization = value
	return b
}

// WithOAuth2 sets the OAuth2 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OAuth2 field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithOAuth2(value *v1.OAuth2ApplyConfiguration) *KubernetesSDConfigApplyConfiguration {
	b.OAuth2 = value
	return b
}

// WithProxyURL sets the ProxyURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProxyURL field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithProxyURL(value string) *KubernetesSDConfigApplyConfiguration {
	b.ProxyConfigApplyConfiguration.ProxyURL = &value
	return b
}

// WithNoProxy sets the NoProxy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NoProxy field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithNoProxy(value string) *KubernetesSDConfigApplyConfiguration {
	b.ProxyConfigApplyConfiguration.NoProxy = &value
	return b
}

// WithProxyFromEnvironment sets the ProxyFromEnvironment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProxyFromEnvironment field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithProxyFromEnvironment(value bool) *KubernetesSDConfigApplyConfiguration {
	b.ProxyConfigApplyConfiguration.ProxyFromEnvironment = &value
	return b
}

// WithProxyConnectHeader puts the entries into the ProxyConnectHeader field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ProxyConnectHeader field,
// overwriting an existing map entries in ProxyConnectHeader field with the same key.
func (b *KubernetesSDConfigApplyConfiguration) WithProxyConnectHeader(entries map[string][]corev1.SecretKeySelector) *KubernetesSDConfigApplyConfiguration {
	if b.ProxyConfigApplyConfiguration.ProxyConnectHeader == nil && len(entries) > 0 {
		b.ProxyConfigApplyConfiguration.ProxyConnectHeader = make(map[string][]corev1.SecretKeySelector, len(entries))
	}
	for k, v := range entries {
		b.ProxyConfigApplyConfiguration.ProxyConnectHeader[k] = v
	}
	return b
}

// WithFollowRedirects sets the FollowRedirects field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FollowRedirects field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithFollowRedirects(value bool) *KubernetesSDConfigApplyConfiguration {
	b.FollowRedirects = &value
	return b
}

// WithEnableHTTP2 sets the EnableHTTP2 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableHTTP2 field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithEnableHTTP2(value bool) *KubernetesSDConfigApplyConfiguration {
	b.EnableHTTP2 = &value
	return b
}

// WithTLSConfig sets the TLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSConfig field is set to the value of the last call.
func (b *KubernetesSDConfigApplyConfiguration) WithTLSConfig(value *v1.SafeTLSConfigApplyConfiguration) *KubernetesSDConfigApplyConfiguration {
	b.TLSConfig = value
	return b
}
