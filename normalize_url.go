package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	inputURL = strings.TrimSpace(inputURL)
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("could not parse URL: %w", err)
	}

	normalized := parsedURL.Host + parsedURL.Path
	normalized = strings.TrimSuffix(normalized, "/")
	return normalized, nil
}
