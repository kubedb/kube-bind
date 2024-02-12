/*
Copyright The Kube Bind Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1alpha1 "go.bytebuilders.dev/kube-bind/pkg/apis/kubebind/v1alpha1"
)

// APIServiceNamespaceLister helps list APIServiceNamespaces.
// All objects returned here must be treated as read-only.
type APIServiceNamespaceLister interface {
	// List lists all APIServiceNamespaces in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.APIServiceNamespace, err error)
	// APIServiceNamespaces returns an object that can list and get APIServiceNamespaces.
	APIServiceNamespaces(namespace string) APIServiceNamespaceNamespaceLister
	APIServiceNamespaceListerExpansion
}

// aPIServiceNamespaceLister implements the APIServiceNamespaceLister interface.
type aPIServiceNamespaceLister struct {
	indexer cache.Indexer
}

// NewAPIServiceNamespaceLister returns a new APIServiceNamespaceLister.
func NewAPIServiceNamespaceLister(indexer cache.Indexer) APIServiceNamespaceLister {
	return &aPIServiceNamespaceLister{indexer: indexer}
}

// List lists all APIServiceNamespaces in the indexer.
func (s *aPIServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.APIServiceNamespace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.APIServiceNamespace))
	})
	return ret, err
}

// APIServiceNamespaces returns an object that can list and get APIServiceNamespaces.
func (s *aPIServiceNamespaceLister) APIServiceNamespaces(namespace string) APIServiceNamespaceNamespaceLister {
	return aPIServiceNamespaceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// APIServiceNamespaceNamespaceLister helps list and get APIServiceNamespaces.
// All objects returned here must be treated as read-only.
type APIServiceNamespaceNamespaceLister interface {
	// List lists all APIServiceNamespaces in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.APIServiceNamespace, err error)
	// Get retrieves the APIServiceNamespace from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.APIServiceNamespace, error)
	APIServiceNamespaceNamespaceListerExpansion
}

// aPIServiceNamespaceNamespaceLister implements the APIServiceNamespaceNamespaceLister
// interface.
type aPIServiceNamespaceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all APIServiceNamespaces in the indexer for a given namespace.
func (s aPIServiceNamespaceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.APIServiceNamespace, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.APIServiceNamespace))
	})
	return ret, err
}

// Get retrieves the APIServiceNamespace from the indexer for a given namespace and name.
func (s aPIServiceNamespaceNamespaceLister) Get(name string) (*v1alpha1.APIServiceNamespace, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apiservicenamespace"), name)
	}
	return obj.(*v1alpha1.APIServiceNamespace), nil
}
