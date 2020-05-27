package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	var myI32 int32 = 19
	bockChan := make(chan bool)
	quitChan := make(chan bool)

	go func() {
		for {
			select {
			case <-quitChan:
				fmt.Println("a.................................................")
				return
			default:
				atomic.StoreInt32(&myI32, 19)
				time.Sleep(100 * time.Millisecond)
				close(bockChan)
				bockChan = make(chan bool)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-quitChan:
				fmt.Println("b.................................................")
				return
			default:
				p := atomic.AddInt32(&myI32, -1)
				if p < 0 {
					<-bockChan
				} else {
					fmt.Printf("atomic.AddInt32(...) === %d\n", p)
				}
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
