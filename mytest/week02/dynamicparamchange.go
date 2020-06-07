package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	var myUrl string = "abc"
	quitChan := make(chan bool)
	bookChan := make(chan bool)
	var pivot int32
	go func() {
		timer := time.NewTicker(50 * time.Second)
		for {
			select {
			case <-timer.C:
				atomic.StoreInt32(&pivot, -1)
				myUrl = "bcd"
				time.Sleep(10 * time.Second)
				close(bookChan)
				bookChan = make(chan bool)
				atomic.StoreInt32(&pivot, 1)
				time.Sleep(2 * time.Second)
				atomic.StoreInt32(&pivot, -1)
				time.Sleep(10 * time.Second)
				myUrl = "abc"
				close(bookChan)
				bookChan = make(chan bool)
				atomic.StoreInt32(&pivot, 1)
			default:
				continue
			}
		}
	}()
	for i := 0; i < 100; i++ {
		go func() {
			for {
				select {
				case <-quitChan:
					return
				default:
					if atomic.AddInt32(&pivot, 0) < 0 {
						<-bookChan
					}
					fmt.Print(myUrl)

				}
			}

		}()
	}

	time.Sleep(5 * time.Minute)
}
