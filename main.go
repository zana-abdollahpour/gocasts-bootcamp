package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID         int
	Title      string
	DueDate    string
	categoryID int
	IsDone     bool
	UserID     int
}

type Category struct {
	ID     int
	Title  string
	Color  string
	UserID int
}

const (
	userStoragePath         = "user.txt"
	customSerializationMode = "custom-made"
	jsonSerializationMode   = "json"
)

var (
	userStorage     []User
	taskStorage     []Task
	categoryStorage []Category

	authenticatedUser *User
	serializationMode string
)

var userFileStore = fileStore{
	filePath: userStoragePath,
}

func main() {
	fmt.Println("Hello, welcome to TODO app!")

	enteredSerializationMode := flag.String("serialization-mode", jsonSerializationMode, "serializationMode to write to file")

	switch *enteredSerializationMode {
	case customSerializationMode:
		serializationMode = customSerializationMode
	default:
		serializationMode = jsonSerializationMode

	}

	loadUserFromStorage(userFileStore, serializationMode)

	command := flag.String("command", "no command", " command to run")
	flag.Parse()

	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("=> Please enter another command:")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(command string) {
	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()

		if authenticatedUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser(userFileStore)
	case "login":
		login()
	case "list-tasks":
		listTasks()
	case "exit":
		os.Exit(0)
	default:
		fmt.Printf("'%s' command is not valid!\n", command)
	}
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, categoryID, duedate string

	fmt.Println("Please enter the task's title:")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("Please enter the task's category id:")
	scanner.Scan()
	categoryID = scanner.Text()

	fmt.Println("Please enter the task's duedate:")
	scanner.Scan()
	duedate = scanner.Text()

	parsedCategoryID, err := strconv.Atoi(categoryID)

	isFound := false
	for _, c := range categoryStorage {
		if c.ID == parsedCategoryID && c.UserID == authenticatedUser.ID {
			isFound = true

			break
		}
	}

	if !isFound {
		fmt.Println("category id is not found.")

		return
	}

	if err != nil {
		fmt.Printf("category id is not a valid integer, %v\n", err)

		return
	}

	task := Task{
		ID:         (len(taskStorage) + 1),
		categoryID: parsedCategoryID,
		Title:      title,
		DueDate:    duedate,
		IsDone:     false,
		UserID:     authenticatedUser.ID,
	}

	taskStorage = append(taskStorage, task)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)

	var title, color string

	fmt.Println("Please enter the category's title:")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("Please enter the category's color:")
	scanner.Scan()
	color = scanner.Text()

	category := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, category)
}

type userWriteStore interface {
	Save(u User)
}

type userReadStore interface {
	Load(serializationMode string) []User
}

func registerUser(store userWriteStore) {
	scanner := bufio.NewScanner(os.Stdin)
	var name, email, password string

	fmt.Println("Please enter your name:")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Please enter your email:")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Please enter your password:")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("User:", email, password)

	id := len(userStorage) + 1

	user := User{ID: id, Name: name, Email: email, Password: hashPassword(password)}
	userStorage = append(userStorage, user)

	store.Save(user)

}

func login() {
	fmt.Println("---> LOGIN PROCESS <---")
	scn := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter your email:")
	scn.Scan()
	email := scn.Text()

	fmt.Println("Please enter your password:")
	scn.Scan()
	password := scn.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == hashPassword(password) {
			authenticatedUser = &user
			fmt.Println("You have logged in successfully :)")

			break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("The email or password is NOT correct!")
	}
}

func listTasks() {
	for _, task := range taskStorage {
		if task.ID == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}

func loadUserFromStorage(store userReadStore, serializationMode string) {
	users := store.Load(serializationMode)
	userStorage = append(userStorage, users...)
}

func deserializeFromCustomMode(userString string) (User, error) {

	if userString == "" {
		return User{}, errors.New("user string is empty")
	}

	userFields := strings.SplitSeq(userString, ",")

	user := User{}
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

				return User{}, errors.New("strconv error happened")
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

func hashPassword(password string) string {
	hashed := md5.Sum([]byte(password))

	return hex.EncodeToString(hashed[:])
}

type fileStore struct {
	filePath string
}

func (f fileStore) Save(user User) {
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
	switch serializationMode {
	case customSerializationMode:
		userDataString := fmt.Sprintf(
			"id: %d, name: %s, email: %s, password: %s\n",
			user.ID, user.Name, user.Email, user.Password,
		)
		userData = []byte(userDataString)

	case jsonSerializationMode:
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

func (f fileStore) Load(serializationMode string) []User {
	var uStore []User

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
		var userStruct = User{}

		switch serializationMode {
		case customSerializationMode:
			userStruct, deserializeErr := deserializeFromCustomMode(userData)
			if deserializeErr != nil {
				fmt.Println("can't deserialize user record to user struct in custom mode")

				return nil
			}

			uStore = append(uStore, userStruct)

		case jsonSerializationMode:
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
