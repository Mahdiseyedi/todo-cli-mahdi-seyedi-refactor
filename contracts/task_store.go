package contract

import "todo-cli-refactor/models"

type TaskWriteStore interface {
	Save(u models.Task)
}

type TaskReadStore interface {
	Load() []models.Task
}
