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

	scheme "github.com/llmos-ai/llmos-operator/pkg/generated/clientset/versioned/scheme"
	cephrookiov1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// CephCOSIDriversGetter has a method to return a CephCOSIDriverInterface.
// A group's client should implement this interface.
type CephCOSIDriversGetter interface {
	CephCOSIDrivers(namespace string) CephCOSIDriverInterface
}

// CephCOSIDriverInterface has methods to work with CephCOSIDriver resources.
type CephCOSIDriverInterface interface {
	Create(ctx context.Context, cephCOSIDriver *cephrookiov1.CephCOSIDriver, opts metav1.CreateOptions) (*cephrookiov1.CephCOSIDriver, error)
	Update(ctx context.Context, cephCOSIDriver *cephrookiov1.CephCOSIDriver, opts metav1.UpdateOptions) (*cephrookiov1.CephCOSIDriver, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*cephrookiov1.CephCOSIDriver, error)
	List(ctx context.Context, opts metav1.ListOptions) (*cephrookiov1.CephCOSIDriverList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *cephrookiov1.CephCOSIDriver, err error)
	CephCOSIDriverExpansion
}

// cephCOSIDrivers implements CephCOSIDriverInterface
type cephCOSIDrivers struct {
	*gentype.ClientWithList[*cephrookiov1.CephCOSIDriver, *cephrookiov1.CephCOSIDriverList]
}

// newCephCOSIDrivers returns a CephCOSIDrivers
func newCephCOSIDrivers(c *CephV1Client, namespace string) *cephCOSIDrivers {
	return &cephCOSIDrivers{
		gentype.NewClientWithList[*cephrookiov1.CephCOSIDriver, *cephrookiov1.CephCOSIDriverList](
			"cephcosidrivers",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *cephrookiov1.CephCOSIDriver { return &cephrookiov1.CephCOSIDriver{} },
			func() *cephrookiov1.CephCOSIDriverList { return &cephrookiov1.CephCOSIDriverList{} },
		),
	}
}
