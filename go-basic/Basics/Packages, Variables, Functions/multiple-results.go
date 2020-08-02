package main

import "fmt"

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	a := "ab"
	b := "ba"

	a, b = swap(a, b)
	fmt.Println(a, b)
}
