package main

import (
	"fmt"
	"todo-cli-refactor/repositories/fileRepository"
)

func main() {
	res, _ := fileRepository.TextDeserializer("id: 10, name: h@h, email: 1, password: c4ca4238a0b923820dcc509a6f75849b")
	fmt.Println("*", res.ID, "*")
	fmt.Println("*", res.Name, "*")
	fmt.Println("*", res.Email, "*")
	fmt.Println("*", res.Password, "*")
}
