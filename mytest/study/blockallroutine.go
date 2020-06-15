package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	threshold, bookChan := int32(0), make(chan bool)
	for i := 0; i < 100; i++ {
		go func(j int) {
			for {
				if atomic.LoadInt32(&threshold) < 0 {
					fmt.Printf("携程【%d】阻塞中......\n", j)
					<-bookChan
				} else {
					fmt.Printf("携程【%d】运行中%v\n", j, time.Now().UnixNano())
				}
			}
		}(i)
	}
	go func() {
		timer := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-timer.C:
				atomic.StoreInt32(&threshold, -1)
				time.Sleep(5 * time.Second)
				atomic.StoreInt32(&threshold, 1)
				close(bookChan)
				bookChan = make(chan bool)
			default:
				continue
			}

		}
	}()
	time.Sleep(2 * time.Minute)
}
