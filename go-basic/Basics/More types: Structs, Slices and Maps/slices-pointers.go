package main

import "fmt"

func main() {
	names := [4]string{
		"A", "B", "C", "D",
	}

	fmt.Println(names) //=> a, b, c, d

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b) //=> [a, b] [b, c]

	b[0] = "XXX"
	fmt.Println(a, b)  //=>[a, XXX] [XXX 1]
	fmt.Println(names) //=>[a XXX c d]
}
