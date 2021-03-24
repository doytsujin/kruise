/*
Copyright 2021 The Kruise Authors.

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

package fake

import (
	v1alpha1 "github.com/openkruise/kruise/apis/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeContainerRecreateRequests implements ContainerRecreateRequestInterface
type FakeContainerRecreateRequests struct {
	Fake *FakeAppsV1alpha1
	ns   string
}

var containerrecreaterequestsResource = schema.GroupVersionResource{Group: "apps.kruise.io", Version: "v1alpha1", Resource: "containerrecreaterequests"}

var containerrecreaterequestsKind = schema.GroupVersionKind{Group: "apps.kruise.io", Version: "v1alpha1", Kind: "ContainerRecreateRequest"}

// Get takes name of the containerRecreateRequest, and returns the corresponding containerRecreateRequest object, and an error if there is any.
func (c *FakeContainerRecreateRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.ContainerRecreateRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(containerrecreaterequestsResource, c.ns, name), &v1alpha1.ContainerRecreateRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerRecreateRequest), err
}

// List takes label and field selectors, and returns the list of ContainerRecreateRequests that match those selectors.
func (c *FakeContainerRecreateRequests) List(opts v1.ListOptions) (result *v1alpha1.ContainerRecreateRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(containerrecreaterequestsResource, containerrecreaterequestsKind, c.ns, opts), &v1alpha1.ContainerRecreateRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ContainerRecreateRequestList{ListMeta: obj.(*v1alpha1.ContainerRecreateRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.ContainerRecreateRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested containerRecreateRequests.
func (c *FakeContainerRecreateRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(containerrecreaterequestsResource, c.ns, opts))

}

// Create takes the representation of a containerRecreateRequest and creates it.  Returns the server's representation of the containerRecreateRequest, and an error, if there is any.
func (c *FakeContainerRecreateRequests) Create(containerRecreateRequest *v1alpha1.ContainerRecreateRequest) (result *v1alpha1.ContainerRecreateRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(containerrecreaterequestsResource, c.ns, containerRecreateRequest), &v1alpha1.ContainerRecreateRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerRecreateRequest), err
}

// Update takes the representation of a containerRecreateRequest and updates it. Returns the server's representation of the containerRecreateRequest, and an error, if there is any.
func (c *FakeContainerRecreateRequests) Update(containerRecreateRequest *v1alpha1.ContainerRecreateRequest) (result *v1alpha1.ContainerRecreateRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(containerrecreaterequestsResource, c.ns, containerRecreateRequest), &v1alpha1.ContainerRecreateRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerRecreateRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeContainerRecreateRequests) UpdateStatus(containerRecreateRequest *v1alpha1.ContainerRecreateRequest) (*v1alpha1.ContainerRecreateRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(containerrecreaterequestsResource, "status", c.ns, containerRecreateRequest), &v1alpha1.ContainerRecreateRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerRecreateRequest), err
}

// Delete takes name of the containerRecreateRequest and deletes it. Returns an error if one occurs.
func (c *FakeContainerRecreateRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(containerrecreaterequestsResource, c.ns, name), &v1alpha1.ContainerRecreateRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeContainerRecreateRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(containerrecreaterequestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ContainerRecreateRequestList{})
	return err
}

// Patch applies the patch and returns the patched containerRecreateRequest.
func (c *FakeContainerRecreateRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ContainerRecreateRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(containerrecreaterequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ContainerRecreateRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ContainerRecreateRequest), err
}
