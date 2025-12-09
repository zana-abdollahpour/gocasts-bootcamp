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

var userStorage []User

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
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Printf("'%s' command is not valid!\n", command)
	}
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, category, duedate string

	fmt.Println("Please enter the task's name:")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Please enter the task's category:")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("Please enter the task's duedate:")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("Task:", name, category, duedate)
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

	fmt.Printf("UserStorage: %+v", userStorage)
}

func login() {
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("Please enter the your email:")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Please enter the your password:")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("LoginData:", email, password)
}
