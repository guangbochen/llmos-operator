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

	v1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRayClusters implements RayClusterInterface
type FakeRayClusters struct {
	Fake *FakeRayV1
}

var rayclustersResource = v1.SchemeGroupVersion.WithResource("rayclusters")

var rayclustersKind = v1.SchemeGroupVersion.WithKind("RayCluster")

// Get takes name of the rayCluster, and returns the corresponding rayCluster object, and an error if there is any.
func (c *FakeRayClusters) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RayCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(rayclustersResource, name), &v1.RayCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayCluster), err
}

// List takes label and field selectors, and returns the list of RayClusters that match those selectors.
func (c *FakeRayClusters) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RayClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(rayclustersResource, rayclustersKind, opts), &v1.RayClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RayClusterList{ListMeta: obj.(*v1.RayClusterList).ListMeta}
	for _, item := range obj.(*v1.RayClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rayClusters.
func (c *FakeRayClusters) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(rayclustersResource, opts))
}

// Create takes the representation of a rayCluster and creates it.  Returns the server's representation of the rayCluster, and an error, if there is any.
func (c *FakeRayClusters) Create(ctx context.Context, rayCluster *v1.RayCluster, opts metav1.CreateOptions) (result *v1.RayCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(rayclustersResource, rayCluster), &v1.RayCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayCluster), err
}

// Update takes the representation of a rayCluster and updates it. Returns the server's representation of the rayCluster, and an error, if there is any.
func (c *FakeRayClusters) Update(ctx context.Context, rayCluster *v1.RayCluster, opts metav1.UpdateOptions) (result *v1.RayCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(rayclustersResource, rayCluster), &v1.RayCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRayClusters) UpdateStatus(ctx context.Context, rayCluster *v1.RayCluster, opts metav1.UpdateOptions) (*v1.RayCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(rayclustersResource, "status", rayCluster), &v1.RayCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayCluster), err
}

// Delete takes name of the rayCluster and deletes it. Returns an error if one occurs.
func (c *FakeRayClusters) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(rayclustersResource, name, opts), &v1.RayCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRayClusters) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(rayclustersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RayClusterList{})
	return err
}

// Patch applies the patch and returns the patched rayCluster.
func (c *FakeRayClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RayCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(rayclustersResource, name, pt, data, subresources...), &v1.RayCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayCluster), err
}
