package main

import (
	"fmt"
	"math"
)

func main() {
	superman()

	result := multiplyIt(2, 3)

	fmt.Println(result)

	myResult := adder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(myResult)

	square := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(square(9))

	var a, b int = 100, 200

	swap(&a, &b)

	fmt.Println(a, b)

}

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func superman() {
	fmt.Println("Superman")
}

func multiplyIt(v1 int, v2 int) int {
	return v1 * v2
}

func adder(values ...int) int {
	sum := 0

	for _, v := range values {
		sum += v
	}

	return sum
}
