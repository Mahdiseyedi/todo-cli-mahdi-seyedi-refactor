package main

import (
	"fmt"
	"todo-cli-refactor/repositories/fileRepository/category"
)

func main() {

	//id: 1, title: Work, color: Red, userID: 10
	res, err := category.TextDeserializer("id: 1, title: Work, color: Red, userID: 10")
	fmt.Println(res, err)
}
