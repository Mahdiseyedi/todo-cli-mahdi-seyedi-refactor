package main

import (
	"todo-cli-refactor/consts"
	"todo-cli-refactor/models"
	"todo-cli-refactor/repositories/fileRepository/user"
)

func main() {
	f := user.New("./testik.txt", consts.TextSerializationMode)
	u := models.User{ID: 1, Name: "mahdi", Email: "mah@ter", Password: "wertwer"}
	f.Save(u)
}
