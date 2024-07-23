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

	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCephObjectZoneGroups implements CephObjectZoneGroupInterface
type FakeCephObjectZoneGroups struct {
	Fake *FakeCephV1
	ns   string
}

var cephobjectzonegroupsResource = v1.SchemeGroupVersion.WithResource("cephobjectzonegroups")

var cephobjectzonegroupsKind = v1.SchemeGroupVersion.WithKind("CephObjectZoneGroup")

// Get takes name of the cephObjectZoneGroup, and returns the corresponding cephObjectZoneGroup object, and an error if there is any.
func (c *FakeCephObjectZoneGroups) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CephObjectZoneGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cephobjectzonegroupsResource, c.ns, name), &v1.CephObjectZoneGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephObjectZoneGroup), err
}

// List takes label and field selectors, and returns the list of CephObjectZoneGroups that match those selectors.
func (c *FakeCephObjectZoneGroups) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CephObjectZoneGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cephobjectzonegroupsResource, cephobjectzonegroupsKind, c.ns, opts), &v1.CephObjectZoneGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.CephObjectZoneGroupList{ListMeta: obj.(*v1.CephObjectZoneGroupList).ListMeta}
	for _, item := range obj.(*v1.CephObjectZoneGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cephObjectZoneGroups.
func (c *FakeCephObjectZoneGroups) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cephobjectzonegroupsResource, c.ns, opts))

}

// Create takes the representation of a cephObjectZoneGroup and creates it.  Returns the server's representation of the cephObjectZoneGroup, and an error, if there is any.
func (c *FakeCephObjectZoneGroups) Create(ctx context.Context, cephObjectZoneGroup *v1.CephObjectZoneGroup, opts metav1.CreateOptions) (result *v1.CephObjectZoneGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cephobjectzonegroupsResource, c.ns, cephObjectZoneGroup), &v1.CephObjectZoneGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephObjectZoneGroup), err
}

// Update takes the representation of a cephObjectZoneGroup and updates it. Returns the server's representation of the cephObjectZoneGroup, and an error, if there is any.
func (c *FakeCephObjectZoneGroups) Update(ctx context.Context, cephObjectZoneGroup *v1.CephObjectZoneGroup, opts metav1.UpdateOptions) (result *v1.CephObjectZoneGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cephobjectzonegroupsResource, c.ns, cephObjectZoneGroup), &v1.CephObjectZoneGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephObjectZoneGroup), err
}

// Delete takes name of the cephObjectZoneGroup and deletes it. Returns an error if one occurs.
func (c *FakeCephObjectZoneGroups) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cephobjectzonegroupsResource, c.ns, name, opts), &v1.CephObjectZoneGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCephObjectZoneGroups) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cephobjectzonegroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.CephObjectZoneGroupList{})
	return err
}

// Patch applies the patch and returns the patched cephObjectZoneGroup.
func (c *FakeCephObjectZoneGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CephObjectZoneGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cephobjectzonegroupsResource, c.ns, name, pt, data, subresources...), &v1.CephObjectZoneGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephObjectZoneGroup), err
}