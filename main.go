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
		{"Noah", "25", "2323434", "dmTech", Address{"Karlsruhe", "BaWü", "Germany", "12345"}},
		{"Paul", "25", "1234542", "Google", Address{"Berlin", "Berlin", "Germany", "45478"}},
		{"Akhil", "25", "2323434", "Facebook", Address{"Karlsruhe", "BaWü", "Germany", "66666"}},
		{"Antonia", "25", "2323434", "Sony", Address{"San Francisco", "Califronia", "Germany", "12345"}},
		{"Marius", "25", "2323434", "Bosch", Address{"Karlsruhe", "BaWü", "Germany", "12345"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, value)
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}
}
