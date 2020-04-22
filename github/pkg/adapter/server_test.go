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

package adapter

import (
	"net/http/httptest"
	"testing"

	logtesting "knative.dev/pkg/logging/testing"
)

func TestGitHubServer(t *testing.T) {
	logger := logtesting.TestLogger(t)
	handler := NewHandler(logger)

	s := httptest.NewServer(handler)
	defer s.Close()

	// Not Found
	resp, err := s.Client().Get(s.URL + "/does/not/exist")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp.StatusCode != 404 {
		t.Fatalf("Unexpected status code. Wanted 404, got %d", resp.StatusCode)
	}

	// Valid
	handler.Register("/valid/path", &fakeHandler{
		handler: sinkAccepted,
	})

	resp, err = s.Client().Get(s.URL + "/valid/path")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Unexpected status code. Wanted 200, got %d", resp.StatusCode)
	}

}
