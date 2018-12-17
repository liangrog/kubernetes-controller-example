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

package v1

import (
	"time"

	v1 "github.com/liangrog/kube-crd-example/pkg/apis/foo/v1"
	scheme "github.com/liangrog/kube-crd-example/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HelloTypesGetter has a method to return a HelloTypeInterface.
// A group's client should implement this interface.
type HelloTypesGetter interface {
	HelloTypes(namespace string) HelloTypeInterface
}

// HelloTypeInterface has methods to work with HelloType resources.
type HelloTypeInterface interface {
	Create(*v1.HelloType) (*v1.HelloType, error)
	Update(*v1.HelloType) (*v1.HelloType, error)
	UpdateStatus(*v1.HelloType) (*v1.HelloType, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.HelloType, error)
	List(opts metav1.ListOptions) (*v1.HelloTypeList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.HelloType, err error)
	HelloTypeExpansion
}

// helloTypes implements HelloTypeInterface
type helloTypes struct {
	client rest.Interface
	ns     string
}

// newHelloTypes returns a HelloTypes
func newHelloTypes(c *ExampleV1Client, namespace string) *helloTypes {
	return &helloTypes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the helloType, and returns the corresponding helloType object, and an error if there is any.
func (c *helloTypes) Get(name string, options metav1.GetOptions) (result *v1.HelloType, err error) {
	result = &v1.HelloType{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hellotypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HelloTypes that match those selectors.
func (c *helloTypes) List(opts metav1.ListOptions) (result *v1.HelloTypeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.HelloTypeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hellotypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested helloTypes.
func (c *helloTypes) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hellotypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a helloType and creates it.  Returns the server's representation of the helloType, and an error, if there is any.
func (c *helloTypes) Create(helloType *v1.HelloType) (result *v1.HelloType, err error) {
	result = &v1.HelloType{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hellotypes").
		Body(helloType).
		Do().
		Into(result)
	return
}

// Update takes the representation of a helloType and updates it. Returns the server's representation of the helloType, and an error, if there is any.
func (c *helloTypes) Update(helloType *v1.HelloType) (result *v1.HelloType, err error) {
	result = &v1.HelloType{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hellotypes").
		Name(helloType.Name).
		Body(helloType).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *helloTypes) UpdateStatus(helloType *v1.HelloType) (result *v1.HelloType, err error) {
	result = &v1.HelloType{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hellotypes").
		Name(helloType.Name).
		SubResource("status").
		Body(helloType).
		Do().
		Into(result)
	return
}

// Delete takes name of the helloType and deletes it. Returns an error if one occurs.
func (c *helloTypes) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hellotypes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *helloTypes) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hellotypes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched helloType.
func (c *helloTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.HelloType, err error) {
	result = &v1.HelloType{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hellotypes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
