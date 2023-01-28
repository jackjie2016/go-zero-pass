/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ServerStorageVersionApplyConfiguration represents an declarative configuration of the ServerStorageVersion type for use
// with apply.
type ServerStorageVersionApplyConfiguration struct {
	APIServerID       *string  `json:"apiServerID,omitempty"`
	EncodingVersion   *string  `json:"encodingVersion,omitempty"`
	DecodableVersions []string `json:"decodableVersions,omitempty"`
}

// ServerStorageVersionApplyConfiguration constructs an declarative configuration of the ServerStorageVersion type for use with
// apply.
func ServerStorageVersion() *ServerStorageVersionApplyConfiguration {
	return &ServerStorageVersionApplyConfiguration{}
}

// WithAPIServerID sets the APIServerID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIServerID field is set to the value of the last call.
func (b *ServerStorageVersionApplyConfiguration) WithAPIServerID(value string) *ServerStorageVersionApplyConfiguration {
	b.APIServerID = &value
	return b
}

// WithEncodingVersion sets the EncodingVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EncodingVersion field is set to the value of the last call.
func (b *ServerStorageVersionApplyConfiguration) WithEncodingVersion(value string) *ServerStorageVersionApplyConfiguration {
	b.EncodingVersion = &value
	return b
}

// WithDecodableVersions adds the given value to the DecodableVersions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the DecodableVersions field.
func (b *ServerStorageVersionApplyConfiguration) WithDecodableVersions(values ...string) *ServerStorageVersionApplyConfiguration {
	for i := range values {
		b.DecodableVersions = append(b.DecodableVersions, values[i])
	}
	return b
}
