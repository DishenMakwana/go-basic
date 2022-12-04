package main

import "fmt"

func main() {
	var numbers [3]string
	numbers[0] = "one"
	numbers[1] = "two"
	numbers[2] = "three"

	fmt.Println(numbers)

	var colors = [4]string{"Red", "Blue", "Green", "Yellow"}

	fmt.Println(colors)
}
