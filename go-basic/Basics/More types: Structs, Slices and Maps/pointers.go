package main

import "fmt"

func main() {
	i, j := 42, 4701

	p := &i         //Point to i
	fmt.Println(*p) //Read i through the pointer
	*p = 21         //set i through the pointer
	fmt.Println(i)  //see the new value of i

	p = &j         //Point to j
	*p = *p / 37   //divide j through the pointer
	fmt.Println(j) //see the new value of j

}
