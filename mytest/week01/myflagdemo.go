package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}
func main() {
	flag.Parse()
	fmt.Printf("hello, %s!\n", name)
	value, ok := interface {}(name).([]string)
	fmt.Println(value, ok)
}
