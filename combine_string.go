package main

import "fmt"

func combineString(strs ...string) string {
	var result string
	for _, str := range strs {
		if result != "" {
			result = result + "," + str
		} else {
			result = str
		}
	}
	return result
}

func swapString(first, second string) (string, string) {
	return second, first
}

func main() {
	fmt.Println(combineString("Hello", "World", "Third", "4th"))
	fmt.Println(swapString("A", "B"))
}
