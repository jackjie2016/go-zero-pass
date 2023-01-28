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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1beta1 "k8s.io/api/authorization/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// SelfSubjectAccessReviewsGetter has a method to return a SelfSubjectAccessReviewInterface.
// A group's client should implement this interface.
type SelfSubjectAccessReviewsGetter interface {
	SelfSubjectAccessReviews() SelfSubjectAccessReviewInterface
}

// SelfSubjectAccessReviewInterface has methods to work with SelfSubjectAccessReview resources.
type SelfSubjectAccessReviewInterface interface {
	Create(ctx context.Context, selfSubjectAccessReview *v1beta1.SelfSubjectAccessReview, opts v1.CreateOptions) (*v1beta1.SelfSubjectAccessReview, error)
	SelfSubjectAccessReviewExpansion
}

// selfSubjectAccessReviews implements SelfSubjectAccessReviewInterface
type selfSubjectAccessReviews struct {
	client rest.Interface
}

// newSelfSubjectAccessReviews returns a SelfSubjectAccessReviews
func newSelfSubjectAccessReviews(c *AuthorizationV1beta1Client) *selfSubjectAccessReviews {
	return &selfSubjectAccessReviews{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a selfSubjectAccessReview and creates it.  Returns the server's representation of the selfSubjectAccessReview, and an error, if there is any.
func (c *selfSubjectAccessReviews) Create(ctx context.Context, selfSubjectAccessReview *v1beta1.SelfSubjectAccessReview, opts v1.CreateOptions) (result *v1beta1.SelfSubjectAccessReview, err error) {
	result = &v1beta1.SelfSubjectAccessReview{}
	err = c.client.Post().
		Resource("selfsubjectaccessreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selfSubjectAccessReview).
		Do(ctx).
		Into(result)
	return
}
