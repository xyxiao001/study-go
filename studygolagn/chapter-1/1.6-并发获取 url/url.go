// go run url.go https://taobao.com  https://baidu.com  http://xyxiao.cn
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// 并发执行  goroutine
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// 接受 channel ch
		fmt.Printf(<-ch)
	}
	fmt.Printf("花费时间总额: %2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// 发送信息
		ch <- fmt.Sprintf("%s\n", err)
		return
	}
	bytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("当前读取网址 %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, bytes, url)
}
