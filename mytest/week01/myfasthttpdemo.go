package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	maxIdleConnDuration time.Duration
)

func myTestRun(cli *fasthttp.Client) {
	if cli == nil {
		cli = &fasthttp.Client{
			MaxConnsPerHost: 2000,
			// DefaultMaxIdleConnDuration 是在空闲的 keep-alive 连接被关闭前默认的持续时间
			// const DefaultMaxIdleConnDuration = 10 * time.Second
			MaxIdleConnDuration: maxIdleConnDuration,
		}
	}

	requ := fasthttp.AcquireRequest()

	requ.Header.SetMethod("GET")
	requ.Header.SetContentType("application/json")
	resp := fasthttp.AcquireResponse()
	requ.SetRequestURI("http://10.168.11.12:8009/getmockjson")
	err := cli.DoTimeout(requ, resp, 10*time.Second)
	if err != nil {
		fmt.Printf(err.Error())
		fasthttp.ReleaseRequest(requ)
		fasthttp.ReleaseResponse(resp)
		return
	}
	fmt.Printf(string(resp.Body()))
	fasthttp.ReleaseRequest(requ)
	fasthttp.ReleaseResponse(resp)
}

func main() {
	maxIdleConnDuration = 60 * time.Second

	for i := 0; i < 5000; i++ {
		go myTestRun(nil)
	}
	time.Sleep(maxIdleConnDuration)
}
