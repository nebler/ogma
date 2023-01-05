package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

type Database interface {
}

func New(dir string, huh string) {

}

func main() {
	dir := "./"

	db, err := New(dir, "")
	if err != nil {
		fmt.Print("error")
	}

	employees := []User{
		{"John", "23", "2323434", "Myril Tech", Address{"test1", "test", "test", "41434434"}},
		{"John", "23", "2323434", "Myril Tech", Address{"test1", "test", "test", "41434434"}},
		{"John", "23", "2323434", "Myril Tech", Address{"test1", "test", "test", "41434434"}},
		{"John", "23", "2323434", "Myril Tech", Address{"test1", "test", "test", "41434434"}},
	}
}
