package main

import (
	"fmt"
	"todo-cli-refactor/consts"
	"todo-cli-refactor/repositories/fileRepository/user"
)

func main() {

	f := user.New("./user.txt", consts.TextSerializationMode)
	data := []string{"ID: 1,  Name: hooooo,  Email: eeeee@eeee,  Password: 423234",
		"ID: 6,  Name: hosdf,  Email: h@asdgsd,  Password: 2"}
	k := f.UserDeserializer(data)
	fmt.Println("--------------------")
	for _, t := range k {
		fmt.Println(t)
	}
}
