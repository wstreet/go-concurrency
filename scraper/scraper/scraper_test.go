package scraper

import (
	"sync"
	"testing"
)

func TestScrape(t *testing.T) {
	url := "https://www.google.com"
	results := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go Scrape(url, &wg, results)

	go func() {
		wg.Wait()
		close(results)
	}()
	result := <-results
	if result == "" {
		t.Errorf("Expected a non-empty result for %s", url)
	}
}
