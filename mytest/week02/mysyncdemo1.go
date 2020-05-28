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

	// 定义用于控制流程流转的channel
	quitChan, bockChan := make(chan bool), make(chan bool)

	// 定义用于计数控制的原子类型变量
	var myI32 int32
	users, qps := 1000, int32(2000)
	atomic.StoreInt32(&myI32, qps)
	// 恢复计数的控制分支

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

	time.Sleep(2 * time.Minute)
}
