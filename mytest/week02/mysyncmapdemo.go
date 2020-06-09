package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"
)

func main() {
	myBockChan := make(chan bool, 100)
	myQuitChan, threshold := make(chan bool), int32(1)
	go func() {
		timer1 := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-timer1.C:
				atomic.StoreInt32(&threshold, -1)
				for i := 0; i < 100; i++ {
					myBockChan <- true
				}
				time.Sleep(3000 * time.Millisecond)
				atomic.StoreInt32(&threshold, 1)
				close(myQuitChan)
				myQuitChan = make(chan bool)
			default:
				continue
			}
		}
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for {
				if atomic.LoadInt32(&threshold) < 0 {
					<-myQuitChan
				}
				select {
				case <-myBockChan:
					fmt.Println(strconv.FormatInt(time.Now().UnixNano(), 10))
				default:
					continue
				}
			}
		}()
	}

	time.Sleep(5 * time.Minute)
}
