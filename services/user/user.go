package user

import (
	"fmt"
	"todo-cli-refactor/models"
)

type ServiceRepository interface {
	CreateNewUser(user models.User) (models.User, error)
	ListUsers() ([]models.User, error)
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
	Name     string
	Email    string
	Password string
}

type CreateResponse struct {
	User models.User
}

func (u Service) Create(req CreateRequest) (CreateResponse, error) {

	createdUser, cErr := u.repository.CreateNewUser(models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if cErr != nil {
		return CreateResponse{}, fmt.Errorf("can't create new User: %v", cErr)
	}

	return CreateResponse{User: createdUser}, nil
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	User models.User
}

func (u Service) Login(req LoginRequest) (LoginResponse, error) {

	users, err := u.repository.ListUsers()
	if err != nil {
		return LoginResponse{}, fmt.Errorf("can't list users: %v", err)
	}

	// Loop over the users and check if any of them matches the email and password
	var authenticatedUser *models.User
	for _, user := range users {
		//TODO : change req.password after implement hash
		if user.Email == req.Email && user.Password == req.Password {
			authenticatedUser = &user
			break
		}
	}

	// If no user matches, return an error
	if authenticatedUser == nil {
		return LoginResponse{}, fmt.Errorf("the email or password is not correct")
	}

	// Return the authenticated user in the response
	return LoginResponse{User: *authenticatedUser}, nil
}

type ListUsersRequest struct{}

type ListUsersResponse struct {
	Users []models.User
}

func (u Service) ListUsers(req ListUsersRequest) (ListUsersResponse, error) {

	users, err := u.repository.ListUsers()
	if err != nil {
		return ListUsersResponse{}, fmt.Errorf("can't list users: %v", err)
	}

	return ListUsersResponse{Users: users}, nil
}
