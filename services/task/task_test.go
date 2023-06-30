package task

import (
	"reflect"
	"testing"
	"todo-cli-refactor/models"
)

type mockRepository struct {
	data map[int]models.Task
}

func (m mockRepository) CreateNewTask(task models.Task) (models.Task, error) {
	task.ID = len(m.data) + 1

	m.data[task.ID] = task

	return task, nil
}

func (m mockRepository) ListUserTasks(userID int) ([]models.Task, error) {
	var tasks []models.Task

	for _, task := range m.data {
		if task.UserID == userID {
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func TestCreate(t *testing.T) {
	mr := mockRepository{
		data: map[int]models.Task{
			1: {ID: 1, Title: "Buy groceries", DueDate: "2021-12-31", CategoryID: 2, IsDone: false, UserID: 3},
			2: {ID: 2, Title: "Clean the house", DueDate: "2022-01-01", CategoryID: 1, IsDone: true, UserID: 4},
			3: {ID: 3, Title: "Read a book", DueDate: "2022-01-02", CategoryID: 3, IsDone: false, UserID: 5},
		},
	}

	s := NewService(mr)

	req := CreateRequest{
		Title:               "Watch a movie",
		DueDate:             "2022-01-03",
		CategoryID:          4,
		AuthenticatedUserID: 6,
	}

	res, err := s.Create(req)
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}

	expected := models.Task{
		ID:         4,
		Title:      "Watch a movie",
		DueDate:    "2022-01-03",
		CategoryID: 4,
		IsDone:     false,
		UserID:     6,
	}
	if !reflect.DeepEqual(res.Task, expected) {
		t.Errorf("response does not match expected data: got %v, want %v", res.Task, expected)
	}
}

func TestListUserTasks(t *testing.T) {
	mr := mockRepository{
		data: map[int]models.Task{
			1: {ID: 1, Title: "Buy groceries", DueDate: "2021-12-31", CategoryID: 2, IsDone: false, UserID: 3},
			2: {ID: 2, Title: "Clean the house", DueDate: "2022-01-01", CategoryID: 1, IsDone: true, UserID: 4},
			3: {ID: 3, Title: "Read a book", DueDate: "2022-01-02", CategoryID: 3, IsDone: false, UserID: 5},
			4: {ID: 4, Title: "Watch a movie", DueDate: "2022-01-03", CategoryID: 4, IsDone: false, UserID: 6},
		},
	}

	s := NewService(mr)

	req := ListRequest{
		UserID: 3,
	}

	res, err := s.List(req)
	if err != nil {
		t.Errorf("List failed: %v", err)
	}

	expected := []models.Task{
		{ID: 1, Title: "Buy groceries", DueDate: "2021-12-31", CategoryID: 2, IsDone: false, UserID: 3},
	}
	if !reflect.DeepEqual(res.Tasks, expected) {
		t.Errorf("response does not match expected data: got %v, want %v", res.Tasks, expected)
	}
}
