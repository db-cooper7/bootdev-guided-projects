package main

import (
	"errors"
	"net/url"
	"strings"
)

func normalizeURL(urlString string) (string, error) {
	if urlString == "" {
		return "", errors.New("empty string")
	}

	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", errors.New("could not parse URL: " + err.Error())
	}

	scheme := parsedURL.Scheme
	if scheme != "http" && scheme != "https" {
		return "", errors.New("disallowed scheme: " + scheme)
	}

	path := parsedURL.Path
	for strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}

	if path == "" {
		return strings.ToLower(parsedURL.Host), nil
	}

	return strings.ToLower(parsedURL.Host + path), nil
}
