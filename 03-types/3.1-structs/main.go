package main

import (
	"fmt"
)

type Profile struct {
	Name     string
	Age      int
	isActive bool
}

type AppConfig struct {
	Host  string
	Port  int
	Debug bool
}

type Address struct {
	City   string
	Street string
}
type Employee struct {
	Name    string
	Address Address
}

type Package struct {
	ID     string
	Weight int
}
type Destination struct {
	City string
	Zip  string
}
type Shipment struct {
	Package     Package
	Destination Destination
}

type Audit struct {
	CreatedAt string
	UpdatedAt string
}
type Article struct {
	Title string
	Audit
}

func main() {
	// task 3.1.1
	user := Profile{
		Name:     "Nikita",
		Age:      23,
		isActive: true,
	}
	fmt.Println("User: = ", user) // User: =  {Nikita 23 true}

	// task 3.1.2
	var cfg AppConfig = AppConfig{
		Host:  "cofig",
		Port:  12,
		Debug: true,
	}

	fmt.Println("AppConfig = ", cfg)
	fmt.Printf("\nHost:%s\nPort:%d\nisActive:%t\n", cfg.Host, cfg.Port, cfg.Debug) // Host:cofig Port:12 isActive:true

	// task 3.1.3
	employe := Employee{
		Name: "Employe",
		Address: Address{
			City:   "Moscow",
			Street: "Krasnya",
		},
	}
	fmt.Printf("\n%s:%s, %s\n", employe.Name, employe.Address.City, employe.Address.Street) // Employe:Moscow, Krasnya

	// task 3.1.4
	ship := Shipment{
		Package: Package{
			ID:     "IDIDID",
			Weight: 100,
		},
		Destination: Destination{
			City: "Moscow",
			Zip:  "Zip",
		},
	}

	fmt.Printf("\nId:%s | Address: %s\n", ship.Package.ID, ship.Destination.City) // Id:IDIDID | Address: Moscow

	// task 3.1.5
	article := Article{
		Title: "Title",
		Audit: Audit{
			CreatedAt: "today",
			UpdatedAt: "today",
		},
	}

	fmt.Printf("\nTitle:%s | CreatedAt: %s | UpdatedAt: %s\n", article.Title, article.CreatedAt, article.UpdatedAt) // Title:Title | CreatedAt: today | UpdatedAt: today
}
