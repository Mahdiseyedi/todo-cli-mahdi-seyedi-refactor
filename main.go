package main

import (
	"fmt"
	"todo-cli-refactor/models"
	"todo-cli-refactor/repositories/fileRepository/user"
)

func main() {

	//f := user.New("./user.txt", consts.TextSerializationMode)
	//data := []string{
	//	"ID: 1,  Name: hooooo,  Email: eeeee@eeee,  Password: 423234 ",
	//	"ID: 60,  Name: hosdf,  Email: h@asdgsd,  Password: 2 ",
	//	"ID: 3,  Name: 234,  Email: hsdg@sd,  Password: 653ac11ca60b3e021a8c609c7198acfc ",
	//	"ID: 4,  Name: Hossein,  Email: h@h.com,  Password: 3f62d610adac92bbf2996bb4d8ff7657 ",
	//	"ID: 50,  Name: hh,  Email: hhhhh,  Password: 94b40c6db280230b4211b06fa04c7be1 ",
	//	"ID: 6,  Name: hhhhhh,  Email: h234,  Password: 789406d01073ca1782d86293dcfc0764 ",
	//	"ID: 7,  Name: h3,  Email: h3,  Password: 6f207f8b5dfe1eebac63467930df5189 ",
	//	"ID: 8,  Name: hhhh,  Email: h@hhhh.com,  Password: c4ca4238a0b923820dcc509a6f75849b ",
	//}
	t2Data := "ID: 300,  Name: 234, Email: hsdg@sd, Password: 653ac11ca60b3e021a8c609c7198acfc"
	user1 := models.User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "123456",
	}

	// Format the user data as a text string
	tData := fmt.Sprintf("ID: %d, Name: %s, Email: %s, Password: %s\n", user1.ID, user1.Name,
		user1.Email, user1.Password)
	fmt.Println(tData)
	k, _ := user.TextDeserializer(tData)
	//k := f.UserDeserializer(data)

	//for _, o := range k {
	//	fmt.Println(o)
	//}
	fmt.Println(k.ID)
	fmt.Println(k.Name)
	fmt.Println(k.Email)
	fmt.Println(k.Password)
	fmt.Println("-----------------------------------")
	k2, _ := user.TextDeserializer(t2Data)
	fmt.Println(t2Data)
	fmt.Println(k2.ID)
	fmt.Println(k2.Name)
	fmt.Println(k2.Email)
	fmt.Println(k2.Password)
}
