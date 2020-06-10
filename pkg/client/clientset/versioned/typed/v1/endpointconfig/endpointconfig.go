/*
Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.

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
// Code generated by client-gen. DO NOT EDIT.

package endpointconfig

import (
	"time"

	endpointconfig "github.com/aws/amazon-sagemaker-operator-for-k8s/api/v1/endpointconfig"
	scheme "github.com/aws/amazon-sagemaker-operator-for-k8s/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EndpointConfigsGetter has a method to return a EndpointConfigInterface.
// A group's client should implement this interface.
type EndpointConfigsGetter interface {
	EndpointConfigs(namespace string) EndpointConfigInterface
}

// EndpointConfigInterface has methods to work with EndpointConfig resources.
type EndpointConfigInterface interface {
	Create(*endpointconfig.EndpointConfig) (*endpointconfig.EndpointConfig, error)
	Update(*endpointconfig.EndpointConfig) (*endpointconfig.EndpointConfig, error)
	UpdateStatus(*endpointconfig.EndpointConfig) (*endpointconfig.EndpointConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*endpointconfig.EndpointConfig, error)
	List(opts v1.ListOptions) (*endpointconfig.EndpointConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *endpointconfig.EndpointConfig, err error)
	EndpointConfigExpansion
}

// endpointConfigs implements EndpointConfigInterface
type endpointConfigs struct {
	client rest.Interface
	ns     string
}

// newEndpointConfigs returns a EndpointConfigs
func newEndpointConfigs(c *V1EndpointconfigClient, namespace string) *endpointConfigs {
	return &endpointConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the endpointConfig, and returns the corresponding endpointConfig object, and an error if there is any.
func (c *endpointConfigs) Get(name string, options v1.GetOptions) (result *endpointconfig.EndpointConfig, err error) {
	result = &endpointconfig.EndpointConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpointconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EndpointConfigs that match those selectors.
func (c *endpointConfigs) List(opts v1.ListOptions) (result *endpointconfig.EndpointConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &endpointconfig.EndpointConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpointconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested endpointConfigs.
func (c *endpointConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("endpointconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a endpointConfig and creates it.  Returns the server's representation of the endpointConfig, and an error, if there is any.
func (c *endpointConfigs) Create(endpointConfig *endpointconfig.EndpointConfig) (result *endpointconfig.EndpointConfig, err error) {
	result = &endpointconfig.EndpointConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("endpointconfigs").
		Body(endpointConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a endpointConfig and updates it. Returns the server's representation of the endpointConfig, and an error, if there is any.
func (c *endpointConfigs) Update(endpointConfig *endpointconfig.EndpointConfig) (result *endpointconfig.EndpointConfig, err error) {
	result = &endpointconfig.EndpointConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("endpointconfigs").
		Name(endpointConfig.Name).
		Body(endpointConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *endpointConfigs) UpdateStatus(endpointConfig *endpointconfig.EndpointConfig) (result *endpointconfig.EndpointConfig, err error) {
	result = &endpointconfig.EndpointConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("endpointconfigs").
		Name(endpointConfig.Name).
		SubResource("status").
		Body(endpointConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the endpointConfig and deletes it. Returns an error if one occurs.
func (c *endpointConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpointconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *endpointConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpointconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched endpointConfig.
func (c *endpointConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *endpointconfig.EndpointConfig, err error) {
	result = &endpointconfig.EndpointConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("endpointconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
