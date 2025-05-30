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
	cephrookiov1 "github.com/llmos-ai/llmos-operator/pkg/generated/clientset/versioned/typed/ceph.rook.io/v1"
	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeCephCOSIDrivers implements CephCOSIDriverInterface
type fakeCephCOSIDrivers struct {
	*gentype.FakeClientWithList[*v1.CephCOSIDriver, *v1.CephCOSIDriverList]
	Fake *FakeCephV1
}

func newFakeCephCOSIDrivers(fake *FakeCephV1, namespace string) cephrookiov1.CephCOSIDriverInterface {
	return &fakeCephCOSIDrivers{
		gentype.NewFakeClientWithList[*v1.CephCOSIDriver, *v1.CephCOSIDriverList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("cephcosidrivers"),
			v1.SchemeGroupVersion.WithKind("CephCOSIDriver"),
			func() *v1.CephCOSIDriver { return &v1.CephCOSIDriver{} },
			func() *v1.CephCOSIDriverList { return &v1.CephCOSIDriverList{} },
			func(dst, src *v1.CephCOSIDriverList) { dst.ListMeta = src.ListMeta },
			func(list *v1.CephCOSIDriverList) []*v1.CephCOSIDriver { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.CephCOSIDriverList, items []*v1.CephCOSIDriver) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
