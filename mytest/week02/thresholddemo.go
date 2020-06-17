package main

import (
	"flag"
	"fmt"
	"sync/atomic"
	"time"
)

var qps int = 1000
var user int = 3000
var threshold int32 = 1
var quitCHan chan bool = make(chan bool)

func main() {
	flag.IntVar(&user, "user", 3000, "")
	flag.Parse()
	go func() {
		ticker := time.NewTicker(30*time.Second)
		for {
			select {
			case <-ticker.C:
				atomic.StoreInt32(&threshold, -1)
				time.Sleep(5*time.Second)
				atomic.StoreInt32(&threshold, 1)
				close(quitCHan)
				quitCHan = make(chan bool)
			default:
				continue
			}
		}
	}()
	for i := 0; i < user; i++ {
		go func(j int) {
			for {
				if atomic.LoadInt32(&threshold) < 0 {
					<-quitCHan
				} else {
					fmt.Printf("the running goroutine【%d】\n", j)
				}

			}

		}(i)
	}

	time.Sleep(2 * time.Minute)
}
