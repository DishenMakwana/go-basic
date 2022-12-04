package main

import "fmt"

func main() {

	var score int = 34

	p := &score

	if p != nil {
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var x int = 100

	var ptr1 *int = &x
	var ptr2 **int = &ptr1

	fmt.Println("Value of x:", x)
	fmt.Println("Value of ptr1:", ptr1)
	fmt.Println("Value of ptr2:", *ptr2)
}
