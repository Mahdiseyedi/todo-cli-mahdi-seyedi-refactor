package fileRepository

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

func (f FileStore) Save(u models.User) {
	//f.writeUserToFile(u)
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

func (f FileStore) Deserializer(pData []string) []models.User {
	var users []models.User

	for _, i := range pData {
		switch f.serializationMode {
		case consts.TextSerializationMode:
			user, err := TextDeserializer(i)
			if err != nil {
				fmt.Println(err)
				continue
			}
			users = append(users, user)
		case consts.JsonSerializationMode:
			user, err := JsonDeserializer(i)
			if err != nil {
				fmt.Println(err)
				continue
			}
			users = append(users, user)
		}
	}

	return users
}

func TextDeserializer(userStr string) (models.User, error) {

	fields := strings.Split(userStr, ",")

	if len(fields) != 4 {
		return models.User{}, fmt.Errorf("invalid user string: %s", userStr)
	}

	idStr := strings.TrimSpace(strings.Trim(fields[0], "id: "))
	name := strings.TrimSpace(strings.Trim(fields[1], "name:"))
	email := strings.TrimSpace(strings.Trim(fields[2], "email:"))
	password := strings.TrimSpace(strings.Trim(fields[3], "password:"))

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid id: %s", idStr)
	}

	user := models.User{
		ID:       id,
		Name:     name[6:],
		Email:    email[7:],
		Password: password[10:],
	}

	return user, nil
}

func JsonDeserializer(userStr string) (models.User, error) {
	var user models.User

	err := json.Unmarshal([]byte(userStr), &user)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid json: %s", userStr)
	}

	return user, nil
}

func (f FileStore) writeUserToFile(user models.User) error {
	file, err := os.OpenFile(f.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("can't create or open file: %w", err)
	}
	defer file.Close()

	var data []byte
	switch f.serializationMode {
	case consts.TextSerializationMode:
		data = []byte(fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n", user.ID, user.Name,
			user.Email, user.Password))
	case consts.JsonSerializationMode:
		data, err = json.Marshal(user)
		if err != nil {
			return fmt.Errorf("can't marshal user struct to json: %w", err)
		}
		data = append(data, '\n')
	default:
		return fmt.Errorf("invalid serialization mode")
	}

	n, err := io.WriteString(file, string(data))
	if err != nil {
		return fmt.Errorf("can't write to the file: %w", err)
	}

	// Print the number of written bytes for debugging purposes
	fmt.Println("numberOfWrittenBytes", n)

	// Return nil error if everything went well
	return nil
}
