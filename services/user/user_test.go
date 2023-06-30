package user

import (
	"reflect"
	"testing"
	"todo-cli-refactor/models"
)

type mockRepository struct {
	data map[int]models.User
}

func (m mockRepository) CreateNewUser(user models.User) (models.User, error) {
	user.ID = len(m.data) + 1

	m.data[user.ID] = user

	return user, nil
}

func (m mockRepository) ListUsers() ([]models.User, error) {
	var users []models.User

	for _, user := range m.data {
		users = append(users, user)
	}

	return users, nil
}

func TestCreate(t *testing.T) {
	mr := mockRepository{
		data: map[int]models.User{
			1: {ID: 1, Name: "Alice", Email: "alice@example.com", Password: "123456"},
			2: {ID: 2, Name: "Bob", Email: "bob@example.com", Password: "654321"},
			3: {ID: 3, Name: "Charlie", Email: "charlie@example.com", Password: "abcdef"},
		},
	}

	s := NewService(mr)

	req := CreateRequest{
		Name:     "David",
		Email:    "david@example.com",
		Password: "123456",
	}

	res, err := s.Create(req)
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}

	expected := models.User{
		ID:       4,
		Name:     "David",
		Email:    "david@example.com",
		Password: "123456", // TODO: change this to the hashed password when implemented
	}
	if !reflect.DeepEqual(res.User, expected) {
		t.Errorf("response does not match expected data: got %v, want %v", res.User, expected)
	}
}

func TestLogin(t *testing.T) {
	mr := mockRepository{
		data: map[int]models.User{
			1: {ID: 1, Name: "Alice", Email: "alice@example.com", Password: "123456"},
			2: {ID: 2, Name: "Bob", Email: "bob@example.com", Password: "654321"},
			3: {ID: 3, Name: "Charlie", Email: "charlie@example.com", Password: "abcdef"},
			4: {ID: 4, Name: "David", Email: "david@example.com", Password: "123456"},
		},
	}

	s := NewService(mr)

	t.Run("valid email and password", func(t *testing.T) {

		req := LoginRequest{
			Email:    "david@example.com",
			Password: "123456",
		}

		res, err := s.Login(req)
		if err != nil {
			t.Errorf("Login failed : %v", err)
		}

		expected := models.User{
			ID:       4,
			Name:     "David",
			Email:    "david@example.com",
			Password: "123456", // TODO : change this to hashed password when implemented
		}
		if !reflect.DeepEqual(res.User, expected) {
			t.Errorf("response does not match expected data : got %v , want %v ", res.User, expected)
		}
	})

	t.Run("invalid email or password", func(t *testing.T) {

		req := LoginRequest{
			Email:    "david@example.com",
			Password: "wrongpassword",
		}

		_, err := s.Login(req)
		if err == nil {
			t.Errorf("Login should fail with invalid email or password")
		}
	})
}

func TestListUsers(t *testing.T) {
	mr := mockRepository{
		data: map[int]models.User{
			1: {ID: 1, Name: "Alice", Email: "alice@example.com", Password: "123456"},
			2: {ID: 2, Name: "Bob", Email: "bob@example.com", Password: "654321"},
			3: {ID: 3, Name: "Charlie", Email: "charlie@example.com", Password: "abcdef"},
			4: {ID: 4, Name: "David", Email: "david@example.com", Password: "123456"},
		},
	}

	s := NewService(mr)

	req := ListUsersRequest{}

	res, err := s.ListUsers(req)
	if err != nil {
		t.Errorf("ListUsers failed : %v", err)
	}

	expected := []models.User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Password: "123456"},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Password: "654321"},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com", Password: "abcdef"},
		{ID: 4, Name: "David", Email: "david@example.com", Password: "123456"},
	}
	if !reflect.DeepEqual(res.Users, expected) {
		t.Errorf("response does not match expected data : got %v , want %v ", res.Users, expected)
	}
}
