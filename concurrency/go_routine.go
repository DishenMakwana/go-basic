package main

import (
	"fmt"
	"time"
)

func main() {
	// 1
	//go display()
	//
	//fmt.Println("hello")
	//
	//time.Sleep(1 * time.Second)
	//fmt.Println("Good bye")

	// 2
	//fmt.Println("Go routine start ..........")
	//
	//go name()
	//
	//go id()
	//
	//time.Sleep(1000 * time.Millisecond)
	//fmt.Println("Main go routine end")

	// 3
	//fmt.Println("Hello from main function")
	//
	//go func() {
	//	fmt.Println("Inside go routine")
	//}()
	//
	//time.Sleep(1 * time.Second)
}

func name() {
	array := [4]string{"abc", "xyz", "pqr", "mno"}

	for i := 0; i < len(array); i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(array[i])
	}
}

func id() {
	array := [4]int{1, 2, 3, 4}

	for i := 0; i < len(array); i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(array[i])
	}
}

func display() {
	go display2()
	fmt.Println("Inside go routine display")
}

func display2() {
	fmt.Println("Inside go routine display2")
}
