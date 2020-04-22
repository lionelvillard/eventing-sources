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

package source

import (
	"os"
	"testing"

	"knative.dev/pkg/configmap"
	. "knative.dev/pkg/reconciler/testing"

	// Fake injection clients and informers
	_ "knative.dev/eventing-contrib/github/pkg/client/injection/informers/sources/v1alpha1/githubsource/fake"
	_ "knative.dev/eventing/pkg/client/injection/client/fake"
	_ "knative.dev/pkg/client/injection/ducks/duck/v1/addressable/fake"
	_ "knative.dev/pkg/client/injection/kube/client/fake"
	_ "knative.dev/pkg/injection/clients/dynamicclient/fake"
	_ "knative.dev/serving/pkg/client/injection/client/fake"
	_ "knative.dev/serving/pkg/client/injection/informers/serving/v1/service/fake"
)

func TestNew(t *testing.T) {
	ctx, _ := SetupFakeContext(t)

	os.Setenv(raImageEnvVar, "dummy")
	os.Setenv("CONTROLLER_NAME", "dummy")
	os.Setenv("CONTROLLER_UID", "dummy")

	c := NewController(ctx, &configmap.InformedWatcher{})

	if c == nil {
		t.Fatal("Expected NewController to return a non-nil value")
	}
}
