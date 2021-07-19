/*

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
	v1alpha2 "kubesphere.io/fluentbit-operator/api/fluentbitoperator/v1alpha2"
)

// FakeFilters implements FilterInterface
type FakeFilters struct {
	Fake *FakeLoggingV1alpha2
	ns   string
}

var filtersResource = schema.GroupVersionResource{Group: "logging.kubesphere.io", Version: "v1alpha2", Resource: "filters"}

var filtersKind = schema.GroupVersionKind{Group: "logging.kubesphere.io", Version: "v1alpha2", Kind: "Filter"}

// Get takes name of the filter, and returns the corresponding filter object, and an error if there is any.
func (c *FakeFilters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Filter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(filtersResource, c.ns, name), &v1alpha2.Filter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Filter), err
}

// List takes label and field selectors, and returns the list of Filters that match those selectors.
func (c *FakeFilters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.FilterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(filtersResource, filtersKind, c.ns, opts), &v1alpha2.FilterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.FilterList{ListMeta: obj.(*v1alpha2.FilterList).ListMeta}
	for _, item := range obj.(*v1alpha2.FilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested filters.
func (c *FakeFilters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(filtersResource, c.ns, opts))

}

// Create takes the representation of a filter and creates it.  Returns the server's representation of the filter, and an error, if there is any.
func (c *FakeFilters) Create(ctx context.Context, filter *v1alpha2.Filter, opts v1.CreateOptions) (result *v1alpha2.Filter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(filtersResource, c.ns, filter), &v1alpha2.Filter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Filter), err
}

// Update takes the representation of a filter and updates it. Returns the server's representation of the filter, and an error, if there is any.
func (c *FakeFilters) Update(ctx context.Context, filter *v1alpha2.Filter, opts v1.UpdateOptions) (result *v1alpha2.Filter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(filtersResource, c.ns, filter), &v1alpha2.Filter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Filter), err
}

// Delete takes name of the filter and deletes it. Returns an error if one occurs.
func (c *FakeFilters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(filtersResource, c.ns, name), &v1alpha2.Filter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFilters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(filtersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.FilterList{})
	return err
}

// Patch applies the patch and returns the patched filter.
func (c *FakeFilters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Filter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(filtersResource, c.ns, name, pt, data, subresources...), &v1alpha2.Filter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Filter), err
}
