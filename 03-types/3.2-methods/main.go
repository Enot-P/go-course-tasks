package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

func (b Book) Summary() string {
	return fmt.Sprintf("%s (%d pages)", b.Title, b.Pages)
}

type Wallet struct {
	Balance int
}

func (w *Wallet) Deposit(amount int) {
	w.Balance += amount
}

func (w Wallet) DepositCopy(amount int) {
	w.Balance += amount
}

func resetScore(score *int) {
	*score = 0
}

type User struct {
	Name string
}

func (*User) printUserName(u *User) {
	if u == nil {
		fmt.Println("user is nil")
		return
	}
	fmt.Println("User exist: ", u.Name)
}

type Timer struct {
	Seconds int
	Running bool
}

func (t *Timer) Start() {
	t.Running = true
}

func (t *Timer) Stop() {
	t.Running = false
}

func (t *Timer) Status() string {
	if t == nil {
		return "t is nil"
	}
	if t.Running == true {
		return "running"
	} else {
		return "stopped"
	}
}

func main() {
	// task 3.2.1
	fmt.Printf("%s", Book{Title: "title", Pages: 12}.Summary()) // title (12 pages)

	// task 3.2.2
	wallet := Wallet{Balance: 0}
	wallet.Deposit(1)
	wallet.Deposit(1)
	wallet.Deposit(1)
	wallet.Deposit(1)
	fmt.Printf("\n\nwallet: %d\n", wallet.Balance) // wallet: 4

	// task 3.2.3
	wallet.DepositCopy(4)
	fmt.Printf("\n\nwallet: %d <--- after copy\n", wallet.Balance) // wallet: 4 <--- after copy

	// task 3.2.4
	score := 1111
	fmt.Printf("\nCurrent score:%d", score)
	resetScore(&score)
	fmt.Printf("\nScore after reset:%d\n\n", score)

	// task 3.2.5
	user := User{Name: "Nikita"}
	var user2 User
	var user3 *User = nil
	user.printUserName(&user)  // User exist:  Nikita
	user.printUserName(&user2) // User exist:
	user.printUserName(user3)  // user is nil

	// task 3.2.6
	timer := Timer{
		Seconds: 0,
		Running: false,
	}
	fmt.Println("\n")
	timer.Start()
	fmt.Println(timer.Status()) // running
	timer.Stop()
	fmt.Println(timer.Status()) // stopped
}
