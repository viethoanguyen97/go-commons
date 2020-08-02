package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v //set P to be the pointer of v
	p.X = 1000
	fmt.Println(v)
}
