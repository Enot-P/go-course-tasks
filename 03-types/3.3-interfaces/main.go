package main

import "fmt"

type Greeter interface {
	Greet() string
}

type User struct{}

func (u User) Greet() string {
	return "Hi from User"
}

type Guest struct{}

func (g Guest) Greet() string {
	return "Hi from Guest"
}

func printGreeting(g Greeter) string {
	// return "Hi from printGreeting"
	return g.Greet()
}

func main() {
	// task 3.3.1
	user := User{}
	guest := Guest{}

	fmt.Println(printGreeting(user))
	fmt.Println(printGreeting(guest))
}
