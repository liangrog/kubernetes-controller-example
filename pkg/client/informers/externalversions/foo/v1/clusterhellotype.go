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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	foov1 "github.com/liangrog/kube-crd-example/pkg/apis/foo/v1"
	versioned "github.com/liangrog/kube-crd-example/pkg/client/clientset/versioned"
	internalinterfaces "github.com/liangrog/kube-crd-example/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/liangrog/kube-crd-example/pkg/client/listers/foo/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterHelloTypeInformer provides access to a shared informer and lister for
// ClusterHelloTypes.
type ClusterHelloTypeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterHelloTypeLister
}

type clusterHelloTypeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterHelloTypeInformer constructs a new informer for ClusterHelloType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterHelloTypeInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterHelloTypeInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterHelloTypeInformer constructs a new informer for ClusterHelloType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterHelloTypeInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1().ClusterHelloTypes().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1().ClusterHelloTypes().Watch(options)
			},
		},
		&foov1.ClusterHelloType{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterHelloTypeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterHelloTypeInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterHelloTypeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&foov1.ClusterHelloType{}, f.defaultInformer)
}

func (f *clusterHelloTypeInformer) Lister() v1.ClusterHelloTypeLister {
	return v1.NewClusterHelloTypeLister(f.Informer().GetIndexer())
}
