package task

import (
	"reflect"
	"testing"
	"todo-cli-refactor/models"
)

type mockRepository struct {
	data map[int]models.Task // map of task ID to task struct
}

func (m mockRepository) CreateNewTask(task models.Task) (models.Task, error) {
	// Generate a unique ID for the task by incrementing the length of the map
	task.ID = len(m.data) + 1

	// Add the task to the map with the ID as the key
	m.data[task.ID] = task

	// Return the created task
	return task, nil
}

func (m mockRepository) ListUserTasks(userID int) ([]models.Task, error) {
	// Create an empty slice of tasks
	var tasks []models.Task

	// Loop over the map and append the tasks that match the user ID to the slice
	for _, task := range m.data {
		if task.UserID == userID {
			tasks = append(tasks, task)
		}
	}

	// Return the slice of tasks
	return tasks, nil
}

func TestCreate(t *testing.T) {
	// Create a mockRepository instance with some initial data
	mr := mockRepository{
		data: map[int]models.Task{
			1: {ID: 1, Title: "Buy groceries", DueDate: "2021-12-31", CategoryID: 2, IsDone: false, UserID: 3},
			2: {ID: 2, Title: "Clean the house", DueDate: "2022-01-01", CategoryID: 1, IsDone: true, UserID: 4},
			3: {ID: 3, Title: "Read a book", DueDate: "2022-01-02", CategoryID: 3, IsDone: false, UserID: 5},
		},
	}

	// Create a Service instance with the mockRepository
	s := NewService(mr)

	// Create a CreateRequest instance with some sample data
	req := CreateRequest{
		Title:               "Watch a movie",
		DueDate:             "2022-01-03",
		CategoryID:          4,
		AuthenticatedUserID: 6,
	}

	// Call the Create method and check for errors
	res, err := s.Create(req)
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}

	// Check if the response contains the expected data
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
	// Create a mockRepository instance with some initial data
	mr := mockRepository{
		data: map[int]models.Task{
			1: {ID: 1, Title: "Buy groceries", DueDate: "2021-12-31", CategoryID: 2, IsDone: false, UserID: 3},
			2: {ID: 2, Title: "Clean the house", DueDate: "2022-01-01", CategoryID: 1, IsDone: true, UserID: 4},
			3: {ID: 3, Title: "Read a book", DueDate: "2022-01-02", CategoryID: 3, IsDone: false, UserID: 5},
			4: {ID: 4, Title: "Watch a movie", DueDate: "2022-01-03", CategoryID: 4, IsDone: false, UserID: 6},
		},
	}

	// Create a Service instance with the mockRepository
	s := NewService(mr)

	// Create a ListRequest instance with a sample user ID
	req := ListRequest{
		UserID: 3,
	}

	// Call the List method and check for errors
	res, err := s.List(req)
	if err != nil {
		t.Errorf("List failed: %v", err)
	}

	// Check if the response contains the expected data
	expected := []models.Task{
		{ID: 1, Title: "Buy groceries", DueDate: "2021-12-31", CategoryID: 2, IsDone: false, UserID: 3},
	}
	if !reflect.DeepEqual(res.Tasks, expected) {
		t.Errorf("response does not match expected data: got %v, want %v", res.Tasks, expected)
	}
}
