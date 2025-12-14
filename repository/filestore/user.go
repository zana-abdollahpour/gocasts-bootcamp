package filestore

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"todoapp/constants"
	"todoapp/entity"
)

type FileStore struct {
	filePath          string
	serializationMode string
}

func New(path, serializationMode string) FileStore {
	return FileStore{filePath: path, serializationMode: serializationMode}
}

func (f FileStore) Save(user entity.User) {
	const (
		flag       = os.O_APPEND | os.O_CREATE | os.O_WRONLY
		permission = os.FileMode(0644)
	)

	file, openErr := os.OpenFile(f.filePath, flag, permission)

	if openErr != nil {
		fmt.Printf("can't access or open the file %v\n", openErr)

		return
	}

	defer file.Close()

	var userData []byte
	switch f.serializationMode {
	case constants.CustomSerializationMode:
		userDataString := fmt.Sprintf(
			"id: %d, name: %s, email: %s, password: %s\n",
			user.ID, user.Name, user.Email, user.Password,
		)
		userData = []byte(userDataString)

	case constants.JsonSerializationMode:
		serializedData, err := json.Marshal(user)
		if err != nil {
			fmt.Println("Can't marshal user struct to json", err)

			return
		}

		userData = fmt.Appendf(nil, "%s\n", serializedData)
	default:
		fmt.Println("Invalid serialization mode!")

		return
	}

	_, writeError := file.Write([]byte(userData))

	if writeError != nil {
		fmt.Printf("can't write to the file %v\n", writeError)
	}

	fmt.Println("User created successfully!")
}

func (f FileStore) Load() []entity.User {
	var uStore []entity.User

	file, err := os.Open(f.filePath)

	if err != nil {
		fmt.Println("Can't open the file", err)
	}

	var data = make([]byte, 1024)
	_, openError := file.Read(data)

	if openError != nil {
		fmt.Println("Can't read from the file", openError)
	}

	dataString := string(data)
	dataString = strings.Trim(dataString, "\n")

	usersSlice := strings.SplitSeq(dataString, "\n")

	for userData := range usersSlice {
		var userStruct = entity.User{}

		switch f.serializationMode {
		case constants.CustomSerializationMode:
			userStruct, deserializeErr := deserializeFromCustomMode(userData)
			if deserializeErr != nil {
				fmt.Println("can't deserialize user record to user struct in custom mode")

				return nil
			}

			uStore = append(uStore, userStruct)

		case constants.JsonSerializationMode:
			if userData[0] != '{' && userData[len(userData)-1] != '}' {
				continue
			}

			deserializeErr := json.Unmarshal([]byte(userData), &userStruct)
			if deserializeErr != nil {
				fmt.Println("can't deserialize user record to user struct in json mode")

				return nil
			}

		default:
			fmt.Println("invalid serialization mode")

			return nil
		}

		uStore = append(uStore, userStruct)
	}

	return uStore
}

func deserializeFromCustomMode(userString string) (entity.User, error) {

	if userString == "" {
		return entity.User{}, errors.New("user string is empty")
	}

	userFields := strings.SplitSeq(userString, ",")

	user := entity.User{}
	for userField := range userFields {
		keyValuePair := strings.Split(userField, ": ")

		if len(keyValuePair) != 2 {
			continue
		}

		fieldName := strings.Trim(keyValuePair[0], " ")
		fieldValue := strings.Trim(keyValuePair[1], " ")

		switch fieldName {
		case "id":
			id, err := strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("strconv.Atoi error", err)

				return entity.User{}, errors.New("strconv error happened")
			}
			user.ID = id

		case "name":
			user.Name = fieldValue
		case "email":
			user.Email = fieldValue
		case "password":
			user.Password = fieldValue

		}
	}

	return user, nil
}
