package contract

import (
	"todo-cli-refactor/models"
)

type UserWriteStore interface {
	Save(u models.User)
}

type UserReadStore interface {
	Load() []models.User
}
