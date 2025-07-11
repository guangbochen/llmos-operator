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
	v1 "github.com/llmos-ai/llmos-operator/pkg/generated/clientset/versioned/typed/ml.llmos.ai/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMlV1 struct {
	*testing.Fake
}

func (c *FakeMlV1) Datasets(namespace string) v1.DatasetInterface {
	return newFakeDatasets(c, namespace)
}

func (c *FakeMlV1) DatasetVersions(namespace string) v1.DatasetVersionInterface {
	return newFakeDatasetVersions(c, namespace)
}

func (c *FakeMlV1) LocalModels(namespace string) v1.LocalModelInterface {
	return newFakeLocalModels(c, namespace)
}

func (c *FakeMlV1) LocalModelVersions(namespace string) v1.LocalModelVersionInterface {
	return newFakeLocalModelVersions(c, namespace)
}

func (c *FakeMlV1) Models(namespace string) v1.ModelInterface {
	return newFakeModels(c, namespace)
}

func (c *FakeMlV1) ModelServices(namespace string) v1.ModelServiceInterface {
	return newFakeModelServices(c, namespace)
}

func (c *FakeMlV1) Notebooks(namespace string) v1.NotebookInterface {
	return newFakeNotebooks(c, namespace)
}

func (c *FakeMlV1) Registries() v1.RegistryInterface {
	return newFakeRegistries(c)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMlV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
