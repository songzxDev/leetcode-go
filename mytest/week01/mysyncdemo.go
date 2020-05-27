package main

import (
	"log"
	"os"
	"sync/atomic"
	"time"
)
func init() {
	file := "G:\\mygolandlog\\myrunlog1.log"
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

	var myI32 int32
	bockChan := make(chan bool)
	quitChan := make(chan bool)
	// 1000用户在线，1000 qps
	users, qps := 3000, int32(2000)
	atomic.StoreInt32(&myI32, qps)
	go func() {
		for {
			select {
			case <-quitChan:
				return
			default:
				atomic.StoreInt32(&myI32, qps)
				time.Sleep(20 * time.Millisecond)
				close(bockChan)
				bockChan = make(chan bool)
			}
		}
	}()
	for i := 0; i < users; i++ {
		time.Sleep(1 * time.Millisecond)
		select {
		case <-quitChan:
			return
		default:
			go func() {
				for {
					select {
					case <-quitChan:
						return
					default:
						p := atomic.AddInt32(&myI32, -1)
						if p < 0 {
							<-bockChan
						} else {
							log.Printf("atomic.AddInt32===%d", p)
						}
					}
				}
			}()
		}

	}

	time.Sleep(60 * time.Second)
}
