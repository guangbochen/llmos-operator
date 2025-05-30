/*
Copyright 2025 llmos.ai.

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

package v1

import (
	context "context"

	managementllmosaiv1 "github.com/llmos-ai/llmos-operator/pkg/apis/management.llmos.ai/v1"
	scheme "github.com/llmos-ai/llmos-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// RoleTemplateBindingsGetter has a method to return a RoleTemplateBindingInterface.
// A group's client should implement this interface.
type RoleTemplateBindingsGetter interface {
	RoleTemplateBindings() RoleTemplateBindingInterface
}

// RoleTemplateBindingInterface has methods to work with RoleTemplateBinding resources.
type RoleTemplateBindingInterface interface {
	Create(ctx context.Context, roleTemplateBinding *managementllmosaiv1.RoleTemplateBinding, opts metav1.CreateOptions) (*managementllmosaiv1.RoleTemplateBinding, error)
	Update(ctx context.Context, roleTemplateBinding *managementllmosaiv1.RoleTemplateBinding, opts metav1.UpdateOptions) (*managementllmosaiv1.RoleTemplateBinding, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*managementllmosaiv1.RoleTemplateBinding, error)
	List(ctx context.Context, opts metav1.ListOptions) (*managementllmosaiv1.RoleTemplateBindingList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *managementllmosaiv1.RoleTemplateBinding, err error)
	RoleTemplateBindingExpansion
}

// roleTemplateBindings implements RoleTemplateBindingInterface
type roleTemplateBindings struct {
	*gentype.ClientWithList[*managementllmosaiv1.RoleTemplateBinding, *managementllmosaiv1.RoleTemplateBindingList]
}

// newRoleTemplateBindings returns a RoleTemplateBindings
func newRoleTemplateBindings(c *ManagementV1Client) *roleTemplateBindings {
	return &roleTemplateBindings{
		gentype.NewClientWithList[*managementllmosaiv1.RoleTemplateBinding, *managementllmosaiv1.RoleTemplateBindingList](
			"roletemplatebindings",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *managementllmosaiv1.RoleTemplateBinding { return &managementllmosaiv1.RoleTemplateBinding{} },
			func() *managementllmosaiv1.RoleTemplateBindingList {
				return &managementllmosaiv1.RoleTemplateBindingList{}
			},
		),
	}
}
