package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var myString string
	fmt.Scanln(&myString)
	fmt.Println(myString)

	var name string = "dishen"
	var a, _ = fmt.Println(name)
	fmt.Println(a)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	myName, _ := reader.ReadString('\n')

	fmt.Println(myName)

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter user rating: ")
	myRating, _ := (reader2.ReadString('\n'))

	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64)

	fmt.Println(myNumRating + 2)
}
