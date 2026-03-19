package main

import (
	"fmt"
	"time"
)

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

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Square struct {
	Side float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

type Runner interface {
	Run() string
}

type Athlete struct{}

func (a Athlete) Run() string {
	return "Athlete is running"
}

type Starter interface {
	Start()
}
type Stopper interface {
	Stop()
}
type Machine interface {
	Starter
	Stopper
}
type Engine struct{}

func (e Engine) Start() {
	fmt.Println("Machine is starting")
}

func (e Engine) Stop() {
	fmt.Println("Machine is stopped")
}

func runCycle(m Machine) {
	m.Start()
	time.Sleep(1 * time.Second)
	m.Stop()
}

func printStringLength(x any) {
	result, ok := x.(string)
	if !ok {
		fmt.Println("Not a string")
		return
	}
	fmt.Println(len(result))
}

func main() {
	// task 3.3.1
	user := User{}
	guest := Guest{}

	fmt.Println(printGreeting(user))  // Hi from User
	fmt.Println(printGreeting(guest)) // Hi from Guest

	// task 3.3.2
	rectangle := Rectangle{Width: 10, Height: 15}
	sqare := Square{Side: 5}
	shapes := []Shape{rectangle, sqare}

	for _, shape := range shapes { // 150
		fmt.Println(shape.Area()) // 25
	}

	// task 3.3.3
	athlete := Athlete{}
	fmt.Println(athlete.Run()) // Athlete is running

	var _ Runner = Athlete{} // Компилится, в чем смысл?

	// task 3.3.4
	bmwEngine := Engine{} // Machine is starting
	runCycle(bmwEngine)   // Machine is stopped

	// task 3.3.5
	printStringLength(42)        // Not a string
	printStringLength("2334324") // 7
}
