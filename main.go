package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"

	"todoapp/constants"
	"todoapp/contract"
	"todoapp/entity"
	"todoapp/filestore"
)

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

const userStoragePath = "user.txt"

var (
	userStorage     []entity.User
	taskStorage     []Task
	categoryStorage []Category

	authenticatedUser *entity.User
	serializationMode string
)

func main() {
	fmt.Println("Hello, welcome to TODO app!")

	enteredSerializationMode := flag.String("serialization-mode", constants.JsonSerializationMode, "serializationMode to write to file")

	switch *enteredSerializationMode {
	case constants.CustomSerializationMode:
		serializationMode = constants.CustomSerializationMode
	default:
		serializationMode = constants.JsonSerializationMode

	}

	userFileStore := filestore.New(userStoragePath, serializationMode)

	loadedUsers := userFileStore.Load()

	userStorage = append(userStorage, loadedUsers...)

	command := flag.String("command", "no command", " command to run")
	flag.Parse()

	for {
		runCommand(userFileStore, *command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("=> Please enter another command:")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(store contract.UserWriteStore, command string) {
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
		registerUser(store)
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

func registerUser(store contract.UserWriteStore) {
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

	user := entity.User{ID: id, Name: name, Email: email, Password: hashPassword(password)}
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

func hashPassword(password string) string {
	hashed := md5.Sum([]byte(password))

	return hex.EncodeToString(hashed[:])
}
