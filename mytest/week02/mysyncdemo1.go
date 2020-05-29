package main

import (
	"log"
	"os"
	"os/exec"
	"sync/atomic"
	"time"
)

func init() {
	myCmd := exec.Command("echo", ".> G:\\mygolandlog\\myrunlog2.log")
	myCmd.Stdout = os.Stdout
	_ = myCmd.Run()
	file := "G:\\mygolandlog\\myrunlog2.log"
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
	qps, users := 3000, 1000
	var threshold int32 = int32(qps)
	quitChan, bookChan := make(chan bool), make(chan bool)

	for i := 0; i < users; i++ {
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
						p := atomic.AddInt32(&threshold, -1)
						if p < 0 {
							<-bookChan
						} else {
							log.Printf("atomic.AddInt32===%d", p)
						}
					}
				}
			}()
		}
	}
	go func() {
		for {
			select {
			case <-quitChan:
				return
			default:
				atomic.StoreInt32(&threshold, int32(qps))
				time.Sleep(1 * time.Second)
				close(bookChan)
				bookChan = make(chan bool)
			}
		}
	}()

	time.Sleep(2 * time.Minute)
}
