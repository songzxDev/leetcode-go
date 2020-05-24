package main

import (
	"errors"
	"fmt"
)

func main() {
	myChan := make(chan int, 8)
	defer func() {
		fmt.Println("Run defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic %s\n", p)
		}
		_, ok := <-myChan
		if ok {
			fmt.Printf("管道容量：%d\n", len(myChan))
			close(myChan)
		} else {
			fmt.Println(len(myChan))
		}
		fmt.Println("Exit defer function.")
	}()
	for i := 0; i < 8; i++ {
		myChan <- i + 1
	}
	fmt.Printf("管道初始化后容量：%d\n", len(myChan))
	elmt := <- myChan
	panic(errors.New("something wrong"))
	fmt.Println(elmt)
	fmt.Println("Exit main function.")
}
