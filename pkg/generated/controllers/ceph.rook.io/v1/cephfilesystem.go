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
	"github.com/rancher/wrangler/v3/pkg/generic"
	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
)

// CephFilesystemController interface for managing CephFilesystem resources.
type CephFilesystemController interface {
	generic.ControllerInterface[*v1.CephFilesystem, *v1.CephFilesystemList]
}

// CephFilesystemClient interface for managing CephFilesystem resources in Kubernetes.
type CephFilesystemClient interface {
	generic.ClientInterface[*v1.CephFilesystem, *v1.CephFilesystemList]
}

// CephFilesystemCache interface for retrieving CephFilesystem resources in memory.
type CephFilesystemCache interface {
	generic.CacheInterface[*v1.CephFilesystem]
}
