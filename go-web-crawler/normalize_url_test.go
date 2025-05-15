package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		errorContains string
	}{
		{
			name:     "standard https url",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "https url w/ trailling /",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:          "empty string",
			inputURL:      "",
			expected:      "",
			errorContains: "empty string",
		},
		{
			name:          "invalid protocol",
			inputURL:      "ftp://example.com",
			expected:      "",
			errorContains: "disallowed scheme",
		},
		{
			name:     "url with subdomain",
			inputURL: "https://api.example.com/v1/users",
			expected: "api.example.com/v1/users",
		},
		{
			name:     "url with port number and trailling /",
			inputURL: "https://localhost:8080/dashboard/",
			expected: "localhost:8080/dashboard",
		},
		{
			name:     "url with query parameters",
			inputURL: "https://search.boot.dev/results?q=golang&page=1",
			expected: "search.boot.dev/results",
		},
		{
			name:     "url with fragment",
			inputURL: "https://docs.boot.dev/tutorial#section-3",
			expected: "docs.boot.dev/tutorial",
		},
		{
			name:     "with no path",
			inputURL: "https://boot.dev",
			expected: "boot.dev",
		},
		{
			name:     "url with multiple trailing /",
			inputURL: "https://blog.boot.dev/path///",
			expected: "blog.boot.dev/path",
		},
		{
			name:          "malformed url with spaces",
			inputURL:      "https://example com/path",
			expected:      "",
			errorContains: "could not parse URL",
		},
		{
			name:     "lowercase capital letters",
			inputURL: "https://BLOG.boot.dev/PATH",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme and capitals and trailing slash",
			inputURL: "http://BLOG.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:          "handle invalid URL",
			inputURL:      `:\\invalidURL`,
			expected:      "",
			errorContains: "could not parse URL",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
