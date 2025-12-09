package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID       int
	Title    string
	DueDate  string
	category string
	IsDone   bool
	UserID   int
}

var userStorage []User
var taskStorage []Task

var authenticatedUser *User

func main() {
	fmt.Println("Hello, welcome to TODO app!")

	command := flag.String("command", "no command", " command to run")
	flag.Parse()

	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please enter another command:")
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
		registerUser()
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
	var title, category, duedate string

	fmt.Println("Please enter the task's title:")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("Please enter the task's category:")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("Please enter the task's duedate:")
	scanner.Scan()
	duedate = scanner.Text()

	task := Task{
		ID:       (len(taskStorage) + 1),
		category: category,
		Title:    title,
		DueDate:  duedate,
		IsDone:   false,
		UserID:   authenticatedUser.ID,
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

	fmt.Println("Category:", title, color)
}

func registerUser() {
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

	user := User{ID: id, Name: name, Email: email, Password: password}
	userStorage = append(userStorage, user)

	fmt.Println("User created successfully!")
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
		if user.Email == email && user.Password == password {
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
