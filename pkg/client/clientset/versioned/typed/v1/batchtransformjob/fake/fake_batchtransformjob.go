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

package fake

import (
	batchtransformjob "github.com/aws/amazon-sagemaker-operator-for-k8s/api/v1/batchtransformjob"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBatchTransformJobs implements BatchTransformJobInterface
type FakeBatchTransformJobs struct {
	Fake *FakeV1Batchtransformjob
	ns   string
}

var batchtransformjobsResource = schema.GroupVersionResource{Group: "v1", Version: "batchtransformjob", Resource: "batchtransformjobs"}

var batchtransformjobsKind = schema.GroupVersionKind{Group: "v1", Version: "batchtransformjob", Kind: "BatchTransformJob"}

// Get takes name of the batchTransformJob, and returns the corresponding batchTransformJob object, and an error if there is any.
func (c *FakeBatchTransformJobs) Get(name string, options v1.GetOptions) (result *batchtransformjob.BatchTransformJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(batchtransformjobsResource, c.ns, name), &batchtransformjob.BatchTransformJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchtransformjob.BatchTransformJob), err
}

// List takes label and field selectors, and returns the list of BatchTransformJobs that match those selectors.
func (c *FakeBatchTransformJobs) List(opts v1.ListOptions) (result *batchtransformjob.BatchTransformJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(batchtransformjobsResource, batchtransformjobsKind, c.ns, opts), &batchtransformjob.BatchTransformJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &batchtransformjob.BatchTransformJobList{ListMeta: obj.(*batchtransformjob.BatchTransformJobList).ListMeta}
	for _, item := range obj.(*batchtransformjob.BatchTransformJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested batchTransformJobs.
func (c *FakeBatchTransformJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(batchtransformjobsResource, c.ns, opts))

}

// Create takes the representation of a batchTransformJob and creates it.  Returns the server's representation of the batchTransformJob, and an error, if there is any.
func (c *FakeBatchTransformJobs) Create(batchTransformJob *batchtransformjob.BatchTransformJob) (result *batchtransformjob.BatchTransformJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(batchtransformjobsResource, c.ns, batchTransformJob), &batchtransformjob.BatchTransformJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchtransformjob.BatchTransformJob), err
}

// Update takes the representation of a batchTransformJob and updates it. Returns the server's representation of the batchTransformJob, and an error, if there is any.
func (c *FakeBatchTransformJobs) Update(batchTransformJob *batchtransformjob.BatchTransformJob) (result *batchtransformjob.BatchTransformJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(batchtransformjobsResource, c.ns, batchTransformJob), &batchtransformjob.BatchTransformJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchtransformjob.BatchTransformJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBatchTransformJobs) UpdateStatus(batchTransformJob *batchtransformjob.BatchTransformJob) (*batchtransformjob.BatchTransformJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(batchtransformjobsResource, "status", c.ns, batchTransformJob), &batchtransformjob.BatchTransformJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchtransformjob.BatchTransformJob), err
}

// Delete takes name of the batchTransformJob and deletes it. Returns an error if one occurs.
func (c *FakeBatchTransformJobs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(batchtransformjobsResource, c.ns, name), &batchtransformjob.BatchTransformJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBatchTransformJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(batchtransformjobsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &batchtransformjob.BatchTransformJobList{})
	return err
}

// Patch applies the patch and returns the patched batchTransformJob.
func (c *FakeBatchTransformJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batchtransformjob.BatchTransformJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(batchtransformjobsResource, c.ns, name, pt, data, subresources...), &batchtransformjob.BatchTransformJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batchtransformjob.BatchTransformJob), err
}
