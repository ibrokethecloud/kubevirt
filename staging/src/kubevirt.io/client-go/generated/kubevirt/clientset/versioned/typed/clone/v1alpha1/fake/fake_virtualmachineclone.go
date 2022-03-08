/*
Copyright 2022 The KubeVirt Authors.

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

	v1alpha1 "kubevirt.io/api/clone/v1alpha1"
)

// FakeVirtualMachineClones implements VirtualMachineCloneInterface
type FakeVirtualMachineClones struct {
	Fake *FakeCloneV1alpha1
	ns   string
}

var virtualmachineclonesResource = schema.GroupVersionResource{Group: "clone.kubevirt.io", Version: "v1alpha1", Resource: "virtualmachineclones"}

var virtualmachineclonesKind = schema.GroupVersionKind{Group: "clone.kubevirt.io", Version: "v1alpha1", Kind: "VirtualMachineClone"}

// Get takes name of the virtualMachineClone, and returns the corresponding virtualMachineClone object, and an error if there is any.
func (c *FakeVirtualMachineClones) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineClone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineclonesResource, c.ns, name), &v1alpha1.VirtualMachineClone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineClone), err
}

// List takes label and field selectors, and returns the list of VirtualMachineClones that match those selectors.
func (c *FakeVirtualMachineClones) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualMachineCloneList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineclonesResource, virtualmachineclonesKind, c.ns, opts), &v1alpha1.VirtualMachineCloneList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualMachineCloneList{ListMeta: obj.(*v1alpha1.VirtualMachineCloneList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualMachineCloneList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineClones.
func (c *FakeVirtualMachineClones) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineclonesResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineClone and creates it.  Returns the server's representation of the virtualMachineClone, and an error, if there is any.
func (c *FakeVirtualMachineClones) Create(ctx context.Context, virtualMachineClone *v1alpha1.VirtualMachineClone, opts v1.CreateOptions) (result *v1alpha1.VirtualMachineClone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineclonesResource, c.ns, virtualMachineClone), &v1alpha1.VirtualMachineClone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineClone), err
}

// Update takes the representation of a virtualMachineClone and updates it. Returns the server's representation of the virtualMachineClone, and an error, if there is any.
func (c *FakeVirtualMachineClones) Update(ctx context.Context, virtualMachineClone *v1alpha1.VirtualMachineClone, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineClone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineclonesResource, c.ns, virtualMachineClone), &v1alpha1.VirtualMachineClone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineClone), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMachineClones) UpdateStatus(ctx context.Context, virtualMachineClone *v1alpha1.VirtualMachineClone, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineClone, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachineclonesResource, "status", c.ns, virtualMachineClone), &v1alpha1.VirtualMachineClone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineClone), err
}

// Delete takes name of the virtualMachineClone and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineClones) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachineclonesResource, c.ns, name), &v1alpha1.VirtualMachineClone{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineClones) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachineclonesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualMachineCloneList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineClone.
func (c *FakeVirtualMachineClones) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineClone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineclonesResource, c.ns, name, pt, data, subresources...), &v1alpha1.VirtualMachineClone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineClone), err
}
