package main

import (
	"fmt"
)

func main() {
	start := 1
	things := []string{"maleta", "bolso", "boleto", "vaso", "casa"}

	fmt.Println(things)

	for i := 0; i < len(things); i++ {
		fmt.Println(i, things[i])
	}

	for i := range things {
		fmt.Println(i, things[i])
	}

	for start < 100 {
		fmt.Println(start)

		if start == 32 {
			goto superman // goto is not a good practice
		}

		start += start
	}

superman:
	fmt.Println("Superman")
}
