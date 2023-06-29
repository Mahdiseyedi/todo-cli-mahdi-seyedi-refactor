package user

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

func TestWriteUserToFile(t *testing.T) {
	// Create a FileStore instance with a test file path and serialization mode
	f := FileStore{
		Filepath:          "./test.txt",
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
func TestUserJsonDeserializer(t *testing.T) {
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
func TestUserTextDeserializer(t *testing.T) {
	// Create a user instance with some sample data
	user := models.User{
		ID:       10,
		Name:     "h@hmahdi",
		Email:    "1",
		Password: "c4ca4238a0b923820dcc509a6f75849b",
	}

	// Format the user data as a text string
	data := fmt.Sprintf("ID: %d, Name: %s, Email: %s, Password: %s\n", user.ID, user.Name,
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
		Name:     "h@hmahdi",
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
func TestUserDeserializer(t *testing.T) {

	fs := FileStore{serializationMode: consts.TextSerializationMode}
	pData := []string{
		"ID: 10, Name: h@h, Email: 1, Password: c4ca4238a0b923820dcc509a6f75849b",
		"ID: 20, Name: j@j, Email: 2, Password: c81e728d9d4c2f636f067f89cc14862c",
		"ID: 30, Name: k@k, Email: 3, Password: eccbc87e4b5ce2fe28308fd9f2a7baf3",
		"ID: 40, Name: l@l, Email: 4, Password: a87ff679a2f3e71d9181a67b7542122c",
		"ID: jk, Name: m@m, Email: 5, Password: e4da3b7fbbce2345d7772b0674a318d5",
	}

	users := fs.UserDeserializer(pData)
	if len(users) != 4 {
		t.Errorf("expected 4 users, got %d", len(users))
	}

	expectedUsers := []models.User{
		{ID: 10, Name: "h@h", Email: "1", Password: "c4ca4238a0b923820dcc509a6f75849b"},
		{ID: 20, Name: "j@j", Email: "2", Password: "c81e728d9d4c2f636f067f89cc14862c"},
		{ID: 30, Name: "k@k", Email: "3", Password: "eccbc87e4b5ce2fe28308fd9f2a7baf3"},
		{ID: 40, Name: "l@l", Email: "4", Password: "a87ff679a2f3e71d9181a67b7542122c"},
	}

	for i, user := range users {
		if user != expectedUsers[i] {
			t.Errorf("expected user %v, got %v", expectedUsers[i], user)
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
	user := models.User{ID: 50, Name: "n@n", Email: "6", Password: "1679091c5a880faf6fb5e6087eb1b2dc"}

	fs.Save(user)

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

	expectedLine := "ID: 50, Name: n@n, Email: 6, Password: 1679091c5a880faf6fb5e6087eb1b2dc"
	if line != expectedLine {
		t.Errorf("expected line %s, got %s", expectedLine, line)
	}
}
func TestCreateNewUser(t *testing.T) {
	// Create a temporary file and a FileStore instance with the file path and serialization mode
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	fs := FileStore{Filepath: tmpfile.Name(), serializationMode: consts.TextSerializationMode}

	// Create a user instance with some sample data
	user := models.User{
		Name:     "David",
		Email:    "David@example.com",
		Password: "123456",
	}

	// Call the CreateNewUser method and check for errors
	result, err := fs.CreateNewUser(user)
	if err != nil {
		t.Errorf("CreateNewUser failed: %v", err)
	}

	// Check if the result contains the expected data
	expected := models.User{
		ID:       1, // The first ID should be 1
		Name:     "David",
		Email:    "David@example.com",
		Password: "123456", // TODO: change this to the hashed Password when implemented
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result does not match expected data: got %v, want %v", result, expected)
	}
}
func TestGenerateID(t *testing.T) {
	// Create a temporary file and a FileStore instance with the file path and serialization mode
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	fs := FileStore{Filepath: tmpfile.Name(), serializationMode: consts.TextSerializationMode}

	// Write some sample data to the file using the writeUserToFile method
	users := []models.User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Password: "123456"},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Password: "654321"},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com", Password: "abcdef"},
	}
	for _, user := range users {
		err := fs.writeUserToFile(user)
		if err != nil {
			t.Errorf("can't write user to file: %v", err)
		}
	}

	// Call the generateID method and check for errors
	ID, err := fs.generateID()
	if err != nil {
		t.Errorf("generateID failed: %v", err)
	}

	// Check if the generated ID is one more than the last user ID
	expectedID := users[len(users)-1].ID + 1
	if ID != expectedID {
		t.Errorf("expected ID %d, got %d", expectedID, ID)
	}
}
func TestListUsers(t *testing.T) {
	// Create a temporary file and a FileStore instance with the file path and serialization mode
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	fs := FileStore{Filepath: tmpfile.Name(), serializationMode: consts.TextSerializationMode}

	// Write some sample users to the file
	users := []models.User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Password: "123456"},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Password: "654321"},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com", Password: "abcdef"},
	}
	for _, user := range users {
		err = fs.writeUserToFile(user)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Call the ListUsers method and check for errors
	result, err := fs.ListUsers()
	if err != nil {
		t.Errorf("ListUsers failed: %v", err)
	}

	// Check if the result matches the expected users
	if !reflect.DeepEqual(result, users) {
		t.Errorf("result does not match expected users: got %v, want %v", result, users)
	}
}
