package category

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"todo-cli-refactor/consts"
	"todo-cli-refactor/models"
)

type FileStore struct {
	Filepath          string
	serializationMode string
}

func New(path, serializationMode string) FileStore {
	return FileStore{Filepath: path, serializationMode: serializationMode}
}

func (f FileStore) Save(t models.Category) {
	f.writeCategoryToFile(t)
}

func (f FileStore) Load() ([]string, error) {
	var err error
	var pData []string

	file, oErr := os.Open(f.Filepath)
	if oErr != nil {
		err = oErr
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pData = append(pData, scanner.Text())
	}

	return pData, err
}

func (f FileStore) CategoryDeserializer(pData []string) []models.Category {
	var Categorys []models.Category

	for _, i := range pData {
		switch f.serializationMode {
		case consts.TextSerializationMode:
			Category, err := TextDeserializer(i)
			if err != nil {
				fmt.Println(err)
				continue
			}
			Categorys = append(Categorys, Category)
		case consts.JsonSerializationMode:
			Category, err := JsonDeserializer(i)
			if err != nil {
				fmt.Println(err)
				continue
			}
			Categorys = append(Categorys, Category)
		}
	}

	return Categorys
}

func TextDeserializer(categoryStr string) (models.Category, error) {

	categoryStr = strings.TrimRight(categoryStr, "\n")
	fields := strings.Split(categoryStr, ",")

	if len(fields) != 4 {
		return models.Category{}, fmt.Errorf("invalid category string: %s", categoryStr)
	}
	idStr := fields[0][4:]
	title := fields[1][8:]
	color := fields[2][8:]
	userIDStr := fields[3][9:]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return models.Category{}, fmt.Errorf("invalid id: %s", idStr)
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return models.Category{}, fmt.Errorf("invalid userID: %s", userIDStr)
	}

	category := models.Category{
		ID:     id,
		Title:  title,
		Color:  color,
		UserID: userID,
	}

	return category, nil
}

func JsonDeserializer(CategoryStr string) (models.Category, error) {
	var Category models.Category

	err := json.Unmarshal([]byte(CategoryStr), &Category)
	if err != nil {
		return models.Category{}, fmt.Errorf("invalid json: %s", CategoryStr)
	}

	return Category, nil
}

func (f FileStore) writeCategoryToFile(category models.Category) error {
	file, err := os.OpenFile(f.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("can't create or open file: %w", err)
	}
	defer file.Close()

	var data []byte
	switch f.serializationMode {
	case consts.TextSerializationMode:
		data = []byte(fmt.Sprintf("id: %d, title: %s, color: %s, userID: %d\n", category.ID, category.Title,
			category.Color, category.UserID))
	case consts.JsonSerializationMode:
		data, err = json.Marshal(category)
		if err != nil {
			return fmt.Errorf("can't marshal category struct to json: %w", err)
		}
		data = append(data, '\n')
	default:
		return fmt.Errorf("invalid serialization mode")
	}

	n, err := io.WriteString(file, string(data))
	if err != nil {
		return fmt.Errorf("can't write to the file: %w", err)
	}

	fmt.Println("numberOfWrittenBytes", n)

	return nil
}

func (f FileStore) CreateNewCategory(category models.Category) (models.Category, error) {

	refID, err := f.generateID()
	if err != nil {
		return models.Category{}, err
	}
	category.ID = refID

	err = f.writeCategoryToFile(category)
	if err != nil {
		return models.Category{}, fmt.Errorf("can't write category to file: %v", err)
	}

	return category, nil
}

func (f FileStore) generateID() (int, error) {

	lines, err := f.Load()
	if err != nil {
		return 0, fmt.Errorf("files can't load for counting lines: %w", err)
	}

	if len(lines) == 0 {
		return 1, nil
	}

	lastLine := lines[len(lines)-1]

	lastCategory := f.CategoryDeserializer([]string{lastLine})
	if err != nil {
		return 0, fmt.Errorf("can't deserialize last line to category: %w", err)
	}

	return lastCategory[0].ID + 1, nil
}
