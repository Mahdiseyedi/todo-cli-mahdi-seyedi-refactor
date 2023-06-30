package category

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

func TestWriteCategoryToFile(t *testing.T) {
	f := FileStore{
		Filepath:          "test.txt",
		serializationMode: consts.JsonSerializationMode,
	}

	category := models.Category{
		ID:     1,
		Title:  "Work",
		Color:  "blue",
		UserID: 2,
	}

	err := f.writeCategoryToFile(category)
	if err != nil {
		t.Errorf("writeCategoryToFile failed: %v", err)
	}

	data, err := ioutil.ReadFile(f.Filepath)
	if err != nil {
		t.Errorf("can't read test file: %v", err)
	}
	expected := "{\"ID\":1,\"Title\":\"Work\",\"Color\":\"blue\",\"UserID\":2}\n"
	if string(data) != expected {
		t.Errorf("test file does not match expected data: got %s, want %s", data, expected)
	}

	err = os.Remove(f.Filepath)
	if err != nil {
		t.Errorf("can't delete test file: %v", err)
	}
}

func TestCategoryJsonDeserializer(t *testing.T) {
	category := models.Category{
		ID:     1,
		Title:  "Work",
		Color:  "blue",
		UserID: 2,
	}

	data, err := json.Marshal(category)
	if err != nil {
		t.Errorf("can't marshal category struct to json: %v", err)
	}

	err = ioutil.WriteFile("test.txt", data, 0644)
	if err != nil {
		t.Errorf("can't write to test file: %v", err)
	}

	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		t.Errorf("can't read test file: %v", err)
	}

	category, err = JsonDeserializer(string(data))
	if err != nil {
		t.Errorf("JsonDeserializer failed: %v", err)
	}

	expected := models.Category{
		ID:     1,
		Title:  "Work",
		Color:  "blue",
		UserID: 2,
	}
	if !reflect.DeepEqual(category, expected) {
		t.Errorf("category does not match expected data: got %v, want %v", category, expected)
	}

	err = os.Remove("test.txt")
	if err != nil {
		t.Errorf("can't delete test file: %v", err)
	}
}

func TestCategoryTextDeserializer(t *testing.T) {

	category := models.Category{
		ID:     1,
		Title:  "Work",
		Color:  "Red",
		UserID: 10,
	}

	data := fmt.Sprintf("id: %d, title: %s, color: %s, userID: %d\n", category.ID, category.Title,
		category.Color, category.UserID)

	err := ioutil.WriteFile("test.txt", []byte(data), 0644)
	if err != nil {
		t.Errorf("can't write to test file: %v", err)
	}

	data2, err1 := ioutil.ReadFile("test.txt")
	if err1 != nil {
		t.Errorf("can't read test file: %v", err1)
	}

	category, err = TextDeserializer(string(data2))
	if err != nil {
		t.Errorf("TextDeserializer failed: %v", err)
	}

	expected := models.Category{
		ID:     1,
		Title:  "Work",
		Color:  "Red",
		UserID: 10,
	}
	if !reflect.DeepEqual(category, expected) {
		t.Errorf("category does not match expected data: got %v, want %v", category, expected)
	}

	err = os.Remove("test.txt")
	if err != nil {
		t.Errorf("can't delete test file: %v", err)
	}
}

func TestCategoryDeserializer(t *testing.T) {

	fs := FileStore{serializationMode: consts.TextSerializationMode}
	pData := []string{
		"id: 1, title: Work, color: blue, userID: 2",
		"id: 2, title: Home, color: green, userID: 3",
		"id: 3, title: Hobby, color: red, userID: 4",
		"id: 4, title: Travel, color: yellow, userID: 5",
		"id: invalid, title: None, color: black, userID: 6",
	}

	categories := fs.CategoryDeserializer(pData)
	if len(categories) != 4 {
		t.Errorf("expected 4 categories, got %d", len(categories))
	}

	expectedCategories := []models.Category{
		{ID: 1, Title: "Work", Color: "blue", UserID: 2},
		{ID: 2, Title: "Home", Color: "green", UserID: 3},
		{ID: 3, Title: "Hobby", Color: "red", UserID: 4},
		{ID: 4, Title: "Travel", Color: "yellow", UserID: 5},
	}

	for i, category := range categories {
		if category != expectedCategories[i] {
			t.Errorf("expected category %v, got %v", expectedCategories[i], category)
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
	category := models.Category{ID: 1, Title: "Work", Color: "blue", UserID: 2}

	fs.Save(category)

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

	expectedLine := "id: 1, title: Work, color: blue, userID: 2"
	if line != expectedLine {
		t.Errorf("expected line %s, got %s", expectedLine, line)
	}
}

func TestCreateNewCategory(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	fs := FileStore{Filepath: tmpfile.Name(), serializationMode: consts.TextSerializationMode}

	category := models.Category{
		Title:  "Movies",
		Color:  "Blue",
		UserID: 6,
	}

	result, err := fs.CreateNewCategory(category)
	if err != nil {
		t.Errorf("CreateNewCategory failed: %v", err)
	}

	expected := models.Category{
		ID:     1,
		Title:  "Movies",
		Color:  "Blue",
		UserID: 6,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result does not match expected data: got %v, want %v", result, expected)
	}
}

func TestGenerateID(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	fs := FileStore{Filepath: tmpfile.Name(), serializationMode: consts.TextSerializationMode}

	categories := []models.Category{
		{ID: 1, Title: "Movies", Color: "Blue", UserID: 6},
		{ID: 2, Title: "Books", Color: "Red", UserID: 7},
		{ID: 3, Title: "Games", Color: "Green", UserID: 8},
	}
	for _, category := range categories {
		err := fs.writeCategoryToFile(category)
		if err != nil {
			t.Errorf("can't write category to file: %v", err)
		}
	}

	id, err := fs.generateID()
	if err != nil {
		t.Errorf("generateID failed: %v", err)
	}

	expectedID := categories[len(categories)-1].ID + 1
	if id != expectedID {
		t.Errorf("expected ID %d, got %d", expectedID, id)
	}
}
