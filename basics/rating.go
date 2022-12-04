package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ = reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating: ")
	userRating, _ = reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	fmt.Printf("%v, %v, %v", name, myNumRating, time.Now().Format(time.Stamp))
	fmt.Println()

	if myNumRating == 5 {
		fmt.Println("Bonus for team member")
	} else if myNumRating == 4 || myNumRating == 3 {
		fmt.Println("Good to go for team member")
	} else {
		fmt.Println("Need to improve")
	}
}
