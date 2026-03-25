package main

import (
	"errors"
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

type PaymentProcessor interface {
	Pay(amount int) error
}

type CardProcessor struct {
	CardNumber string
}

func (c CardProcessor) Pay(amount int) error {
	if amount <= 0 {
		return errors.New("сумма платежа должна быть положительной")
	}

	if len(c.CardNumber) != 16 {
		return errors.New("неверный номер карты")
	}

	fmt.Printf("Оплата картой %s на сумму %d рублей выполнена успешно\n", c.CardNumber, amount)

	return nil
}

type CashProcessor struct{}

func (c CashProcessor) Pay(amount int) error {
	if amount <= 0 {
		return errors.New("сумма платежа должна быть положительной")
	}

	fmt.Printf("Оплата наличными на сумму %d рублей выполнена успешно\n", amount)
	return nil
}

func checkout(p PaymentProcessor, amount int) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Printf("Ошибка платежа: %v\n", err)
	}
}

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Printf("\nConsoleLogger => %s\n", message)
}

type PrefixLogger struct {
	Prefix string
}

func (pf PrefixLogger) Log(message string) {
	fmt.Printf("\n%s - %s\n", pf.Prefix, message)
}

func processOrder(logger Logger, id string) {
	logger.Log("order start")
	time.Sleep(1 * time.Second)
	logger.Log("order complete")
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

	// task 3.3.6
	cardProcessor := CardProcessor{CardNumber: "1234567890123456"}
	cashProcessor := CashProcessor{}

	// Тестируем различные платежи
	checkout(cardProcessor, 1500) // Оплата картой 1234567890123456 на сумму 1500 рублей выполнена успешно
	checkout(cashProcessor, 3500) // Оплата наличными на сумму 3500 рублей выполнена успешно
	invalidCard := CardProcessor{CardNumber: "123"}
	checkout(invalidCard, 1000)    // Ошибка платежа: неверный номер карты
	checkout(cashProcessor, 60000) // Оплата наличными на сумму 60000 рублей выполнена успешно
	checkout(cardProcessor, -100)  // Ошибка платежа: сумма платежа должна быть положительной

	// task 3.3.7
	console := ConsoleLogger{}
	prefix := PrefixLogger{Prefix: "order"}

	processOrder(console, "0")
	// ConsoleLogger => order start
	// ConsoleLogger => order complete

	processOrder(prefix, "1")
	// order - order start
	// order - order complete
}
