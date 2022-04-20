package main

import (
	"encoding/json"
	"fmt"
)

const Version = "1.0.1"

type Address struct {
	City    string
	Door_No json.Number
	Street  string
	Country string
	Pincode json.Number
	State   string
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Printf("Error! %v\n", err)
	}

	employees := []User{
		{"Lakshmi Narasimman", "22", "9453754225", Address{"Theni", "453", "AVR Compound", "India", "625531", "Tamil Nadu"}},
		{"Kishore", "30", "9876543234", Address{"Mumbai", "742", "Main Road", "India", "625637", "Maharastra"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Printf("Error! %v", err)
	}
	fmt.Printf("%v\n", records)

}
<<<<<<< HEAD
=======

func New() {
	// Hello
}
>>>>>>> a92d978614a31c29208c9dfbde514b6b9788a6e6
