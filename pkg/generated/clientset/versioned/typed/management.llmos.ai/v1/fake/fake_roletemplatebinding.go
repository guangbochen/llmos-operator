/*
Copyright 2024 llmos.ai.

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
// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/llmos-ai/llmos-operator/pkg/apis/management.llmos.ai/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRoleTemplateBindings implements RoleTemplateBindingInterface
type FakeRoleTemplateBindings struct {
	Fake *FakeManagementV1
}

var roletemplatebindingsResource = v1.SchemeGroupVersion.WithResource("roletemplatebindings")

var roletemplatebindingsKind = v1.SchemeGroupVersion.WithKind("RoleTemplateBinding")

// Get takes name of the roleTemplateBinding, and returns the corresponding roleTemplateBinding object, and an error if there is any.
func (c *FakeRoleTemplateBindings) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RoleTemplateBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(roletemplatebindingsResource, name), &v1.RoleTemplateBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RoleTemplateBinding), err
}

// List takes label and field selectors, and returns the list of RoleTemplateBindings that match those selectors.
func (c *FakeRoleTemplateBindings) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RoleTemplateBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(roletemplatebindingsResource, roletemplatebindingsKind, opts), &v1.RoleTemplateBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RoleTemplateBindingList{ListMeta: obj.(*v1.RoleTemplateBindingList).ListMeta}
	for _, item := range obj.(*v1.RoleTemplateBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested roleTemplateBindings.
func (c *FakeRoleTemplateBindings) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(roletemplatebindingsResource, opts))
}

// Create takes the representation of a roleTemplateBinding and creates it.  Returns the server's representation of the roleTemplateBinding, and an error, if there is any.
func (c *FakeRoleTemplateBindings) Create(ctx context.Context, roleTemplateBinding *v1.RoleTemplateBinding, opts metav1.CreateOptions) (result *v1.RoleTemplateBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(roletemplatebindingsResource, roleTemplateBinding), &v1.RoleTemplateBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RoleTemplateBinding), err
}

// Update takes the representation of a roleTemplateBinding and updates it. Returns the server's representation of the roleTemplateBinding, and an error, if there is any.
func (c *FakeRoleTemplateBindings) Update(ctx context.Context, roleTemplateBinding *v1.RoleTemplateBinding, opts metav1.UpdateOptions) (result *v1.RoleTemplateBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(roletemplatebindingsResource, roleTemplateBinding), &v1.RoleTemplateBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RoleTemplateBinding), err
}

// Delete takes name of the roleTemplateBinding and deletes it. Returns an error if one occurs.
func (c *FakeRoleTemplateBindings) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(roletemplatebindingsResource, name, opts), &v1.RoleTemplateBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoleTemplateBindings) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(roletemplatebindingsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RoleTemplateBindingList{})
	return err
}

// Patch applies the patch and returns the patched roleTemplateBinding.
func (c *FakeRoleTemplateBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RoleTemplateBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(roletemplatebindingsResource, name, pt, data, subresources...), &v1.RoleTemplateBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RoleTemplateBinding), err
}