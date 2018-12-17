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

package fake

import (
	foov1 "github.com/liangrog/kube-crd-example/pkg/apis/foo/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHelloTypes implements HelloTypeInterface
type FakeHelloTypes struct {
	Fake *FakeExampleV1
	ns   string
}

var hellotypesResource = schema.GroupVersionResource{Group: "example.com", Version: "v1", Resource: "hellotypes"}

var hellotypesKind = schema.GroupVersionKind{Group: "example.com", Version: "v1", Kind: "HelloType"}

// Get takes name of the helloType, and returns the corresponding helloType object, and an error if there is any.
func (c *FakeHelloTypes) Get(name string, options v1.GetOptions) (result *foov1.HelloType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hellotypesResource, c.ns, name), &foov1.HelloType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*foov1.HelloType), err
}

// List takes label and field selectors, and returns the list of HelloTypes that match those selectors.
func (c *FakeHelloTypes) List(opts v1.ListOptions) (result *foov1.HelloTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hellotypesResource, hellotypesKind, c.ns, opts), &foov1.HelloTypeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &foov1.HelloTypeList{ListMeta: obj.(*foov1.HelloTypeList).ListMeta}
	for _, item := range obj.(*foov1.HelloTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested helloTypes.
func (c *FakeHelloTypes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hellotypesResource, c.ns, opts))

}

// Create takes the representation of a helloType and creates it.  Returns the server's representation of the helloType, and an error, if there is any.
func (c *FakeHelloTypes) Create(helloType *foov1.HelloType) (result *foov1.HelloType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hellotypesResource, c.ns, helloType), &foov1.HelloType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*foov1.HelloType), err
}

// Update takes the representation of a helloType and updates it. Returns the server's representation of the helloType, and an error, if there is any.
func (c *FakeHelloTypes) Update(helloType *foov1.HelloType) (result *foov1.HelloType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hellotypesResource, c.ns, helloType), &foov1.HelloType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*foov1.HelloType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHelloTypes) UpdateStatus(helloType *foov1.HelloType) (*foov1.HelloType, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hellotypesResource, "status", c.ns, helloType), &foov1.HelloType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*foov1.HelloType), err
}

// Delete takes name of the helloType and deletes it. Returns an error if one occurs.
func (c *FakeHelloTypes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hellotypesResource, c.ns, name), &foov1.HelloType{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHelloTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hellotypesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &foov1.HelloTypeList{})
	return err
}

// Patch applies the patch and returns the patched helloType.
func (c *FakeHelloTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *foov1.HelloType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hellotypesResource, c.ns, name, pt, data, subresources...), &foov1.HelloType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*foov1.HelloType), err
}
