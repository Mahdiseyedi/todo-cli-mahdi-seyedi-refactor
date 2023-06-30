package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"todo-cli-refactor/delivery/deliveryParam"
)

func main() {
	fmt.Println("command", os.Args[0])

	if len(os.Args) < 2 {
		log.Fatalln("you input your server ip address")
	}

	serverAddress := os.Args[1]

	message := "default message"
	if len(os.Args) > 2 {
		message = os.Args[2]
	}

	connection, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalln("cant dial the server ...", err)
	}

	defer connection.Close()

	fmt.Println("local address", connection.LocalAddr())

	req := deliveryParam.Request{Command: message}
	if req.Command == "create-task" {
		req.CreateTaskRequest = deliveryParam.CreateTaskRequest{
			Title:      "test",
			DueDate:    "test",
			CategoryID: 1,
		}
	}

	serializedData, mErr := json.Marshal(&req)
	if mErr != nil {
		log.Fatalln("cant marshal request ", mErr)
	}

	numberOfWrittenBytes, wErr := connection.Write(serializedData)
	if wErr != nil {
		log.Fatalln("cant write to connection ", wErr)
	}

	fmt.Println("number of written bytes: ", numberOfWrittenBytes)

	var data = make([]byte, 1024)
	_, rErr := connection.Read(data)
	if rErr != nil {
		log.Fatalln("cant read data from connection: ", rErr)
	}

	fmt.Println("server response: ", string(data))
}
