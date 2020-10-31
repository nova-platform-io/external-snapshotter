/*
Copyright 2020 The Kubernetes Authors.

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

package v1beta1

import (
	"context"
	time "time"

	volumesnapshotv1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v3/apis/volumesnapshot/v1beta1"
	versioned "github.com/kubernetes-csi/external-snapshotter/client/v3/clientset/versioned"
	internalinterfaces "github.com/kubernetes-csi/external-snapshotter/client/v3/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v3/listers/volumesnapshot/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VolumeSnapshotContentInformer provides access to a shared informer and lister for
// VolumeSnapshotContents.
type VolumeSnapshotContentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.VolumeSnapshotContentLister
}

type volumeSnapshotContentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewVolumeSnapshotContentInformer constructs a new informer for VolumeSnapshotContent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVolumeSnapshotContentInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVolumeSnapshotContentInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredVolumeSnapshotContentInformer constructs a new informer for VolumeSnapshotContent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVolumeSnapshotContentInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SnapshotV1beta1().VolumeSnapshotContents().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SnapshotV1beta1().VolumeSnapshotContents().Watch(context.TODO(), options)
			},
		},
		&volumesnapshotv1beta1.VolumeSnapshotContent{},
		resyncPeriod,
		indexers,
	)
}

func (f *volumeSnapshotContentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVolumeSnapshotContentInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *volumeSnapshotContentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&volumesnapshotv1beta1.VolumeSnapshotContent{}, f.defaultInformer)
}

func (f *volumeSnapshotContentInformer) Lister() v1beta1.VolumeSnapshotContentLister {
	return v1beta1.NewVolumeSnapshotContentLister(f.Informer().GetIndexer())
}