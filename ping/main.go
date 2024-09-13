package main

import (
	"fmt"
	"os/exec"
	"sync"
)

type PingResult struct {
	IP     string
	Status string
}

func main() {
	// 需要 Ping 的 IP 列表
	targets := []string{
		"8.8.8.8", "8.8.4.4", "1.1.1.1", "192.168.1.1", "invalid_ip",
	}

	var wg sync.WaitGroup

	// 用于传递 IP 的 channel
	ipChan := make(chan string, len(targets))
	// 用于传递 Ping 结果的 channel
	resultChan := make(chan PingResult, len(targets))
	workerCount := 3 // 控制并发数量

	// 启动 worker Goroutines
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(ipChan, resultChan, &wg)
	}
	// 将所有的目标 IP 发送到 ipChan
	go func() {
		for _, ip := range targets {
			ipChan <- ip
		}
		close(ipChan)
	}()

	// 等待所有的 worker 完成
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 处理结果
	for result := range resultChan {
		fmt.Printf("Ping %s: %s\n", result.IP, result.Status)
	}
}

// worker 是每个 Goroutine 执行的 Ping 任务
func worker(ips <-chan string, results chan<- PingResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for ip := range ips {
		status := "unreachable"
		if ping(ip) {
			status = "reachable"
		}
		results <- PingResult{IP: ip, Status: status}
	}
}

func ping(ip string) bool {
	cmd := exec.Command("ping", "-c", "1", ip)
	err := cmd.Run()
	return err == nil
}
