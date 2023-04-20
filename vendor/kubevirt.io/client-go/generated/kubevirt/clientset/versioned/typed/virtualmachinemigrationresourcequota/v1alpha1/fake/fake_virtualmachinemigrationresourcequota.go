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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubevirt.io/api/virtualMachineMigrationResourceQuota/v1alpha1"
)

// FakeVirtualMachineMigrationResourceQuotas implements VirtualMachineMigrationResourceQuotaInterface
type FakeVirtualMachineMigrationResourceQuotas struct {
	Fake *FakeVirtualMachineMigrationResourceQuotaV1alpha1
	ns   string
}

var virtualmachinemigrationresourcequotasResource = schema.GroupVersionResource{Group: "virtualMachineMigrationResourceQuota.kubevirt.io", Version: "v1alpha1", Resource: "virtualmachinemigrationresourcequotas"}

var virtualmachinemigrationresourcequotasKind = schema.GroupVersionKind{Group: "virtualMachineMigrationResourceQuota.kubevirt.io", Version: "v1alpha1", Kind: "VirtualMachineMigrationResourceQuota"}

// Get takes name of the virtualMachineMigrationResourceQuota, and returns the corresponding virtualMachineMigrationResourceQuota object, and an error if there is any.
func (c *FakeVirtualMachineMigrationResourceQuotas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineMigrationResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachinemigrationresourcequotasResource, c.ns, name), &v1alpha1.VirtualMachineMigrationResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineMigrationResourceQuota), err
}

// List takes label and field selectors, and returns the list of VirtualMachineMigrationResourceQuotas that match those selectors.
func (c *FakeVirtualMachineMigrationResourceQuotas) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualMachineMigrationResourceQuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachinemigrationresourcequotasResource, virtualmachinemigrationresourcequotasKind, c.ns, opts), &v1alpha1.VirtualMachineMigrationResourceQuotaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualMachineMigrationResourceQuotaList{ListMeta: obj.(*v1alpha1.VirtualMachineMigrationResourceQuotaList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualMachineMigrationResourceQuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineMigrationResourceQuotas.
func (c *FakeVirtualMachineMigrationResourceQuotas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachinemigrationresourcequotasResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineMigrationResourceQuota and creates it.  Returns the server's representation of the virtualMachineMigrationResourceQuota, and an error, if there is any.
func (c *FakeVirtualMachineMigrationResourceQuotas) Create(ctx context.Context, virtualMachineMigrationResourceQuota *v1alpha1.VirtualMachineMigrationResourceQuota, opts v1.CreateOptions) (result *v1alpha1.VirtualMachineMigrationResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachinemigrationresourcequotasResource, c.ns, virtualMachineMigrationResourceQuota), &v1alpha1.VirtualMachineMigrationResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineMigrationResourceQuota), err
}

// Update takes the representation of a virtualMachineMigrationResourceQuota and updates it. Returns the server's representation of the virtualMachineMigrationResourceQuota, and an error, if there is any.
func (c *FakeVirtualMachineMigrationResourceQuotas) Update(ctx context.Context, virtualMachineMigrationResourceQuota *v1alpha1.VirtualMachineMigrationResourceQuota, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineMigrationResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachinemigrationresourcequotasResource, c.ns, virtualMachineMigrationResourceQuota), &v1alpha1.VirtualMachineMigrationResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineMigrationResourceQuota), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMachineMigrationResourceQuotas) UpdateStatus(ctx context.Context, virtualMachineMigrationResourceQuota *v1alpha1.VirtualMachineMigrationResourceQuota, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineMigrationResourceQuota, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachinemigrationresourcequotasResource, "status", c.ns, virtualMachineMigrationResourceQuota), &v1alpha1.VirtualMachineMigrationResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineMigrationResourceQuota), err
}

// Delete takes name of the virtualMachineMigrationResourceQuota and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineMigrationResourceQuotas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachinemigrationresourcequotasResource, c.ns, name), &v1alpha1.VirtualMachineMigrationResourceQuota{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineMigrationResourceQuotas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachinemigrationresourcequotasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualMachineMigrationResourceQuotaList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineMigrationResourceQuota.
func (c *FakeVirtualMachineMigrationResourceQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineMigrationResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachinemigrationresourcequotasResource, c.ns, name, pt, data, subresources...), &v1alpha1.VirtualMachineMigrationResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineMigrationResourceQuota), err
}
