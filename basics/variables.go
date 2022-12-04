package main

import "fmt"

func main() {
	var batman string = "Bruce Wayne"
	fmt.Println(batman)

	var superman string
	superman = "Clark Kent"
	fmt.Println(superman)

	thor := "Thor Odinson"
	fmt.Println(thor)

	// thorRating := 3.
	// fmt.Printf("%v, %T", thorRating, thorRating)

	var ironMan, captionAmerica string = "Tony Stark", "Steve Rogers"
	fmt.Println(ironMan, captionAmerica)

	var defaultValue int
	fmt.Println(defaultValue)

	var (
		spiderMan    = "Peter Parker"
		blackPanther = "T'Challa"
	)
	fmt.Println(spiderMan, blackPanther)
}
