/*

Don't alter this file, it was generated.

*/
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "code.cloudfoundry.org/cf-operator/pkg/kube/apis/quarkssecret/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// QuarksSecretLister helps list QuarksSecrets.
type QuarksSecretLister interface {
	// List lists all QuarksSecrets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.QuarksSecret, err error)
	// QuarksSecrets returns an object that can list and get QuarksSecrets.
	QuarksSecrets(namespace string) QuarksSecretNamespaceLister
	QuarksSecretListerExpansion
}

// quarksSecretLister implements the QuarksSecretLister interface.
type quarksSecretLister struct {
	indexer cache.Indexer
}

// NewQuarksSecretLister returns a new QuarksSecretLister.
func NewQuarksSecretLister(indexer cache.Indexer) QuarksSecretLister {
	return &quarksSecretLister{indexer: indexer}
}

// List lists all QuarksSecrets in the indexer.
func (s *quarksSecretLister) List(selector labels.Selector) (ret []*v1alpha1.QuarksSecret, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.QuarksSecret))
	})
	return ret, err
}

// QuarksSecrets returns an object that can list and get QuarksSecrets.
func (s *quarksSecretLister) QuarksSecrets(namespace string) QuarksSecretNamespaceLister {
	return quarksSecretNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// QuarksSecretNamespaceLister helps list and get QuarksSecrets.
type QuarksSecretNamespaceLister interface {
	// List lists all QuarksSecrets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.QuarksSecret, err error)
	// Get retrieves the QuarksSecret from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.QuarksSecret, error)
	QuarksSecretNamespaceListerExpansion
}

// quarksSecretNamespaceLister implements the QuarksSecretNamespaceLister
// interface.
type quarksSecretNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all QuarksSecrets in the indexer for a given namespace.
func (s quarksSecretNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.QuarksSecret, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.QuarksSecret))
	})
	return ret, err
}

// Get retrieves the QuarksSecret from the indexer for a given namespace and name.
func (s quarksSecretNamespaceLister) Get(name string) (*v1alpha1.QuarksSecret, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("quarkssecret"), name)
	}
	return obj.(*v1alpha1.QuarksSecret), nil
}
