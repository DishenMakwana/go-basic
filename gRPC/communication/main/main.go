package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dishenmakwana/gRPC/communication/protocol"
	"github.com/dishenmakwana/gRPC/communication/server"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type Configuration struct {
	DriverName     string `json:"driverName"`
	DataSourceName string `json:"dataSourceName"`
}

var conf Configuration = Configuration{}

func main() {

	file, err := os.Open("configuration/config.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	json.NewDecoder(file).Decode(&conf)

	option := flag.String("admin", "server", "Communication btw client and server")
	flag.Parse()

	switch *option {
	case "client":
		runGrpcClient()
	case "server":
		runGrpcServer()
	}

}

func runGrpcClient() {
	connection, err := grpc.Dial("127.0.0.1:8085", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err.Error())
	}

	defer connection.Close()

	client := protocol.NewStudentServiceClient(connection)

	fmt.Println("\n\nWelcome, Please enter number to select your requests")

	fmt.Println("1:Select All Students Based Name")

	fmt.Println("2:Select a Student Based ID Number")

	fmt.Print("\nPlease enter your Number:")

	input := ""
	fmt.Scanln(&input)

	if strings.EqualFold(input, "1") {
		value := ""

		fmt.Print("Enter your Name:")

		fmt.Scanln(&value)
		fmt.Println("")

		students, err := client.GetStudentsByName(context.Background(), &protocol.SearchByName{Name: value})

		if err != nil {
			log.Fatal(err.Error())
		}

		for {

			student, err := students.Recv()

			if err == io.EOF {

				break
			}

			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(student)

		}

		fmt.Println("")
		return

	}

	fmt.Print("Enter your ID:")

	fmt.Scanln(&input)

	fmt.Println("")

	id, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err.Error())
	}

	student, err := client.GetStudentByID(context.Background(), &protocol.SearchByID{Id: int64(id)})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(student)

	fmt.Println("")
}

func runGrpcServer() {
	fmt.Println("Starting Server ..................!")

	listener, err := net.Listen("tcp", "127.0.0.1:8085")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Server is Listening .............!")

	var options []grpc.ServerOption

	newServer := grpc.NewServer(options...)

	studentServer, err := server.GrpcServerInitializer(conf.DriverName, conf.DataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	protocol.RegisterStudentServiceServer(newServer, studentServer)

	err = newServer.Serve(listener)

	if err != nil {
		log.Fatal(err.Error())
	}
}
