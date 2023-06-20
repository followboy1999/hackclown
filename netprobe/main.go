package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var operation string
	var addr string
	var concurrency int
	flag.StringVar(&operation, "m", "scanport", "scanport or httprobe")
	flag.StringVar(&addr, "t", "127.0.0.1", "target")
	flag.IntVar(&concurrency, "c", 100, "concurrency")
	flag.Parse()

	// 定义要扫描的端口范围
	startPort := 1
	endPort := 65535

	// 定义要扫描的协议
	protocols := []string{"http", "https"}

	// 创建等待组
	var wg sync.WaitGroup

	// 创建通道
	jobs := make(chan int, concurrency)

	// 循环扫描端口
	for port := startPort; port <= endPort; port++ {
		// 添加任务到通道
		jobs <- port

		// 如果通道已满，等待任务完成
		if len(jobs) == concurrency {
			wg.Add(concurrency)
			go func() {
				for i := 0; i < concurrency; i++ {
					port := <-jobs
					if operation == "scanport" {
						scanPort(port, addr)
					} else {
						httprobe(port, addr, protocols)
					}
					wg.Done()
				}
			}()
			wg.Wait()
		}
	}
	// 执行剩余任务
	for len(jobs) > 0 {
		wg.Add(1)
		go func() {
			port := <-jobs
			if operation == "scanport" {
				scanPort(port, addr)
			} else {
				httprobe(port, addr, protocols)
			}
			wg.Done()
		}()
	}

	// 定义超时时间为 5 秒
	timeout := time.After(5 * time.Second)
	// 在超时时间内完成任务，输出结果
	select {
	case <-timeout:
		fmt.Println("Timeout")
		os.Exit(1)
	default:
		fmt.Println("Task completed")
		os.Exit(0)
	}
	// 等待所有任务完成
	wg.Wait()
}

func scanPort(port int, addr string) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", addr, port), time.Second)
	if err != nil {
		return
	}
	defer conn.Close()
	if conn != nil {
		fmt.Printf("%s %d open\n", addr, port)
	}
}

func httprobe(port int, addr string, protocols []string) {
	for _, protocol := range protocols {
		if isProtocol(addr, port, protocol) {
			fmt.Printf("%s %d %s\n", addr, port, protocol)
			return
		}
	}
}

func isProtocol(addr string, port int, protocol string) bool {
	if protocol == "http" {
		client := http.Client{Timeout: 1 * time.Second}                   //添加超时
		resp, err := client.Get(fmt.Sprintf("http://%s:%d/", addr, port)) //使用client
		if err == nil && (resp.StatusCode == 200 || resp.StatusCode == 302) {
			return true
		}

	} else if protocol == "https" {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", addr, port), time.Second)
		if err != nil {
			return false
		}
		defer conn.Close()
		tlsConn := tls.Client(conn, &tls.Config{InsecureSkipVerify: true})
		tlsConn.SetDeadline(time.Now().Add(time.Second * 1)) //1秒超时
		err = tlsConn.Handshake()
		if err != nil {
			return false
		}
		_, err = tlsConn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
		if err != nil {
			return false
		}
		resp, err := http.ReadResponse(bufio.NewReader(tlsConn), nil)
		if err != nil {
			return false
		}
		if resp.StatusCode == 200 {
			return true
		}
	}
	return false
}
