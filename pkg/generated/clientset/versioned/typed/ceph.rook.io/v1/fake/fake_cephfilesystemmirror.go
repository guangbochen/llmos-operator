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

// FakeCephFilesystemMirrors implements CephFilesystemMirrorInterface
type FakeCephFilesystemMirrors struct {
	Fake *FakeCephV1
	ns   string
}

var cephfilesystemmirrorsResource = v1.SchemeGroupVersion.WithResource("cephfilesystemmirrors")

var cephfilesystemmirrorsKind = v1.SchemeGroupVersion.WithKind("CephFilesystemMirror")

// Get takes name of the cephFilesystemMirror, and returns the corresponding cephFilesystemMirror object, and an error if there is any.
func (c *FakeCephFilesystemMirrors) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CephFilesystemMirror, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cephfilesystemmirrorsResource, c.ns, name), &v1.CephFilesystemMirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephFilesystemMirror), err
}

// List takes label and field selectors, and returns the list of CephFilesystemMirrors that match those selectors.
func (c *FakeCephFilesystemMirrors) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CephFilesystemMirrorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cephfilesystemmirrorsResource, cephfilesystemmirrorsKind, c.ns, opts), &v1.CephFilesystemMirrorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.CephFilesystemMirrorList{ListMeta: obj.(*v1.CephFilesystemMirrorList).ListMeta}
	for _, item := range obj.(*v1.CephFilesystemMirrorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cephFilesystemMirrors.
func (c *FakeCephFilesystemMirrors) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cephfilesystemmirrorsResource, c.ns, opts))

}

// Create takes the representation of a cephFilesystemMirror and creates it.  Returns the server's representation of the cephFilesystemMirror, and an error, if there is any.
func (c *FakeCephFilesystemMirrors) Create(ctx context.Context, cephFilesystemMirror *v1.CephFilesystemMirror, opts metav1.CreateOptions) (result *v1.CephFilesystemMirror, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cephfilesystemmirrorsResource, c.ns, cephFilesystemMirror), &v1.CephFilesystemMirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephFilesystemMirror), err
}

// Update takes the representation of a cephFilesystemMirror and updates it. Returns the server's representation of the cephFilesystemMirror, and an error, if there is any.
func (c *FakeCephFilesystemMirrors) Update(ctx context.Context, cephFilesystemMirror *v1.CephFilesystemMirror, opts metav1.UpdateOptions) (result *v1.CephFilesystemMirror, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cephfilesystemmirrorsResource, c.ns, cephFilesystemMirror), &v1.CephFilesystemMirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephFilesystemMirror), err
}

// Delete takes name of the cephFilesystemMirror and deletes it. Returns an error if one occurs.
func (c *FakeCephFilesystemMirrors) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cephfilesystemmirrorsResource, c.ns, name, opts), &v1.CephFilesystemMirror{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCephFilesystemMirrors) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cephfilesystemmirrorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.CephFilesystemMirrorList{})
	return err
}

// Patch applies the patch and returns the patched cephFilesystemMirror.
func (c *FakeCephFilesystemMirrors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CephFilesystemMirror, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cephfilesystemmirrorsResource, c.ns, name, pt, data, subresources...), &v1.CephFilesystemMirror{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephFilesystemMirror), err
}
