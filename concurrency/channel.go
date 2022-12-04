package main

import (
	"fmt"
)

func main() {
	//var chan1 chan int
	//
	//fmt.Println("channel 1 value", chan1)
	//fmt.Printf("Type %T\n", chan1)
	//
	//chan2 := make(chan int)
	//
	//fmt.Println("channel 2 value", chan2)
	//fmt.Printf("Type %T\n", chan2)

	//fmt.Println("Start main fund")
	//
	//ch := make(chan int)
	//
	//go send(ch)
	//
	//go receive(ch)
	//
	//time.Sleep(100 * time.Millisecond)
	//
	//fmt.Println("end of main func")

	//ch := make(chan int)
	//
	//go send(ch)
	//
	//for {
	//	result, ok := <-ch
	//
	//	if !ok {
	//		fmt.Println("Channel Close", ok)
	//		break
	//	}
	//	fmt.Println("channel open ", result)
	//
	//}

	//ch := make(chan string, 1)
	//
	//ch <- "Go lang"
	//
	//message := <-ch
	//fmt.Println(message)
	//
	//fmt.Println("len :", len(ch))
	//fmt.Println("cap :", cap(ch))
	//
	//ch <- "hello"
	//
	//fmt.Println("len :", len(ch))
	//fmt.Println("cap :", cap(ch))

	ch := make(chan int, 5)

	process(&ch)

	fmt.Println(ch)
}

func process(ch *chan int) {
	*ch <- 2

	s := <-*ch

	*ch = nil

	fmt.Println("s :", s)
	fmt.Println("ch :", *ch)
}

func send(ch chan int) {
	//ch <- 50

	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch)
}

func receive(ch chan int) {
	fmt.Println("value ", <-ch)
}
