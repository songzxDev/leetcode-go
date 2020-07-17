package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	threshold int32         = 1
	bookChan  chan bool     = make(chan bool)
	workerNum int           = 100
	roundTime time.Duration = 30 * time.Second
	runTime   time.Duration = 4 * roundTime
)

func working(g int) {
	for {
		if atomic.LoadInt32(&threshold) < 0 {
			fmt.Printf("goroutine【%d】 sleep at %v......\n", g, time.Now().UnixNano())
			<-bookChan
		} else {
			fmt.Printf("goroutine【%d】 at %v running......\n", g, time.Now().UnixNano())
		}
	}
}
func main() {
	go func() {
		ticker := time.NewTicker(roundTime)
		for {
			select {
			case <-ticker.C:
				atomic.StoreInt32(&threshold, -1)
				fmt.Println(roundTime/ 2)
				time.Sleep(roundTime / 2)
				atomic.StoreInt32(&threshold, 1)
				close(bookChan)
				bookChan = make(chan bool)
			}
		}
	}()
	for i := 0; i < workerNum; i++ {
		go working(i)
	}

	time.Sleep(runTime)
}
