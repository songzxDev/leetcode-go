package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func myTestRun() {
	cli := &fasthttp.Client{
		MaxConnsPerHost: 2000,
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
	myTestRun()
}
