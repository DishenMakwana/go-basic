package main

import (
	"flag"
	"fmt"
	"github.com/dishenmakwana/protocol_buffer/information/protocol"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	option := flag.String("admin", "server", "Communication")

	flag.Parse()

	switch *option {
	case "client":
		runClient()
	case "server":
		runServer()
	}
}

func runClient() {
	person := protocol.Person{
		Id:   1,
		Name: "Dishen",
		Age:  22,
	}

	data, err := proto.Marshal(&person)

	if err != nil {
		log.Fatal("marshal err: ", err.Error())
	}

	sendData(data)
}

func runServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:8085")

	if err != nil {
		log.Fatal("listener err: ", err.Error())
	}

	for {
		connection, err := listener.Accept()

		if err != nil {
			log.Fatal("listener err:", err.Error())
		}

		go func(c net.Conn) {
			defer c.Close()

			data, err := ioutil.ReadAll(connection)

			if err != nil {
				log.Fatal(err.Error())
			}

			person := &protocol.Person{}

			proto.Unmarshal(data, person)

			fmt.Println("Person: ", person)
		}(connection)
	}
}

func sendData(data []byte) {
	connection, err := net.Dial("tcp", "127.0.0.1:8085")

	if err != nil {
		log.Fatal("connection err: ", err.Error())
	}

	defer connection.Close()

	connection.Write(data)
}
