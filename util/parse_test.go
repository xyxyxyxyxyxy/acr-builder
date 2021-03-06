// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package util

import (
	"testing"
)

// TestParseTags tests parsing tags off a build command.
func TestParseTags(t *testing.T) {
	tests := []struct {
		id       int
		build    string
		expected []string
	}{
		// Tag tests
		{1, "-f Dockerfile -t {{.Run.ID}}:latest --tag blah https://github.com/Azure/acr-builder.git", []string{"{{.Run.ID}}:latest", "blah"}},
		{2, "--tag foo https://github.com/Azure/acr-builder.git --tag bar -t qux", []string{"foo", "bar", "qux"}},
	}

	for _, test := range tests {
		actual := ParseTags(test.build)
		if !StringSequenceEquals(actual, test.expected) {
			t.Errorf("Test %d failed. Expected %v, got %v", test.id, test.expected, actual)
		}
	}
}

// TestParseBuildArgs tests parsing build args off a build command.
func TestParseBuildArgs(t *testing.T) {
	tests := []struct {
		id       int
		build    string
		expected []string
	}{
		{1, "-f Dockerfile -t hello:world --build-arg foo https://github.com/Azure/acr-builder.git --build-arg bar", []string{"foo", "bar"}},
		{2, "-f Dockerfile -t hello:world --buildarg ignored --build-arg foo=bar https://github.com/Azure/acr-builder.git --build-arg hello=world", []string{"foo=bar", "hello=world"}},
	}

	for _, test := range tests {
		actual := ParseBuildArgs(test.build)
		if !StringSequenceEquals(actual, test.expected) {
			t.Errorf("Test %d failed. Expected %v, got %v", test.id, test.expected, actual)
		}
	}
}
