package main

import "fmt"

func sayHello() {
	fmt.Println("hello world")
}
func sayHi() {
	fmt.Println("hiiiiii")
}

func main() {
	fmt.Println("go routine")
	sayHi()
	sayHello()
}
