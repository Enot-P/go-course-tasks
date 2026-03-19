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

func main() {
	// task 3.3.1
	user := User{}
	guest := Guest{}

	fmt.Println(printGreeting(user))
	fmt.Println(printGreeting(guest))

	// task 3.3.2
	rectangle := Rectangle{Width: 10, Height: 15}
	sqare := Square{Side: 5}
	shapes := []Shape{rectangle, sqare}

	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}

	// task 3.3.3
	athlete := Athlete{}
	fmt.Println(athlete.Run())

	var _ Runner = Athlete{} // Компилится, в чем смысл?
}
