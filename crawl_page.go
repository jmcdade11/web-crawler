package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}
	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizeURL: %v\n", err)
		return
	}

	if !cfg.addPageVisit(normalizedCurrent) {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v\n", err)
		return
	}

	nextURLs, err := getURLsFromHTML(html, rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v\n", err)
		return
	}

	for _, nextURL := range nextURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
