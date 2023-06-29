package user

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
	f.writeUserToFile(u)
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

func (f FileStore) UserDeserializer(pData []string) []models.User {
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

	idStr := strings.TrimSpace(strings.Trim(fields[0], "ID: "))
	name := strings.TrimSpace(strings.Trim(fields[1], "Name:"))
	email := strings.TrimSpace(strings.Trim(fields[2], "Email:"))
	password := strings.TrimSpace(strings.Trim(fields[3], "Password:"))

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
		data = []byte(fmt.Sprintf("ID: %d, Name: %s, Email: %s, Password: %s\n", user.ID, user.Name,
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

	fmt.Println("numberOfWrittenBytes", n)

	return nil
}

func (f FileStore) CreateNewUser(user models.User) (models.User, error) {

	refID, err := f.generateID()
	if err != nil {
		return models.User{}, err
	}
	user.ID = refID

	//TODO : implement this section for hash password
	//	user.Password = hashThePassword(user.Password)

	err = f.writeUserToFile(user)
	if err != nil {
		return models.User{}, fmt.Errorf("can't write user to file: %v", err)
	}

	return user, nil
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

	lastUser := f.UserDeserializer([]string{lastLine})
	if err != nil {
		return 0, fmt.Errorf("can't deserialize last line to user: %w", err)
	}

	return lastUser[0].ID + 1, nil
}

func (f FileStore) ListUsers() ([]models.User, error) {

	lines, err := f.Load()

	if err != nil {
		return nil, fmt.Errorf("can't read from file: %w", err)
	}

	var users []models.User

	for _, line := range lines {
		user := f.UserDeserializer([]string{line})
		if err != nil {
			fmt.Println(err)
			continue
		}

		users = append(users, user[0])
	}

	return users, nil
}
