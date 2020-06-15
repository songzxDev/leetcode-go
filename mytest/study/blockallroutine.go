package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	threshold int32         = 1
	bookChan  chan bool     = make(chan bool)
	takeTime  time.Duration = 10 * time.Second
)

func main() {
	for i := 0; i < 100; i++ {
		go func(j int) {
			for {
				if atomic.LoadInt32(&threshold) < 0 {
					fmt.Printf("goroutine【%d】阻塞中%v......\n", j, time.Now().UnixNano())
					<-bookChan
				} else {
					fmt.Printf("goroutine【%d】运行中%v......\n", j, time.Now().UnixNano())
				}
			}
		}(i)
	}
	go func() {
		ticker := time.NewTicker(takeTime)
		for {
			select {
			case <-ticker.C:
				atomic.StoreInt32(&threshold, -1)
				time.Sleep(takeTime / 2)
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
