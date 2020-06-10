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
// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	time "time"

	batchtransformjob "github.com/aws/amazon-sagemaker-operator-for-k8s/api/v1/batchtransformjob"
	internalinterfaces "github.com/aws/amazon-sagemaker-operator-for-k8s/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "github.com/aws/amazon-sagemaker-operator-for-k8s/pkg/client/listers/batchtransformjob/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	internalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
)

// BatchTransformJobInformer provides access to a shared informer and lister for
// BatchTransformJobs.
type BatchTransformJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.BatchTransformJobLister
}

type batchTransformJobInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBatchTransformJobInformer constructs a new informer for BatchTransformJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBatchTransformJobInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBatchTransformJobInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBatchTransformJobInformer constructs a new informer for BatchTransformJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBatchTransformJobInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Batchtransformjob().BatchTransformJobs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Batchtransformjob().BatchTransformJobs(namespace).Watch(options)
			},
		},
		&batchtransformjob.BatchTransformJob{},
		resyncPeriod,
		indexers,
	)
}

func (f *batchTransformJobInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBatchTransformJobInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *batchTransformJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&batchtransformjob.BatchTransformJob{}, f.defaultInformer)
}

func (f *batchTransformJobInformer) Lister() internalversion.BatchTransformJobLister {
	return internalversion.NewBatchTransformJobLister(f.Informer().GetIndexer())
}
