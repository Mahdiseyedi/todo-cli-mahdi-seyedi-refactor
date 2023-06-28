package category

import (
	"fmt"
	"todo-cli-refactor/models"
)

type ServiceRepository interface {
	CreateNewCategory(c models.Category) (models.Category, error)
}

type Service struct {
	repository ServiceRepository
}

func NewService(repo ServiceRepository) Service {
	return Service{
		repository: repo,
	}
}

type CreateRequest struct {
	Title               string
	Color               string
	AuthenticatedUserID int
}

type CreateResponse struct {
	Category models.Category
}

func (c Service) Create(req CreateRequest) (CreateResponse, error) {

	createdCategory, cErr := c.repository.CreateNewCategory(models.Category{
		Title:  req.Title,
		Color:  req.Color, // Added the color field to the category struct
		UserID: req.AuthenticatedUserID,
	})
	if cErr != nil {
		return CreateResponse{}, fmt.Errorf("can't create new category: %v", cErr)
	}

	return CreateResponse{Category: createdCategory}, nil
}
