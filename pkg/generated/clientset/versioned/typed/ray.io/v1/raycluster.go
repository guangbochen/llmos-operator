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
	rayv1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// RayClustersGetter has a method to return a RayClusterInterface.
// A group's client should implement this interface.
type RayClustersGetter interface {
	RayClusters() RayClusterInterface
}

// RayClusterInterface has methods to work with RayCluster resources.
type RayClusterInterface interface {
	Create(ctx context.Context, rayCluster *rayv1.RayCluster, opts metav1.CreateOptions) (*rayv1.RayCluster, error)
	Update(ctx context.Context, rayCluster *rayv1.RayCluster, opts metav1.UpdateOptions) (*rayv1.RayCluster, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, rayCluster *rayv1.RayCluster, opts metav1.UpdateOptions) (*rayv1.RayCluster, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*rayv1.RayCluster, error)
	List(ctx context.Context, opts metav1.ListOptions) (*rayv1.RayClusterList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *rayv1.RayCluster, err error)
	RayClusterExpansion
}

// rayClusters implements RayClusterInterface
type rayClusters struct {
	*gentype.ClientWithList[*rayv1.RayCluster, *rayv1.RayClusterList]
}

// newRayClusters returns a RayClusters
func newRayClusters(c *RayV1Client) *rayClusters {
	return &rayClusters{
		gentype.NewClientWithList[*rayv1.RayCluster, *rayv1.RayClusterList](
			"rayclusters",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *rayv1.RayCluster { return &rayv1.RayCluster{} },
			func() *rayv1.RayClusterList { return &rayv1.RayClusterList{} },
		),
	}
}
