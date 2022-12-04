package main

import (
	"github.com/dishenmakwana/websocket/socket"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func main() {

	config := socket.NewConfig()

	muxServer := http.NewServeMux()

	muxServer.Handle("/", websocket.Handler(func(connection *websocket.Conn) {
		webSocketHandler(connection, config)
	}))

	server := http.Server{
		Addr:    "127.0.0.1:8085",
		Handler: muxServer}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func webSocketHandler(connection *websocket.Conn, config *socket.Config) {

	go config.RunSocket()

	config.RegisterClient <- connection

	for {

		var message socket.Message
		err := websocket.JSON.Receive(connection, &message)

		if err != nil {
			config.RemoveClient <- connection
			continue
		}

		config.MessageData <- message
	}
}
