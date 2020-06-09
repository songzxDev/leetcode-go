package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			for (time.Now().UnixNano() / 1e6) != (time.Now().Unix() * 1000) {
			}
			fmt.Println(time.Now().UnixNano())
		}
	}()

	time.Sleep(3 * time.Minute)
}
