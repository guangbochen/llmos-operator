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
	v1 "github.com/kubernetes-csi/external-snapshotter/client/v8/apis/volumesnapshot/v1"
	snapshotstoragek8siov1 "github.com/llmos-ai/llmos-operator/pkg/generated/clientset/versioned/typed/snapshot.storage.k8s.io/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeVolumeSnapshotContents implements VolumeSnapshotContentInterface
type fakeVolumeSnapshotContents struct {
	*gentype.FakeClientWithList[*v1.VolumeSnapshotContent, *v1.VolumeSnapshotContentList]
	Fake *FakeSnapshotV1
}

func newFakeVolumeSnapshotContents(fake *FakeSnapshotV1) snapshotstoragek8siov1.VolumeSnapshotContentInterface {
	return &fakeVolumeSnapshotContents{
		gentype.NewFakeClientWithList[*v1.VolumeSnapshotContent, *v1.VolumeSnapshotContentList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("volumesnapshotcontents"),
			v1.SchemeGroupVersion.WithKind("VolumeSnapshotContent"),
			func() *v1.VolumeSnapshotContent { return &v1.VolumeSnapshotContent{} },
			func() *v1.VolumeSnapshotContentList { return &v1.VolumeSnapshotContentList{} },
			func(dst, src *v1.VolumeSnapshotContentList) { dst.ListMeta = src.ListMeta },
			func(list *v1.VolumeSnapshotContentList) []*v1.VolumeSnapshotContent {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1.VolumeSnapshotContentList, items []*v1.VolumeSnapshotContent) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
