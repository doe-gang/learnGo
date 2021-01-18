package main

import (
	"fmt"
	"learnGo/hanoi"
	"learnGo/seq"
)

func main() {
	fmt.Println("Hello World!!")
	hanoi.Hanoi(3)

	fmt.Println(seq.Fib(7))
}
