package main

import "fmt"

func main() {
	var m map[string]string

	m = make(map[string]string)

	if m == nil {
		fmt.Println("it's nil value")
	}

	m["a"] = "ant"
	m["b"] = "bat"
	m["c"] = "cat"

	delete(m, "b")
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
