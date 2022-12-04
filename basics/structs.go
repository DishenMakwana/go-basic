package main

import "fmt"

type User struct {
	Name  string
	Email string
	age   int
}

func main() {
	rob := User{"Rob", "rob@gmail.com", 34}

	fmt.Printf("%+v\n", rob)

	var sam = new(User)
	sam.Name = "Sam"
	sam.Email = "sam@gmail.com"
	sam.age = 30

	fmt.Printf("%+v\n", sam)

	var tobby = &User{"Tobby", "tobby@gmail.com", 25}

	fmt.Printf("%+v\n", tobby)
}
