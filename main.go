package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Page struct {
	URL   string
	Count int
}

func main() {
	// usage: ./web-crawler URL maxConcurrency maxPages
	if len(os.Args) < 4 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("could not parse maxConcurrency")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Could not parse maxPages")
		os.Exit(1)
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("configured maxConcurrency: %d, maxPages: %d\n", maxConcurrency, maxPages)

	fmt.Printf("starting crawl: %s\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

	sorted := sortPages(pages)
	for _, page := range sorted {
		url := page.URL
		count := page.Count
		fmt.Printf("Found %d internal links to %s\n", count, url)
	}
}

func sortPages(pages map[string]int) []Page {
	pagesSlice := []Page{}
	for url, count := range pages {
		pagesSlice = append(pagesSlice, Page{URL: url, Count: count})
	}
	sort.Slice(pagesSlice, func(i, j int) bool {
		if pagesSlice[i].Count == pagesSlice[j].Count {
			return pagesSlice[i].URL < pagesSlice[j].URL
		}
		return pagesSlice[i].Count > pagesSlice[j].Count
	})
	return pagesSlice
}
