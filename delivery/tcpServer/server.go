package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"todo-cli-refactor/consts"
	"todo-cli-refactor/delivery/deliveryParam"
	"todo-cli-refactor/repositories/fileRepository/task"
	task2 "todo-cli-refactor/services/task"
)

func main() {
	const (
		network = "tcp"
		address = ":9986"
	)

	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatalln("cant listen on given address", err)
	}

	defer listener.Close()

	fmt.Println("server listening on: ", listener.Addr())

	f := task.New("./task.txt", consts.JsonSerializationMode)
	taskService := task2.NewService(f)

	for {
		connection, aErr := listener.Accept()
		if aErr != nil {
			log.Println("cant listen to new connection, ", aErr)
			continue
		}

		var rawRequest = make([]byte, 1024)
		numberOfBytes, rErr := connection.Read(rawRequest)
		if rErr != nil {
			log.Println("cant read data from connection", rErr)

			continue
		}

		fmt.Printf("client address: %s, numberOfByttes: %d, data: %s\n",
			connection.RemoteAddr(), numberOfBytes, string(rawRequest))

		req := &deliveryParam.Request{}
		if uErr := json.Unmarshal(rawRequest[:numberOfBytes], req); uErr != nil {
			log.Println("bad request...", uErr)
			continue
		}

		switch req.Command {
		case "create-task":
			response, cErr := taskService.Create(task2.CreateRequest{
				Title:               req.CreateTaskRequest.Title,
				DueDate:             req.CreateTaskRequest.DueDate,
				CategoryID:          req.CreateTaskRequest.CategoryID,
				AuthenticatedUserID: 0,
			})
			if cErr != nil {
				_, wErr := connection.Write([]byte(cErr.Error()))
				if wErr != nil {
					log.Println("cant write data to connection,", rErr)
					continue
				}
			}

			data, mErr := json.Marshal(&response)
			if mErr != nil {
				_, wErr := connection.Write([]byte(mErr.Error()))
				if wErr != nil {
					log.Println("cant marshal response,", rErr)

					continue
				}

				continue
			}

			_, wErr := connection.Write(data)
			if wErr != nil {
				log.Println("cant write data to connection", rErr)
				continue
			}
		}

		connection.Close()
	}

}
