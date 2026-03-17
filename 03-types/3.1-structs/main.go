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
	fmt.Printf("\nHost:%s\nPort:%d\nisActive:%t", cfg.Host, cfg.Port, cfg.Debug) // Host:cofig Port:12 isActive:true
}
