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
}
