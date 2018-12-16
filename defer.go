package main

import "fmt"

func main() {
	doSomething(4)
}

func doSomething(n int) {
	defer fmt.Println(n)
	defer func(m int) {
		fmt.Println(m)
	}(n)
	n = n * n
	fmt.Println(n)
}
