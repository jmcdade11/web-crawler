package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	parsedHtml, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	var walkNodes func(*html.Node)
	urls := []string{}
	walkNodes = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					val, err := url.Parse(a.Val)
					if err != nil {
						continue
					}
					url := baseURL.ResolveReference(val)
					urls = append(urls, url.String())
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			walkNodes(c)
		}
	}
	walkNodes(parsedHtml)
	if len(urls) == 0 {
		return nil, fmt.Errorf("no URL found: %s", rawBaseURL)
	}
	return urls, nil
}
