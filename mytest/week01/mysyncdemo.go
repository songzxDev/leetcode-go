package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {


	var myI32 int32 = 0
	bockChan := make(chan bool)
	quitChan := make(chan bool)

	go func() {
		for {
			select {
			case <-quitChan:
				fmt.Println("a.................................................")
				return
			default:
				fmt.Printf("atomic.StoreInt32(...) === %d\n", myI32)
				atomic.StoreInt32(&myI32, 19)
				time.Sleep(500 * time.Millisecond)
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
				fmt.Printf("atomic.AddInt32(...) === %d\n", myI32)
				if p < 0 {
					<-bockChan
				}
			}
		}
	}()

	time.Sleep(20 * time.Second)
	quitChan <- true
	quitChan <- true
	time.Sleep(10 * time.Second)
}
