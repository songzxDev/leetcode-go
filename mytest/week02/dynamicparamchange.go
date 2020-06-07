package main

import (
	"log"
	"os"
	"os/exec"
	"sync/atomic"
	"time"
)

func init() {
	myCmd := exec.Command("echo", "''> D:\\tmp\\mydynamicparam.log")
	myCmd.Stdout = os.Stdout
	_ = myCmd.Run()
	file := "D:\\tmp\\mydynamicparam.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}
func main() {
	bookChan, myPivot, myUrl, quitChan := make(chan bool), int32(0), "abc", make(chan bool)
	mySleep := func(t time.Duration, doSomething func()) {
		atomic.StoreInt32(&myPivot, -1)
		doSomething()
		time.Sleep(t)
		atomic.StoreInt32(&myPivot, 1)
		close(bookChan)
		bookChan = make(chan bool)
	}
	go func() {
		timer := time.NewTicker(51 * time.Second)
		for {
			select {
			case <-timer.C:
				// 所有 goroutine 暂停运行10s
				mySleep(10*time.Second, func() {
					myUrl = "cba"
				})
				time.Sleep(1100 * time.Millisecond)
				mySleep(10*time.Second, func() {
					myUrl = "abc"
				})
			default:
				continue
			}
		}

	}()

	for i := 0; i < 64000; i++ {
		go func(j int) {
			for {
				select {
				case <-quitChan:
					return
				default:
					if atomic.AddInt32(&myPivot, 0) < 0 {
						<-bookChan
					}
					log.Printf("%s+%d", myUrl, j)
				}
			}
		}(i + 1)
	}
	time.Sleep(5 * time.Minute)
}
