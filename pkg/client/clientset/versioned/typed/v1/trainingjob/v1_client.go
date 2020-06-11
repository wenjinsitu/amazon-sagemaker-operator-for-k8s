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

package trainingjob

import (
	trainingjob "github.com/aws/amazon-sagemaker-operator-for-k8s/api/v1/trainingjob"
	"github.com/aws/amazon-sagemaker-operator-for-k8s/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type V1TrainingjobInterface interface {
	RESTClient() rest.Interface
	TrainingJobsGetter
}

// V1TrainingjobClient is used to interact with features provided by the v1 group.
type V1TrainingjobClient struct {
	restClient rest.Interface
}

func (c *V1TrainingjobClient) TrainingJobs(namespace string) TrainingJobInterface {
	return newTrainingJobs(c, namespace)
}

// NewForConfig creates a new V1TrainingjobClient for the given config.
func NewForConfig(c *rest.Config) (*V1TrainingjobClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &V1TrainingjobClient{client}, nil
}

// NewForConfigOrDie creates a new V1TrainingjobClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *V1TrainingjobClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new V1TrainingjobClient for the given RESTClient.
func New(c rest.Interface) *V1TrainingjobClient {
	return &V1TrainingjobClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := trainingjob.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *V1TrainingjobClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
