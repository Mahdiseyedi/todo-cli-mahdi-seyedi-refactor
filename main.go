package main

import (
	"fmt"
	"todo-cli-refactor/models"
	"todo-cli-refactor/repositories/fileRepository/task"
)

func main() {
	task1 := models.Task{
		ID:         1,
		Title:      "Buy groceries",
		DueDate:    "2021-12-31",
		CategoryID: 2,
		IsDone:     false,
		UserID:     3,
	}

	data := fmt.Sprintf("id: %d, title: %s, dueDate: %s, categoryID: %d, isDone: %t, userID: %d\n", task1.ID, task1.Title,
		task1.DueDate, task1.CategoryID, task1.IsDone, task1.UserID)
	v, e := task.TextDeserializer(data)

	fmt.Println(v, e)
}
