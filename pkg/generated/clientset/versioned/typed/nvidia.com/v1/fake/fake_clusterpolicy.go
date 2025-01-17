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

package fake

import (
	"context"

	v1 "github.com/NVIDIA/gpu-operator/api/nvidia/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterPolicies implements ClusterPolicyInterface
type FakeClusterPolicies struct {
	Fake *FakeNvidiaV1
}

var clusterpoliciesResource = v1.SchemeGroupVersion.WithResource("clusterpolicies")

var clusterpoliciesKind = v1.SchemeGroupVersion.WithKind("ClusterPolicy")

// Get takes name of the clusterPolicy, and returns the corresponding clusterPolicy object, and an error if there is any.
func (c *FakeClusterPolicies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterpoliciesResource, name), &v1.ClusterPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterPolicy), err
}

// List takes label and field selectors, and returns the list of ClusterPolicies that match those selectors.
func (c *FakeClusterPolicies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterpoliciesResource, clusterpoliciesKind, opts), &v1.ClusterPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ClusterPolicyList{ListMeta: obj.(*v1.ClusterPolicyList).ListMeta}
	for _, item := range obj.(*v1.ClusterPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterPolicies.
func (c *FakeClusterPolicies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterpoliciesResource, opts))
}

// Create takes the representation of a clusterPolicy and creates it.  Returns the server's representation of the clusterPolicy, and an error, if there is any.
func (c *FakeClusterPolicies) Create(ctx context.Context, clusterPolicy *v1.ClusterPolicy, opts metav1.CreateOptions) (result *v1.ClusterPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterpoliciesResource, clusterPolicy), &v1.ClusterPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterPolicy), err
}

// Update takes the representation of a clusterPolicy and updates it. Returns the server's representation of the clusterPolicy, and an error, if there is any.
func (c *FakeClusterPolicies) Update(ctx context.Context, clusterPolicy *v1.ClusterPolicy, opts metav1.UpdateOptions) (result *v1.ClusterPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterpoliciesResource, clusterPolicy), &v1.ClusterPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterPolicies) UpdateStatus(ctx context.Context, clusterPolicy *v1.ClusterPolicy, opts metav1.UpdateOptions) (*v1.ClusterPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterpoliciesResource, "status", clusterPolicy), &v1.ClusterPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterPolicy), err
}

// Delete takes name of the clusterPolicy and deletes it. Returns an error if one occurs.
func (c *FakeClusterPolicies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterpoliciesResource, name, opts), &v1.ClusterPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterPolicies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterpoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ClusterPolicyList{})
	return err
}

// Patch applies the patch and returns the patched clusterPolicy.
func (c *FakeClusterPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterpoliciesResource, name, pt, data, subresources...), &v1.ClusterPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterPolicy), err
}
