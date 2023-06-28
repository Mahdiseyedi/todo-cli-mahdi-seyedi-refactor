package task

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

func (f FileStore) Save(t models.Task) {
	f.writeTaskToFile(t)
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

func (f FileStore) TaskDeserializer(pData []string) []models.Task {
	var Tasks []models.Task

	for _, i := range pData {
		switch f.serializationMode {
		case consts.TextSerializationMode:
			Task, err := TextDeserializer(i)
			if err != nil {
				fmt.Println(err)
				continue
			}
			Tasks = append(Tasks, Task)
		case consts.JsonSerializationMode:
			Task, err := JsonDeserializer(i)
			if err != nil {
				fmt.Println(err)
				continue
			}
			Tasks = append(Tasks, Task)
		}
	}

	return Tasks
}

func TextDeserializer(taskStr string) (models.Task, error) {

	taskStr = strings.TrimRight(taskStr, "\n")
	fields := strings.Split(taskStr, ",")

	if len(fields) != 6 {
		return models.Task{}, fmt.Errorf("invalid task string: %s", taskStr)
	}

	idStr := fields[0][4:]
	title := fields[1][8:]
	dueDate := fields[2][10:]
	categoryIDStr := fields[3][13:]
	isDoneStr := fields[4][9:]
	userIDStr := fields[5][9:]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return models.Task{}, fmt.Errorf("invalid id: %s", idStr)
	}

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return models.Task{}, fmt.Errorf("invalid categoryID: %s", categoryIDStr)
	}

	isDone, err := strconv.ParseBool(isDoneStr)
	if err != nil {
		return models.Task{}, fmt.Errorf("invalid isDone: %s", isDoneStr)
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return models.Task{}, fmt.Errorf("invalid userID: %s", userIDStr)
	}
	fmt.Println(id)
	fmt.Println(title)
	fmt.Println(dueDate)
	fmt.Println(categoryID)
	fmt.Println(isDone)
	fmt.Println(userID)
	task := models.Task{
		ID:         id,
		Title:      title,
		DueDate:    dueDate,
		CategoryID: categoryID,
		IsDone:     isDone,
		UserID:     userID,
	}
	fmt.Println(task)
	return task, nil
}

func JsonDeserializer(TaskStr string) (models.Task, error) {
	var Task models.Task

	err := json.Unmarshal([]byte(TaskStr), &Task)
	if err != nil {
		return models.Task{}, fmt.Errorf("invalid json: %s", TaskStr)
	}

	return Task, nil
}

func (f FileStore) writeTaskToFile(task models.Task) error {
	file, err := os.OpenFile(f.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("can't create or open file: %w", err)
	}
	defer file.Close()

	var data []byte
	switch f.serializationMode {
	case consts.TextSerializationMode:
		data = []byte(fmt.Sprintf("id: %d, title: %s, dueDate: %s, categoryID: %d, isDone: %t, userID: %d\n", task.ID, task.Title,
			task.DueDate, task.CategoryID, task.IsDone, task.UserID))
	case consts.JsonSerializationMode:
		data, err = json.Marshal(task)
		if err != nil {
			return fmt.Errorf("can't marshal task struct to json: %w", err)
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
