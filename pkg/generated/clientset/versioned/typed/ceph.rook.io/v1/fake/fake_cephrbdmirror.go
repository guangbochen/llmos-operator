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

	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCephRBDMirrors implements CephRBDMirrorInterface
type FakeCephRBDMirrors struct {
	Fake *FakeCephV1
	ns   string
}

var cephrbdmirrorsResource = v1.SchemeGroupVersion.WithResource("cephrbdmirrors")

var cephrbdmirrorsKind = v1.SchemeGroupVersion.WithKind("CephRBDMirror")

// Get takes name of the cephRBDMirror, and returns the corresponding cephRBDMirror object, and an error if there is any.
func (c *FakeCephRBDMirrors) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CephRBDMirror, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cephrbdmirrorsResource, c.ns, name), &v1.CephRBDMirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephRBDMirror), err
}

// List takes label and field selectors, and returns the list of CephRBDMirrors that match those selectors.
func (c *FakeCephRBDMirrors) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CephRBDMirrorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cephrbdmirrorsResource, cephrbdmirrorsKind, c.ns, opts), &v1.CephRBDMirrorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.CephRBDMirrorList{ListMeta: obj.(*v1.CephRBDMirrorList).ListMeta}
	for _, item := range obj.(*v1.CephRBDMirrorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cephRBDMirrors.
func (c *FakeCephRBDMirrors) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cephrbdmirrorsResource, c.ns, opts))

}

// Create takes the representation of a cephRBDMirror and creates it.  Returns the server's representation of the cephRBDMirror, and an error, if there is any.
func (c *FakeCephRBDMirrors) Create(ctx context.Context, cephRBDMirror *v1.CephRBDMirror, opts metav1.CreateOptions) (result *v1.CephRBDMirror, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cephrbdmirrorsResource, c.ns, cephRBDMirror), &v1.CephRBDMirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephRBDMirror), err
}

// Update takes the representation of a cephRBDMirror and updates it. Returns the server's representation of the cephRBDMirror, and an error, if there is any.
func (c *FakeCephRBDMirrors) Update(ctx context.Context, cephRBDMirror *v1.CephRBDMirror, opts metav1.UpdateOptions) (result *v1.CephRBDMirror, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cephrbdmirrorsResource, c.ns, cephRBDMirror), &v1.CephRBDMirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephRBDMirror), err
}

// Delete takes name of the cephRBDMirror and deletes it. Returns an error if one occurs.
func (c *FakeCephRBDMirrors) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cephrbdmirrorsResource, c.ns, name, opts), &v1.CephRBDMirror{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCephRBDMirrors) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cephrbdmirrorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.CephRBDMirrorList{})
	return err
}

// Patch applies the patch and returns the patched cephRBDMirror.
func (c *FakeCephRBDMirrors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CephRBDMirror, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cephrbdmirrorsResource, c.ns, name, pt, data, subresources...), &v1.CephRBDMirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephRBDMirror), err
}
