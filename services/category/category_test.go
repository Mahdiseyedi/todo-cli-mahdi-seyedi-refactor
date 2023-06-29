package category

import (
	"reflect"
	"testing"
	"todo-cli-refactor/models"
)

type mockRepository struct {
	data map[int]models.Category // map of category ID to category struct
}

func (m mockRepository) CreateNewCategory(c models.Category) (models.Category, error) {
	// Generate a unique ID for the category by incrementing the length of the map
	c.ID = len(m.data) + 1

	// Add the category to the map with the ID as the key
	m.data[c.ID] = c

	// Return the created category
	return c, nil
}

func TestCreate(t *testing.T) {
	// Create a mockRepository instance with some initial data
	mr := mockRepository{
		data: map[int]models.Category{
			1: {ID: 1, Title: "Work", Color: "red", UserID: 3},
			2: {ID: 2, Title: "Home", Color: "blue", UserID: 4},
			3: {ID: 3, Title: "Hobby", Color: "green", UserID: 5},
		},
	}

	// Create a Service instance with the mockRepository
	s := NewService(mr)

	// Create a CreateRequest instance with some sample data
	req := CreateRequest{
		Title:               "Travel",
		Color:               "yellow",
		AuthenticatedUserID: 6,
	}

	// Call the Create method and check for errors
	res, err := s.Create(req)
	if err != nil {
		t.Errorf("Create failed : %v", err)
	}

	// Check if the response contains the expected data
	expected := models.Category{
		ID:     4,
		Title:  "Travel",
		Color:  "yellow",
		UserID: 6,
	}
	if !reflect.DeepEqual(res.Category, expected) {
		t.Errorf("response does not match expected data : got %v , want %v ", res.Category, expected)
	}
}
