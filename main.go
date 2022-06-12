package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

const Version = "1.0.1"

type Driver struct {
	mutex   sync.Mutex
	mutexes map[string]*sync.Mutex
	dir     string
}

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

	db, err := New(dir)
	if err != nil {
		fmt.Printf("Error! %v\n", err)
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			records, err := db.ReadAll("users")
			if err != nil {
				fmt.Printf("Error! %v", err)
			}
			fmt.Fprintf(w, "%v\n", records)
		case "POST":
			var user User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				fmt.Printf("Error! %v", err)
			}
			db.Write("users", user.Name, User{
				Name:    user.Name,
				Age:     user.Age,
				Contact: user.Contact,
				Address: user.Address,
			})
			fmt.Fprintf(w, "User %v added successfully!", user.Name)
		case "DELETE":
			var user User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				fmt.Printf("Error! %v", err)
			}
			db.Delete("users", user.Name)
			fmt.Fprintf(w, "User %v deleted successfully!", user.Name)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var user User
			db.Read("users", r.URL.Path[len("/users/"):], &user)
			fmt.Fprintf(w, "%v\n", user)
		}
	})

	fmt.Println("Listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func New(dir string) (*Driver, error) {
	dir = filepath.Clean(dir)

	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
	}

	return &driver, nil
}

func (driver *Driver) Write(table string, key string, value interface{}) error {
	driver.mutex.Lock()
	defer driver.mutex.Unlock()

	if _, ok := driver.mutexes[table]; !ok {
		driver.mutexes[table] = &sync.Mutex{}
	}

	driver.mutexes[table].Lock()
	defer driver.mutexes[table].Unlock()

	tableDir := filepath.Join(driver.dir, table)
	if err := os.MkdirAll(tableDir, 0755); err != nil {
		return err
	}

	file := filepath.Join(tableDir, key)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	return encoder.Encode(value)
}

func (driver *Driver) Read(table string, key string, value interface{}) error {
	driver.mutex.Lock()
	defer driver.mutex.Unlock()

	if _, ok := driver.mutexes[table]; !ok {
		driver.mutexes[table] = &sync.Mutex{}
	}

	driver.mutexes[table].Lock()
	defer driver.mutexes[table].Unlock()

	file := filepath.Join(driver.dir, table, key)
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	return decoder.Decode(value)
}

func (driver *Driver) ReadAll(table string) ([]interface{}, error) {
	driver.mutex.Lock()
	defer driver.mutex.Unlock()

	if _, ok := driver.mutexes[table]; !ok {
		driver.mutexes[table] = &sync.Mutex{}
	}

	driver.mutexes[table].Lock()
	defer driver.mutexes[table].Unlock()

	tableDir := filepath.Join(driver.dir, table)
	files, err := ioutil.ReadDir(tableDir)
	if err != nil {
		return nil, err
	}

	var records []interface{}
	for _, file := range files {
		var record interface{}
		if err := driver.Read(table, file.Name(), &record); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (driver *Driver) Delete(table string, key string) error {
	driver.mutex.Lock()
	defer driver.mutex.Unlock()

	if _, ok := driver.mutexes[table]; !ok {
		driver.mutexes[table] = &sync.Mutex{}
	}

	driver.mutexes[table].Lock()
	defer driver.mutexes[table].Unlock()

	file := filepath.Join(driver.dir, table, key)
	return os.Remove(file)
}

var employees = []struct {
	Name    string
	Age     json.Number
	Contact string
	Address Address
}{}
