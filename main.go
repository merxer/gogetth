package main

import (
	"fmt"

	"github.com/merxer/demo/bank"
)

func init() {
	fmt.Println("Initial")
}

func main() {
	println("Hello, " + bank.Title, func() string {
		return "KTB"
	}())

	counter := counterFactory()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func init() {
	fmt.Println("Initial 2")
}

func counterFactory() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
