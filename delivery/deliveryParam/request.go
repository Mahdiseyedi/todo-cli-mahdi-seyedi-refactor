package deliveryParam

import "todo-cli-refactor/services/task"

type Request struct {
	Command           string
	CreateTaskRequest task.CreateRequest
}

type Response struct {
	Title      string
	DueDate    string
	CategoryID int
}
