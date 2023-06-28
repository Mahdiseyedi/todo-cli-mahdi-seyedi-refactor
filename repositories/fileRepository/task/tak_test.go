package task

import (
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
	// Create a FileStore instance with a test file path and serialization mode
	f := FileStore{
		Filepath:          "test.txt",
		serializationMode: consts.JsonSerializationMode,
	}

	// Create a task instance with some sample data
	task := models.Task{
		ID:         1,
		Title:      "Buy groceries",
		DueDate:    "2021-12-31",
		CategoryID: 2,
		IsDone:     false,
		UserID:     3,
	}

	// Call the writeTaskToFile method and check for errors
	err := f.writeTaskToFile(task)
	if err != nil {
		// Use the t.Errorf method to report an error
		t.Errorf("writeTaskToFile failed: %v", err)
	}

	// Read the test file and check if it contains the expected data
	data, err := ioutil.ReadFile(f.Filepath)
	if err != nil {
		t.Errorf("can't read test file: %v", err)
	}
	expected := "{\"ID\":1,\"Title\":\"Buy groceries\",\"DueDate\":\"2021-12-31\",\"CategoryID\":2,\"IsDone\":false,\"UserID\":3}\n"
	if string(data) != expected {
		t.Errorf("test file does not match expected data: got %s, want %s", data, expected)
	}

	// Delete the test file after the test is done
	err = os.Remove(f.Filepath)
	if err != nil {
		t.Errorf("can't delete test file: %v", err)
	}
}

func TestTaskJsonDeserializer(t *testing.T) {
	// Create a task instance with some sample data
	task := models.Task{
		ID:         1,
		Title:      "Buy groceries",
		DueDate:    "2021-12-31",
		CategoryID: 2,
		IsDone:     false,
		UserID:     3,
	}

	// Marshal the task data to JSON and check for errors
	data, err := json.Marshal(task)
	if err != nil {
		t.Errorf("can't marshal task struct to json: %v", err)
	}

	// Write the JSON data to a test file and check for errors
	err = ioutil.WriteFile("test.txt", data, 0644)
	if err != nil {
		t.Errorf("can't write to test file: %v", err)
	}

	// Read the test file and check for errors
	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		t.Errorf("can't read test file: %v", err)
	}

	// Call the JsonDeserializer function and check for errors
	task, err = JsonDeserializer(string(data))
	if err != nil {
		t.Errorf("JsonDeserializer failed: %v", err)
	}

	// Check if the task struct matches the expected data
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

	// Delete the test file after the test is done
	err = os.Remove("test.txt")
	if err != nil {
		t.Errorf("can't delete test file: %v", err)
	}
}

func TestTaskTextDeserializer(t *testing.T) {
	// Create a task instance with some sample data
	task := models.Task{
		ID:         1,
		Title:      "Buy groceries",
		DueDate:    "2021-12-31",
		CategoryID: 2,
		IsDone:     false,
		UserID:     3,
	}

	// Format the task data as a text string
	data := fmt.Sprintf("id: %d, title: %s, dueDate: %s, categoryID: %d, isDone: %t, userID: %d\n", task.ID, task.Title,
		task.DueDate, task.CategoryID, task.IsDone, task.UserID)

	// Write the text data to a test file and check for errors
	err := ioutil.WriteFile("test.txt", []byte(data), 0644)
	if err != nil {
		t.Errorf("can't write to test file: %v", err)
	}

	// Read the test file and check for errors
	data2, err1 := ioutil.ReadFile("test.txt")
	if err1 != nil {
		t.Errorf("can't read test file: %v", err1)
	}

	// Call the TextDeserializer function and check for errors
	task, err = TextDeserializer(string(data2))
	if err != nil {
		t.Errorf("TextDeserializer failed: %v", err)
	}

	// Check if the task struct matches the expected data
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

	// Delete the test file after the test is done
	err = os.Remove("test.txt")
	if err != nil {
		t.Errorf("can't delete test file: %v", err)
	}
}
