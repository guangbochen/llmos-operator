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
	"net/http"

	v1 "github.com/llmos-ai/llmos-operator/pkg/apis/management.llmos.ai/v1"
	"github.com/llmos-ai/llmos-operator/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type ManagementV1Interface interface {
	RESTClient() rest.Interface
	GlobalRolesGetter
	ManagedAddonsGetter
	RoleTemplatesGetter
	RoleTemplateBindingsGetter
	SettingsGetter
	TokensGetter
	UpgradesGetter
	UsersGetter
	VersionsGetter
}

// ManagementV1Client is used to interact with features provided by the management.llmos.ai group.
type ManagementV1Client struct {
	restClient rest.Interface
}

func (c *ManagementV1Client) GlobalRoles() GlobalRoleInterface {
	return newGlobalRoles(c)
}

func (c *ManagementV1Client) ManagedAddons(namespace string) ManagedAddonInterface {
	return newManagedAddons(c, namespace)
}

func (c *ManagementV1Client) RoleTemplates() RoleTemplateInterface {
	return newRoleTemplates(c)
}

func (c *ManagementV1Client) RoleTemplateBindings() RoleTemplateBindingInterface {
	return newRoleTemplateBindings(c)
}

func (c *ManagementV1Client) Settings() SettingInterface {
	return newSettings(c)
}

func (c *ManagementV1Client) Tokens() TokenInterface {
	return newTokens(c)
}

func (c *ManagementV1Client) Upgrades() UpgradeInterface {
	return newUpgrades(c)
}

func (c *ManagementV1Client) Users() UserInterface {
	return newUsers(c)
}

func (c *ManagementV1Client) Versions() VersionInterface {
	return newVersions(c)
}

// NewForConfig creates a new ManagementV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ManagementV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ManagementV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ManagementV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ManagementV1Client{client}, nil
}

// NewForConfigOrDie creates a new ManagementV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ManagementV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ManagementV1Client for the given RESTClient.
func New(c rest.Interface) *ManagementV1Client {
	return &ManagementV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ManagementV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
