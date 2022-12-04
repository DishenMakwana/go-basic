package main

import (
	"fmt"
	"github.com/dishenmakwana/protocol_buffer/library/protocol"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	book := &protocol.Book{
		Name: "Go Programming language",
		Isbn: 123456,
	}

	data, err := proto.Marshal(book)

	if err != nil {
		log.Fatal("marshal err: ", err.Error())
	}

	newBook := &protocol.Book{}

	err = proto.Unmarshal(data, newBook)

	if err != nil {
		log.Fatal("unmarshal err: ", err.Error())
	}

	fmt.Println("Name: ", newBook.GetName())
	fmt.Println("Isbn: ", newBook.GetIsbn())
}
