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

package v1

import (
	v1 "github.com/llmos-ai/llmos-controller/pkg/apis/management.llmos.ai/v1"
	"github.com/rancher/wrangler/v2/pkg/generic"
)

// UserController interface for managing User resources.
type UserController interface {
	generic.NonNamespacedControllerInterface[*v1.User, *v1.UserList]
}

// UserClient interface for managing User resources in Kubernetes.
type UserClient interface {
	generic.NonNamespacedClientInterface[*v1.User, *v1.UserList]
}

// UserCache interface for retrieving User resources in memory.
type UserCache interface {
	generic.NonNamespacedCacheInterface[*v1.User]
}