/*
Copyright The Kube Bind Authors.

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
)

// FakeAPIServiceBindings implements APIServiceBindingInterface
type FakeAPIServiceBindings struct {
	Fake *FakeKubeBindV1alpha1
}

var apiservicebindingsResource = schema.GroupVersionResource{Group: "kube-bind.io", Version: "v1alpha1", Resource: "apiservicebindings"}

var apiservicebindingsKind = schema.GroupVersionKind{Group: "kube-bind.io", Version: "v1alpha1", Kind: "APIServiceBinding"}

// Get takes name of the aPIServiceBinding, and returns the corresponding aPIServiceBinding object, and an error if there is any.
func (c *FakeAPIServiceBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIServiceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(apiservicebindingsResource, name), &v1alpha1.APIServiceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceBinding), err
}

// List takes label and field selectors, and returns the list of APIServiceBindings that match those selectors.
func (c *FakeAPIServiceBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIServiceBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(apiservicebindingsResource, apiservicebindingsKind, opts), &v1alpha1.APIServiceBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.APIServiceBindingList{ListMeta: obj.(*v1alpha1.APIServiceBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.APIServiceBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIServiceBindings.
func (c *FakeAPIServiceBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(apiservicebindingsResource, opts))
}

// Create takes the representation of a aPIServiceBinding and creates it.  Returns the server's representation of the aPIServiceBinding, and an error, if there is any.
func (c *FakeAPIServiceBindings) Create(ctx context.Context, aPIServiceBinding *v1alpha1.APIServiceBinding, opts v1.CreateOptions) (result *v1alpha1.APIServiceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(apiservicebindingsResource, aPIServiceBinding), &v1alpha1.APIServiceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceBinding), err
}

// Update takes the representation of a aPIServiceBinding and updates it. Returns the server's representation of the aPIServiceBinding, and an error, if there is any.
func (c *FakeAPIServiceBindings) Update(ctx context.Context, aPIServiceBinding *v1alpha1.APIServiceBinding, opts v1.UpdateOptions) (result *v1alpha1.APIServiceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(apiservicebindingsResource, aPIServiceBinding), &v1alpha1.APIServiceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPIServiceBindings) UpdateStatus(ctx context.Context, aPIServiceBinding *v1alpha1.APIServiceBinding, opts v1.UpdateOptions) (*v1alpha1.APIServiceBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(apiservicebindingsResource, "status", aPIServiceBinding), &v1alpha1.APIServiceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceBinding), err
}

// Delete takes name of the aPIServiceBinding and deletes it. Returns an error if one occurs.
func (c *FakeAPIServiceBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(apiservicebindingsResource, name, opts), &v1alpha1.APIServiceBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPIServiceBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(apiservicebindingsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.APIServiceBindingList{})
	return err
}

// Patch applies the patch and returns the patched aPIServiceBinding.
func (c *FakeAPIServiceBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIServiceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apiservicebindingsResource, name, pt, data, subresources...), &v1alpha1.APIServiceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceBinding), err
}
