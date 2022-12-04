package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var str string = "hello"
	fmt.Print(str)

	str1 := "world"
	num := 123
	fmt.Printf("%s", str1)
	fmt.Printf("%d", num)

	fmt.Println(num)

	var name, value = "Go", "Programming Language"

	result := fmt.Sprint(name, " is a ", value, "\n")
	result2 := fmt.Sprintf("%s is a %s \n", name, value)
	result3 := fmt.Sprintln(name, " is a ", value, "\n")

	io.WriteString(os.Stdout, result)
	io.WriteString(os.Stdout, result2)
	io.WriteString(os.Stdout, result3)

	number, err := fmt.Fprint(os.Stdout, result)
	fmt.Println(number, " bytes written")
	fmt.Println(err)

	var name1 string
	var age int

	fmt.Scan(&name1)
	fmt.Scan(&age)

	fmt.Println(name1, age)
}
