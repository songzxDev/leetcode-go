package main

import "fmt"

type My struct {
	name string
}

func (my My) Say() string {
	fmt.Println("where is me")
	return "where is me"
}
func (my My) Hello() string  {
	return "hello word"
}
func (my My) hello() string  {
	return "hello word"
}

func main() {
	my := My{name: "abc"}
	fmt.Println(my.Say())
	fmt.Println(my.Hello())
	fmt.Println(my.hello())
}
