package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error while executing GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("error response: %s", res.Status)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("response is not html: %s", contentType)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("could not process response body: %v", err)
	}

	body := string(bodyBytes)
	return body, nil
}
