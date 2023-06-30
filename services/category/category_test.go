package category

import (
	"reflect"
	"testing"
	"todo-cli-refactor/models"
)

type mockRepository struct {
	data map[int]models.Category
}

func (m mockRepository) CreateNewCategory(c models.Category) (models.Category, error) {
	c.ID = len(m.data) + 1

	m.data[c.ID] = c

	return c, nil
}

func TestCreate(t *testing.T) {
	mr := mockRepository{
		data: map[int]models.Category{
			1: {ID: 1, Title: "Work", Color: "red", UserID: 3},
			2: {ID: 2, Title: "Home", Color: "blue", UserID: 4},
			3: {ID: 3, Title: "Hobby", Color: "green", UserID: 5},
		},
	}

	s := NewService(mr)

	req := CreateRequest{
		Title:               "Travel",
		Color:               "yellow",
		AuthenticatedUserID: 6,
	}

	res, err := s.Create(req)
	if err != nil {
		t.Errorf("Create failed : %v", err)
	}

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
