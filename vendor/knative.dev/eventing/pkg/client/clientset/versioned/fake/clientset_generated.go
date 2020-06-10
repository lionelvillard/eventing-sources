/*
Copyright 2020 The Knative Authors

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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
	clientset "knative.dev/eventing/pkg/client/clientset/versioned"
	configsv1alpha1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/configs/v1alpha1"
	fakeconfigsv1alpha1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/configs/v1alpha1/fake"
	eventingv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1beta1"
	fakeeventingv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1beta1/fake"
	flowsv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/flows/v1beta1"
	fakeflowsv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/flows/v1beta1/fake"
	messagingv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/messaging/v1beta1"
	fakemessagingv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/messaging/v1beta1/fake"
	sourcesv1alpha1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1alpha1"
	fakesourcesv1alpha1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1alpha1/fake"
	sourcesv1alpha2 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1alpha2"
	fakesourcesv1alpha2 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1alpha2/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// ConfigsV1alpha1 retrieves the ConfigsV1alpha1Client
func (c *Clientset) ConfigsV1alpha1() configsv1alpha1.ConfigsV1alpha1Interface {
	return &fakeconfigsv1alpha1.FakeConfigsV1alpha1{Fake: &c.Fake}
}

// EventingV1beta1 retrieves the EventingV1beta1Client
func (c *Clientset) EventingV1beta1() eventingv1beta1.EventingV1beta1Interface {
	return &fakeeventingv1beta1.FakeEventingV1beta1{Fake: &c.Fake}
}

// FlowsV1beta1 retrieves the FlowsV1beta1Client
func (c *Clientset) FlowsV1beta1() flowsv1beta1.FlowsV1beta1Interface {
	return &fakeflowsv1beta1.FakeFlowsV1beta1{Fake: &c.Fake}
}

// MessagingV1beta1 retrieves the MessagingV1beta1Client
func (c *Clientset) MessagingV1beta1() messagingv1beta1.MessagingV1beta1Interface {
	return &fakemessagingv1beta1.FakeMessagingV1beta1{Fake: &c.Fake}
}

// SourcesV1alpha1 retrieves the SourcesV1alpha1Client
func (c *Clientset) SourcesV1alpha1() sourcesv1alpha1.SourcesV1alpha1Interface {
	return &fakesourcesv1alpha1.FakeSourcesV1alpha1{Fake: &c.Fake}
}

// SourcesV1alpha2 retrieves the SourcesV1alpha2Client
func (c *Clientset) SourcesV1alpha2() sourcesv1alpha2.SourcesV1alpha2Interface {
	return &fakesourcesv1alpha2.FakeSourcesV1alpha2{Fake: &c.Fake}
}
