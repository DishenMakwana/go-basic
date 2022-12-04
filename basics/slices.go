package main

import (
	"fmt"
	"sort"
)

func main() {
	var things = []string{"one", "two", "three"}

	things = append(things, "four")

	fmt.Println(things)

	things = append(things[1 : len(things)/2])

	fmt.Println(things)

	heros := make([]string, 3, 3)

	heros[0] = "Superman"
	heros[1] = "Batman"
	heros[2] = "Spiderman"

	fmt.Println(heros)

	heros = append(heros, "Ironman")

	fmt.Println(heros)

	myints := []int{4, 9, 2, 7, 3, 8, 6, 1, 5}

	// isSorted := sort.IntsAreSorted(myints)

	// sortedInts := sort.Sort(myints)
	sort.Sort(sort.IntSlice(myints))

	fmt.Println(myints)

	mainSlice := []int{10, 20, 30, 40, 50, 60}

	newSlice := mainSlice[:]

	fmt.Println(newSlice)

	var slices = make([]int, 4, 7)

	fmt.Println(slices, len(slices), cap(slices))

	for _, v := range mainSlice {
		fmt.Println(v)
	}

	slices = append(slices, 30, 40, 50, 60)

	fmt.Println(slices, len(slices), cap(slices))

	// sort.Ints(slices)
	sort.Ints(slices)

	slices = append(slices, mainSlice...)
}
