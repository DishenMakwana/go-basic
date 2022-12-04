package main

import (
	"fmt"
	"sort"
)

func main() {
	score := make(map[string]int)

	score["dishen"] = 10
	score["james"] = 20

	fmt.Println(score)

	getDishenScore := score["dishen"]

	fmt.Println(getDishenScore)

	delete(score, "dishen")

	fmt.Println(score)

	for k, v := range score {
		fmt.Println(k, v)
	}

	var simpleMap = map[int]string{}

	simpleMap[1] = "one"
	simpleMap[2] = "two"
	simpleMap[3] = "three"
	simpleMap[4] = "four"
	simpleMap[5] = "five"

	fmt.Println(simpleMap)

	keys := make([]int, len(simpleMap))
	values := make([]string, len(simpleMap))

	for i := range simpleMap {
		keys = append(keys, i)
	}

	for i := range simpleMap {
		values = append(values, simpleMap[i])
	}

	sort.Ints(keys)
	sort.Strings(values)

	for _, v := range keys {
		fmt.Println(v, simpleMap[v])
	}

	for _, v := range values {
		fmt.Println(v)
	}

}
