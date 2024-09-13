package main

import (
	"fmt"
	"scraper/scraper"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
	}
	results := make(chan string, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go scraper.Scrape(url, &wg, results) // goroutine获取网页内容
	}
	go func() {
		wg.Wait() // 等待所有goroutine完成
		close(results)
	}()
	for result := range results {
		fmt.Println(result)
		// URL: https://www.github.com, Length: 253413
		// URL: https://www.stackoverflow.com, Length: 127409
		// URL: https://www.google.com, Length: 22212

		// URL: https://www.google.com, Length: 22191
		// URL: https://www.github.com, Length: 253434
		// URL: https://www.stackoverflow.com, Length: 127413
	}
}
