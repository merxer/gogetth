package main

import (
	"fmt"

)

func init() {
	fmt.Println("Initial")
}

func main() {

	counter := counterFactory()
	printCount(counter)
	printCount(counter)
	printCount(counter)
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

func printCount(fn func() int ) {
	fmt.Println(fn())
}
