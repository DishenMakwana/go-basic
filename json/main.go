package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//type Person struct {
//	Name     string `json:"firstName"`
//	Age      int64
//	Location string `json:"location,omitempty"`
//}

type Person struct {
	Name     string
	Age      int64
	Location string
}

func main() {
	//person := Person{
	//	Name: "kim",
	//	Age:  23,
	//}
	//
	//personArray, err := json.Marshal(person)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(personArray))

	//j := []byte(`{"name": "Kim", "age": 32, "location": "abd"}`)
	//
	//var p Person
	//
	//err := json.Unmarshal(j, &p)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(p)

	//jsonStream := `{"name": "Kim", "age": 32, "location": "abd"}{"name": "Rajan", "age": 24, "location": "abd"}{"name": "Dishen", "age": 22, "location": "abd"}`
	//
	//reader := strings.NewReader(jsonStream)
	//
	//writer := os.Stdout
	//
	//decoder := json.NewDecoder(reader)
	//
	//encoder := json.NewEncoder(writer)
	//
	//for {
	//	var v map[string]interface{}
	//
	//	if err := decoder.Decode(&v); err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	for k := range v {
	//		if k == "Age" {
	//			delete(v, k)
	//		}
	//	}
	//
	//	if err := encoder.Encode(&v); err != nil {
	//		log.Println(err)
	//	}
	//}

	file, err := os.Open("config.json")

	if err != nil {
		log.Fatal(err)
	}

	conf := new(Configuration)

	err = json.NewDecoder(file).Decode(conf)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Run server ........")

	//web.RunWeb(conf.ServerPort)
}

type Configuration struct {
	ServerPort string `json:"serverport"`
}
