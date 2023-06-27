package fileRepository

// Import the testing and assert packages
// Import the testing package
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

//Define a test function with the prefix Test
func TestWriteUserToFile(t *testing.T) {
	// Create a FileStore instance with a test file path and serialization mode
	f := FileStore{
		Filepath:          "test.txt",
		serializationMode: consts.JsonSerializationMode,
	}

	// Create a user instance with some sample data
	user := models.User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "123456",
	}

	// Call the writeUserToFile method and check for errors
	err := f.writeUserToFile(user)
	if err != nil {
		// Use the t.Errorf method to report an error
		t.Errorf("writeUserToFile failed: %v", err)
	}

	// Read the test file and check if it contains the expected data
	data, err := ioutil.ReadFile(f.Filepath)
	if err != nil {
		t.Errorf("can't read test file: %v", err)
	}
	expected := "{\"ID\":1,\"Name\":\"Alice\",\"Email\":\"alice@example.com\",\"Password\":\"123456\"}\n"
	if string(data) != expected {
		t.Errorf("test file does not match expected data: got %s, want %s", data, expected)
	}

	// Delete the test file after the test is done
	err = os.Remove(f.Filepath)
	if err != nil {
		t.Errorf("can't delete test file: %v", err)
	}
}

func TestJsonDeserializer(t *testing.T) {
	// Create a user instance with some sample data
	user := models.User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "123456",
	}

	// Marshal the user data to JSON and check for errors
	data, err := json.Marshal(user)
	if err != nil {
		t.Errorf("can't marshal user struct to json: %v", err)
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
	user, err = JsonDeserializer(string(data))
	if err != nil {
		t.Errorf("JsonDeserializer failed: %v", err)
	}

	// Check if the user struct matches the expected data
	expected := models.User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "123456",
	}
	if !reflect.DeepEqual(user, expected) {
		t.Errorf("user does not match expected data: got %v, want %v", user, expected)
	}

	// Delete the test file after the test is done
	err = os.Remove("test.txt")
	if err != nil {
		t.Errorf("can't delete test file: %v", err)
	}
}

func TestTextDeserializer(t *testing.T) {
	// Create a user instance with some sample data
	user := models.User{
		ID:       10,
		Name:     "h@h",
		Email:    "1",
		Password: "c4ca4238a0b923820dcc509a6f75849b",
	}

	// Format the user data as a text string
	data := fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n", user.ID, user.Name,
		user.Email, user.Password)

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
	user, err = TextDeserializer(string(data2))
	if err != nil {
		t.Errorf("TextDeserializer failed: %v", err)
	}

	// Check if the user struct matches the expected data
	expected := models.User{
		ID:       10,
		Name:     "h@h",
		Email:    "1",
		Password: "c4ca4238a0b923820dcc509a6f75849b",
	}
	if !reflect.DeepEqual(user, expected) {
		t.Errorf("user does not match expected data: got %v, want %v", user, expected)
	}

	// Delete the test file after the test is done
	err = os.Remove("test.txt")
	if err != nil {
		t.Errorf("can't delete test file: %v", err)
	}
}
