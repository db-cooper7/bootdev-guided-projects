package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(urlString string) (string, error) {
	if urlString == "" {
		return "", fmt.Errorf("empty string")
	}

	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", fmt.Errorf("could not parse URL: %w", err)
	}

	scheme := parsedURL.Scheme
	if scheme != "http" && scheme != "https" {
		return "", fmt.Errorf("disallowed scheme: %s", scheme)
	}

	path := parsedURL.Path
	for strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}

	return strings.ToLower(parsedURL.Host + path), nil
}
