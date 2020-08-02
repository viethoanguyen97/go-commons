package main

import (
	"fmt"
)

var i, j int = 1, 2

//ii := 3 -> error
func main() {
	c, python, java := true, false, "no"
	fmt.Println(i, j, c, python, java)
}
