package user

import (
	"reflect"
	"testing"
	"todo-cli-refactor/models"
)

// mockRepository is a struct that implements the UserServiceRepository interface using a map as the data source
type mockRepository struct {
	data map[int]models.User // map of user ID to user struct
}

// CreateNewUser is a method that creates a new user and adds it to the map
func (m mockRepository) CreateNewUser(user models.User) (models.User, error) {
	// Generate a unique ID for the user by incrementing the length of the map
	user.ID = len(m.data) + 1

	// Add the user to the map with the ID as the key
	m.data[user.ID] = user

	// Return the created user
	return user, nil
}

// ListUsers is a method that returns a slice of all users in the map
func (m mockRepository) ListUsers() ([]models.User, error) {
	// Create an empty slice of users
	var users []models.User

	// Loop over the map and append the users to the slice
	for _, user := range m.data {
		users = append(users, user)
	}

	// Return the slice of users
	return users, nil
}

func TestCreate(t *testing.T) {
	// Create a mockRepository instance with some initial data
	mr := mockRepository{
		data: map[int]models.User{
			1: {ID: 1, Name: "Alice", Email: "alice@example.com", Password: "123456"},
			2: {ID: 2, Name: "Bob", Email: "bob@example.com", Password: "654321"},
			3: {ID: 3, Name: "Charlie", Email: "charlie@example.com", Password: "abcdef"},
		},
	}

	// Create a Service instance with the mockRepository
	s := NewService(mr)

	// Create a CreateRequest instance with some sample data
	req := CreateRequest{
		Name:     "David",
		Email:    "david@example.com",
		Password: "123456",
	}

	// Call the Create method and check for errors
	res, err := s.Create(req)
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}

	// Check if the response contains the expected data
	expected := models.User{
		ID:       4, // The next ID should be 4
		Name:     "David",
		Email:    "david@example.com",
		Password: "123456", // TODO: change this to the hashed password when implemented
	}
	if !reflect.DeepEqual(res.User, expected) {
		t.Errorf("response does not match expected data: got %v, want %v", res.User, expected)
	}
}

func TestLogin(t *testing.T) {
	// Create a mockRepository instance with some initial data
	mr := mockRepository{
		data: map[int]models.User{
			1: {ID: 1, Name: "Alice", Email: "alice@example.com", Password: "123456"},
			2: {ID: 2, Name: "Bob", Email: "bob@example.com", Password: "654321"},
			3: {ID: 3, Name: "Charlie", Email: "charlie@example.com", Password: "abcdef"},
			4: {ID: 4, Name: "David", Email: "david@example.com", Password: "123456"},
		},
	}

	// Create a Service instance with the mockRepository
	s := NewService(mr)

	t.Run("valid email and password", func(t *testing.T) {

		// Create a LoginRequest instance with valid email and password
		req := LoginRequest{
			Email:    "david@example.com",
			Password: "123456",
		}

		// Call the Login method and check for errors
		res, err := s.Login(req)
		if err != nil {
			t.Errorf("Login failed : %v", err)
		}

		// Check if the response contains the expected data
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

		// Create a LoginRequest instance with invalid email or password
		req := LoginRequest{
			Email:    "david@example.com",
			Password: "wrongpassword",
		}

		// Call the Login method and expect an error
		_, err := s.Login(req)
		if err == nil {
			t.Errorf("Login should fail with invalid email or password")
		}
	})
}
