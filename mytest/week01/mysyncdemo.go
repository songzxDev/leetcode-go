package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var i32 int32 = int32(97)
	var o32 int32 = i32
	for i := 0; i < int(o32); i++ {
		go func(k int) {
			myi32 := atomic.AddInt32(&i32, -1)
			fmt.Printf("i的值是：%d，myi32的值是：%d\n", k, myi32)
		}(i)
	}
	for {
		if atomic.LoadInt32(&i32) == 0 {
			fmt.Printf("myi32的值是：%d\n", i32)
			func() {
				for i := 0; i < 1; i++ {
					fmt.Println("abc")
					func() {
						for j := 1; j < 100; j++ {
							if j == 99 {
								return
							}
						}
					}()
					return
				}
			}()
			return
		}
		time.Sleep(time.Millisecond * 1)
	}
	fmt.Println("Main exit.")
}
