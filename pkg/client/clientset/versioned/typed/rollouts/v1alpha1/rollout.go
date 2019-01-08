/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/argoproj/rollout-controller/pkg/apis/rollouts/v1alpha1"
	scheme "github.com/argoproj/rollout-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RolloutsGetter has a method to return a RolloutInterface.
// A group's client should implement this interface.
type RolloutsGetter interface {
	Rollouts(namespace string) RolloutInterface
}

// RolloutInterface has methods to work with Rollout resources.
type RolloutInterface interface {
	Create(*v1alpha1.Rollout) (*v1alpha1.Rollout, error)
	Update(*v1alpha1.Rollout) (*v1alpha1.Rollout, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Rollout, error)
	List(opts v1.ListOptions) (*v1alpha1.RolloutList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	RolloutExpansion
}

// rollouts implements RolloutInterface
type rollouts struct {
	client rest.Interface
	ns     string
}

// newRollouts returns a Rollouts
func newRollouts(c *ArgoprojV1alpha1Client, namespace string) *rollouts {
	return &rollouts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the rollout, and returns the corresponding rollout object, and an error if there is any.
func (c *rollouts) Get(name string, options v1.GetOptions) (result *v1alpha1.Rollout, err error) {
	result = &v1alpha1.Rollout{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rollouts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Rollouts that match those selectors.
func (c *rollouts) List(opts v1.ListOptions) (result *v1alpha1.RolloutList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RolloutList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rollouts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rollouts.
func (c *rollouts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rollouts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a rollout and creates it.  Returns the server's representation of the rollout, and an error, if there is any.
func (c *rollouts) Create(rollout *v1alpha1.Rollout) (result *v1alpha1.Rollout, err error) {
	result = &v1alpha1.Rollout{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rollouts").
		Body(rollout).
		Do().
		Into(result)
	return
}

// Update takes the representation of a rollout and updates it. Returns the server's representation of the rollout, and an error, if there is any.
func (c *rollouts) Update(rollout *v1alpha1.Rollout) (result *v1alpha1.Rollout, err error) {
	result = &v1alpha1.Rollout{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rollouts").
		Name(rollout.Name).
		Body(rollout).
		Do().
		Into(result)
	return
}

// Delete takes name of the rollout and deletes it. Returns an error if one occurs.
func (c *rollouts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rollouts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rollouts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rollouts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}
