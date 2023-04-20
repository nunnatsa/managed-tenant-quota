/*
Copyright 2023 The KubeVirt Authors.

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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubevirt.io/api/virtualMachineMigrationResourceQuota/v1alpha1"
	scheme "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/scheme"
)

// VirtualMachineMigrationResourceQuotasGetter has a method to return a VirtualMachineMigrationResourceQuotaInterface.
// A group's client should implement this interface.
type VirtualMachineMigrationResourceQuotasGetter interface {
	VirtualMachineMigrationResourceQuotas(namespace string) VirtualMachineMigrationResourceQuotaInterface
}

// VirtualMachineMigrationResourceQuotaInterface has methods to work with VirtualMachineMigrationResourceQuota resources.
type VirtualMachineMigrationResourceQuotaInterface interface {
	Create(ctx context.Context, virtualMachineMigrationResourceQuota *v1alpha1.VirtualMachineMigrationResourceQuota, opts v1.CreateOptions) (*v1alpha1.VirtualMachineMigrationResourceQuota, error)
	Update(ctx context.Context, virtualMachineMigrationResourceQuota *v1alpha1.VirtualMachineMigrationResourceQuota, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineMigrationResourceQuota, error)
	UpdateStatus(ctx context.Context, virtualMachineMigrationResourceQuota *v1alpha1.VirtualMachineMigrationResourceQuota, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineMigrationResourceQuota, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VirtualMachineMigrationResourceQuota, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VirtualMachineMigrationResourceQuotaList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineMigrationResourceQuota, err error)
	VirtualMachineMigrationResourceQuotaExpansion
}

// virtualMachineMigrationResourceQuotas implements VirtualMachineMigrationResourceQuotaInterface
type virtualMachineMigrationResourceQuotas struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineMigrationResourceQuotas returns a VirtualMachineMigrationResourceQuotas
func newVirtualMachineMigrationResourceQuotas(c *VirtualMachineMigrationResourceQuotaV1alpha1Client, namespace string) *virtualMachineMigrationResourceQuotas {
	return &virtualMachineMigrationResourceQuotas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineMigrationResourceQuota, and returns the corresponding virtualMachineMigrationResourceQuota object, and an error if there is any.
func (c *virtualMachineMigrationResourceQuotas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineMigrationResourceQuota, err error) {
	result = &v1alpha1.VirtualMachineMigrationResourceQuota{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinemigrationresourcequotas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineMigrationResourceQuotas that match those selectors.
func (c *virtualMachineMigrationResourceQuotas) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualMachineMigrationResourceQuotaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualMachineMigrationResourceQuotaList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinemigrationresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineMigrationResourceQuotas.
func (c *virtualMachineMigrationResourceQuotas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinemigrationresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualMachineMigrationResourceQuota and creates it.  Returns the server's representation of the virtualMachineMigrationResourceQuota, and an error, if there is any.
func (c *virtualMachineMigrationResourceQuotas) Create(ctx context.Context, virtualMachineMigrationResourceQuota *v1alpha1.VirtualMachineMigrationResourceQuota, opts v1.CreateOptions) (result *v1alpha1.VirtualMachineMigrationResourceQuota, err error) {
	result = &v1alpha1.VirtualMachineMigrationResourceQuota{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinemigrationresourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineMigrationResourceQuota).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualMachineMigrationResourceQuota and updates it. Returns the server's representation of the virtualMachineMigrationResourceQuota, and an error, if there is any.
func (c *virtualMachineMigrationResourceQuotas) Update(ctx context.Context, virtualMachineMigrationResourceQuota *v1alpha1.VirtualMachineMigrationResourceQuota, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineMigrationResourceQuota, err error) {
	result = &v1alpha1.VirtualMachineMigrationResourceQuota{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinemigrationresourcequotas").
		Name(virtualMachineMigrationResourceQuota.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineMigrationResourceQuota).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *virtualMachineMigrationResourceQuotas) UpdateStatus(ctx context.Context, virtualMachineMigrationResourceQuota *v1alpha1.VirtualMachineMigrationResourceQuota, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineMigrationResourceQuota, err error) {
	result = &v1alpha1.VirtualMachineMigrationResourceQuota{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinemigrationresourcequotas").
		Name(virtualMachineMigrationResourceQuota.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineMigrationResourceQuota).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualMachineMigrationResourceQuota and deletes it. Returns an error if one occurs.
func (c *virtualMachineMigrationResourceQuotas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinemigrationresourcequotas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineMigrationResourceQuotas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinemigrationresourcequotas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualMachineMigrationResourceQuota.
func (c *virtualMachineMigrationResourceQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineMigrationResourceQuota, err error) {
	result = &v1alpha1.VirtualMachineMigrationResourceQuota{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinemigrationresourcequotas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
