package socket

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

type Message struct {
	Subject string `json:"subject"`
}

type Config struct {
	Clients map[string]*websocket.Conn

	RegisterClient chan *websocket.Conn

	RemoveClient chan *websocket.Conn

	MessageData chan Message
}

func NewConfig() *Config {
	return &Config{
		Clients:        make(map[string]*websocket.Conn),
		RegisterClient: make(chan *websocket.Conn),
		RemoveClient:   make(chan *websocket.Conn),
		MessageData:    make(chan Message),
	}
}

func (config Config) RegisterClients(client *websocket.Conn) {
	config.Clients[client.RemoteAddr().String()] = client

	fmt.Println("Added Clients: ", config.Clients)
}

func (config Config) RemoveClients(client *websocket.Conn) {
	delete(config.Clients, client.RemoteAddr().String())

	fmt.Println("Removed Clients: ", config.Clients)
}

func (config Config) MessageDatas(message Message) {
	for _, client := range config.Clients {
		err := websocket.JSON.Send(client, message)

		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func (config Config) RunSocket() {
	for {
		select {
		case registerClient := <-config.RegisterClient:
			config.RegisterClients(registerClient)
		case removeClient := <-config.RemoveClient:
			config.RemoveClients(removeClient)
		case messageData := <-config.MessageData:
			config.MessageDatas(messageData)
		}
	}
}
