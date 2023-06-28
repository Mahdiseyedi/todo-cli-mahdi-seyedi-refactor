package task

import (
	"fmt"
	"todo-cli-refactor/models"
)

type ServiceRepository interface {
	CreateNewTask(t models.Task) (models.Task, error)
	ListUserTasks(userID int) ([]models.Task, error)
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
	DueDate             string
	CategoryID          int
	AuthenticatedUserID int
}

type CreateResponse struct {
	Task models.Task
}

func (t Service) Create(req CreateRequest) (CreateResponse, error) {

	createdTask, cErr := t.repository.CreateNewTask(models.Task{
		Title:      req.Title,
		DueDate:    req.DueDate,
		CategoryID: req.CategoryID,
		IsDone:     false,
		UserID:     req.AuthenticatedUserID,
	})
	if cErr != nil {
		return CreateResponse{}, fmt.Errorf("can't create new task: %v", cErr)
	}

	return CreateResponse{Task: createdTask}, nil
}

type ListRequest struct {
	UserID int
}

type ListResponse struct {
	Tasks []models.Task
}

func (t Service) List(req ListRequest) (ListResponse, error) {
	tasks, err := t.repository.ListUserTasks(req.UserID)
	if err != nil {
		return ListResponse{}, fmt.Errorf("can't list user tasks: %v", err)
	}

	return ListResponse{Tasks: tasks}, nil
}
