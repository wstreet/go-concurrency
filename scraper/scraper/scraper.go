package scraper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func Scrape(url string, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()            // 表示一个 Goroutine 完成任务，等价于 Add(-1)
	resp, err := http.Get(url) // 发送 HTTP GET 请求
	if err != nil {
		ch <- fmt.Sprintf("Failed to fetch %s: %s", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // 读取 HTTP 响应体
	if err != nil {
		ch <- fmt.Sprintf("Failed to read response body from %s: %s", url, err)
		return
	}

	ch <- fmt.Sprintf("URL: %s, Length: %d", url, len(body))
}
