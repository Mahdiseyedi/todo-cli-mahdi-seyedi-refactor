package contract

import "todo-cli-refactor/models"

type CategoryWriteStore interface {
	Save(u models.Category)
}

type CategoryReadStore interface {
	Load() []models.Category
}
