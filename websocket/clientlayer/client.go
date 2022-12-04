package main

import (
	"bufio"
	"fmt"
	"github.com/dishenmakwana/websocket/socket"
	"golang.org/x/net/websocket"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	connection, err := websocket.Dial("ws://127.0.0.1:8085", "", createIP())

	if err != nil {
		log.Fatal(err.Error())
	}

	defer connection.Close()

	go receiveMessage(connection)

	SendMessage(connection)
}

func createIP() string {
	var ip [4]int

	for i := 0; i < len(ip); i++ {
		rand.Seed(time.Now().UnixNano())

		ip[i] = rand.Intn(256)
	}

	return fmt.Sprintf("http://%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func receiveMessage(conncetion *websocket.Conn) {
	for {
		var message socket.Message

		err := websocket.JSON.Receive(conncetion, message)

		if err != nil {
			log.Fatal("Error in receive data: ", err)
			continue
		}

		fmt.Println("Message is: ", message.Subject)
	}
}

func SendMessage(connection *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		message := socket.Message{
			Subject: text,
		}

		err := websocket.JSON.Send(connection, message)

		if err != nil {
			log.Fatal("Error in sending message: ", err.Error())
			continue
		}
	}
}
