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
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// VolumeAttachmentsGetter has a method to return a VolumeAttachmentInterface.
// A group's client should implement this interface.
type VolumeAttachmentsGetter interface {
	VolumeAttachments() VolumeAttachmentInterface
}

// VolumeAttachmentInterface has methods to work with VolumeAttachment resources.
type VolumeAttachmentInterface interface {
	Create(ctx context.Context, volumeAttachment *storagev1.VolumeAttachment, opts metav1.CreateOptions) (*storagev1.VolumeAttachment, error)
	Update(ctx context.Context, volumeAttachment *storagev1.VolumeAttachment, opts metav1.UpdateOptions) (*storagev1.VolumeAttachment, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, volumeAttachment *storagev1.VolumeAttachment, opts metav1.UpdateOptions) (*storagev1.VolumeAttachment, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*storagev1.VolumeAttachment, error)
	List(ctx context.Context, opts metav1.ListOptions) (*storagev1.VolumeAttachmentList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *storagev1.VolumeAttachment, err error)
	VolumeAttachmentExpansion
}

// volumeAttachments implements VolumeAttachmentInterface
type volumeAttachments struct {
	*gentype.ClientWithList[*storagev1.VolumeAttachment, *storagev1.VolumeAttachmentList]
}

// newVolumeAttachments returns a VolumeAttachments
func newVolumeAttachments(c *StorageV1Client) *volumeAttachments {
	return &volumeAttachments{
		gentype.NewClientWithList[*storagev1.VolumeAttachment, *storagev1.VolumeAttachmentList](
			"volumeattachments",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *storagev1.VolumeAttachment { return &storagev1.VolumeAttachment{} },
			func() *storagev1.VolumeAttachmentList { return &storagev1.VolumeAttachmentList{} },
		),
	}
}
