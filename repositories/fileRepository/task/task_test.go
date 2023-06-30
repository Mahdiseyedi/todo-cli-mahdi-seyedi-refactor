package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"todo-cli-refactor/consts"
	"todo-cli-refactor/models"
)

func TestWriteTaskToFile(t *testing.T) {
	f := FileStore{
		Filepath:          "test.txt",
		serializationMode: consts.JsonSerializationMode,
	}

	task := models.Task{
		ID:         1,
		Title:      "Buy groceries",
		DueDate:    "2021-12-31",
		CategoryID: 2,
		IsDone:     false,
		UserID:     3,
	}

	err := f.writeTaskToFile(task)
	if err != nil {
		t.Errorf("writeTaskToFile failed: %v", err)
	}

	data, err := ioutil.ReadFile(f.Filepath)
	if err != nil {
		t.Errorf("can't read test file: %v", err)
	}
	expected := "{\"ID\":1,\"Title\":\"Buy groceries\",\"DueDate\":\"2021-12-31\",\"CategoryID\":2,\"IsDone\":false,\"UserID\":3}\n"
	if string(data) != expected {
		t.Errorf("test file does not match expected data: got %s, want %s", data, expected)
	}

	err = os.Remove(f.Filepath)
	if err != nil {
		t.Errorf("can't delete test file: %v", err)
	}
}
func TestTaskJsonDeserializer(t *testing.T) {
	task := models.Task{
		ID:         1,
		Title:      "Buy groceries",
		DueDate:    "2021-12-31",
		CategoryID: 2,
		IsDone:     false,
		UserID:     3,
	}

	data, err := json.Marshal(task)
	if err != nil {
		t.Errorf("can't marshal task struct to json: %v", err)
	}

	err = ioutil.WriteFile("test.txt", data, 0644)
	if err != nil {
		t.Errorf("can't write to test file: %v", err)
	}

	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		t.Errorf("can't read test file: %v", err)
	}

	task, err = JsonDeserializer(string(data))
	if err != nil {
		t.Errorf("JsonDeserializer failed: %v", err)
	}

	expected := models.Task{
		ID:         1,
		Title:      "Buy groceries",
		DueDate:    "2021-12-31",
		CategoryID: 2,
		IsDone:     false,
		UserID:     3,
	}
	if !reflect.DeepEqual(task, expected) {
		t.Errorf("task does not match expected data: got %v, want %v", task, expected)
	}

	err = os.Remove("test.txt")
	if err != nil {
		t.Errorf("can't delete test file: %v", err)
	}
}
func TestTaskTextDeserializer(t *testing.T) {
	task := models.Task{
		ID:         1,
		Title:      "Buy groceries",
		DueDate:    "2021-12-31",
		CategoryID: 2,
		IsDone:     false,
		UserID:     3,
	}

	data := fmt.Sprintf("id: %d, title: %s, dueDate: %s, categoryID: %d, isDone: %t, userID: %d\n", task.ID, task.Title,
		task.DueDate, task.CategoryID, task.IsDone, task.UserID)

	err := ioutil.WriteFile("test.txt", []byte(data), 0644)
	if err != nil {
		t.Errorf("can't write to test file: %v", err)
	}

	data2, err1 := ioutil.ReadFile("test.txt")
	if err1 != nil {
		t.Errorf("can't read test file: %v", err1)
	}

	task, err = TextDeserializer(string(data2))
	if err != nil {
		t.Errorf("TextDeserializer failed: %v", err)
	}

	expected := models.Task{
		ID:         1,
		Title:      "Buy groceries",
		DueDate:    "2021-12-31",
		CategoryID: 2,
		IsDone:     false,
		UserID:     3,
	}
	if !reflect.DeepEqual(task, expected) {
		t.Errorf("task does not match expected data: got %v, want %v", task, expected)
	}

	err = os.Remove("test.txt")
	if err != nil {
		t.Errorf("can't delete test file: %v", err)
	}
}
func TestTaskDeserializer(t *testing.T) {

	fs := FileStore{serializationMode: consts.TextSerializationMode}
	pData := []string{
		"id: 1, title: Buy groceries, dueDate: 2021-12-31, categoryID: 2, isDone: false, userID: 3",
		"id: 2, title: Clean the house, dueDate: 2022-01-01, categoryID: 1, isDone: true, userID: 4",
		"id: 3, title: Read a book, dueDate: 2022-01-02, categoryID: 3, isDone: false, userID: 5",
		"id: 4, title: Watch a movie, dueDate: 2022-01-03, categoryID: 4, isDone: true, userID: 6",
		"id: invalid, title: Do nothing, dueDate: 2022-01-04, categoryID: 5, isDone: false, userID: 7",
	}

	tasks := fs.TaskDeserializer(pData)
	if len(tasks) != 4 {
		t.Errorf("expected 4 tasks, got %d", len(tasks))
	}

	expectedTasks := []models.Task{
		{ID: 1, Title: "Buy groceries", DueDate: "2021-12-31", CategoryID: 2, IsDone: false, UserID: 3},
		{ID: 2, Title: "Clean the house", DueDate: "2022-01-01", CategoryID: 1, IsDone: true, UserID: 4},
		{ID: 3, Title: "Read a book", DueDate: "2022-01-02", CategoryID: 3, IsDone: false, UserID: 5},
		{ID: 4, Title: "Watch a movie", DueDate: "2022-01-03", CategoryID: 4, IsDone: true, UserID: 6},
	}

	for i, task := range tasks {
		if task != expectedTasks[i] {
			t.Errorf("expected task %v, got %v", expectedTasks[i], task)
		}
	}
}
func TestSave(t *testing.T) {

	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	fs := FileStore{Filepath: tmpfile.Name(), serializationMode: consts.TextSerializationMode}
	task := models.Task{ID: 1, Title: "Buy groceries", DueDate: "2021-12-31", CategoryID: 2, IsDone: false, UserID: 3}

	fs.Save(task)

	tmpfile.Close()

	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	expectedLine := "id: 1, title: Buy groceries, dueDate: 2021-12-31, categoryID: 2, isDone: false, userID: 3"
	if line != expectedLine {
		t.Errorf("expected line %s, got %s", expectedLine, line)
	}
}
func TestCreateNewTask(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	fs := FileStore{Filepath: tmpfile.Name(), serializationMode: consts.TextSerializationMode}

	task := models.Task{
		Title:      "Buy groceries",
		DueDate:    "2021-12-31",
		CategoryID: 2,
		IsDone:     false,
		UserID:     3,
	}

	createdTask, err := fs.CreateNewTask(task)
	if err != nil {
		t.Errorf("CreateNewTask failed: %v", err)
	}

	if createdTask.ID == 0 {
		t.Errorf("expected a non-zero ID, got %d", createdTask.ID)
	}

	expectedTask := task
	expectedTask.ID = createdTask.ID
	if !reflect.DeepEqual(createdTask, expectedTask) {
		t.Errorf("created task does not match expected data: got %v, want %v", createdTask, expectedTask)
	}

	data, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Errorf("can't read temporary file: %v", err)
	}
	expectedData := fmt.Sprintf("id: %d, title: Buy groceries, dueDate: 2021-12-31, categoryID: 2, isDone: false, userID: 3\n", createdTask.ID)
	if string(data) != expectedData {
		t.Errorf("temporary file does not match expected data: got %s, want %s", data, expectedData)
	}
}
func TestGenerateID(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	fs := FileStore{Filepath: tmpfile.Name(), serializationMode: consts.TextSerializationMode}

	tasks := []models.Task{
		{ID: 1, Title: "Buy groceries", DueDate: "2021-12-31", CategoryID: 2, IsDone: false, UserID: 3},
		{ID: 2, Title: "Clean the house", DueDate: "2022-01-01", CategoryID: 1, IsDone: true, UserID: 4},
		{ID: 3, Title: "Read a book", DueDate: "2022-01-02", CategoryID: 3, IsDone: false, UserID: 5},
	}
	for _, task := range tasks {
		err := fs.writeTaskToFile(task)
		if err != nil {
			t.Errorf("can't write task to file: %v", err)
		}
	}

	id, err := fs.generateID()
	if err != nil {
		t.Errorf("generateID failed: %v", err)
	}

	expectedID := tasks[len(tasks)-1].ID + 1
	if id != expectedID {
		t.Errorf("expected ID %d, got %d", expectedID, id)
	}
}
func TestListUserTasks(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	fs := FileStore{Filepath: tmpfile.Name(), serializationMode: consts.TextSerializationMode}

	tasks := []models.Task{
		{ID: 1, Title: "Buy groceries", DueDate: "2021-12-31", CategoryID: 2, IsDone: false, UserID: 3},
		{ID: 2, Title: "Clean the house", DueDate: "2022-01-01", CategoryID: 1, IsDone: true, UserID: 4},
		{ID: 3, Title: "Read a book", DueDate: "2022-01-02", CategoryID: 3, IsDone: false, UserID: 5},
	}
	for _, task := range tasks {
		err := fs.writeTaskToFile(task)
		if err != nil {
			t.Errorf("can't write task to file: %v", err)
		}
	}

	userID := 3
	result, err := fs.ListUserTasks(userID)
	if err != nil {
		t.Errorf("ListUserTasks failed: %v", err)
	}

	expected := []models.Task{
		{ID: 1, Title: "Buy groceries", DueDate: "2021-12-31", CategoryID: 2, IsDone: false, UserID: 3},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result does not match expected data: got %v, want %v", result, expected)
	}
}
