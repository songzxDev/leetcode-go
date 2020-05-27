package main

import (
	"github.com/myzhan/boomer"
	"github.com/valyala/fasthttp"
	"time"
)

var client *fasthttp.Client

func myPlay() {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.Header.SetRequestURI("http://10.168.11.12:8009/getmockjson")
	req.Header.SetContentType("application/json")
	resp := fasthttp.AcquireResponse()

	startTime := time.Now()
	err := client.DoTimeout(req, resp, 10*time.Second)
	elapsed := time.Since(startTime)
	if err != nil {
		boomer.RecordFailure("http", "unkown", elapsed.Nanoseconds()/int64(time.Millisecond), err.Error())
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
		return
	}
	boomer.RecordSuccess("http", "ok", elapsed.Nanoseconds()/int64(time.Millisecond), int64(len(resp.Body())))
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)

}
func main() {
	client = &fasthttp.Client{
		MaxConnsPerHost:     2000,
		MaxIdleConnDuration: 30 * time.Second,
	}
	task := &boomer.Task{
		Name:   "worker",
		Weight: 10,
		Fn:     myPlay,
	}
	boomer.Run(task)
}
