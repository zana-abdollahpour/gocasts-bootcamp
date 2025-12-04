package main

import "fmt"

func determineName(name string) string {
	if name == "" {
		return "you"
	}
	return name
}

func shutout(name string) {
	fmt.Printf("one for %s, one for me\n", determineName(name))
}

func main() {
	names := [4]string{"James", "Abraham", "", "Monica"}

	for _, name := range names {
		shutout(name)
	}
}
