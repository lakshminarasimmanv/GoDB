# DB Driver

This is a simple DB driver program that can be used to store and retrieve data from a JSON file.

## Usage

To use this program, you first need to create a new Driver instance, passing in the directory where you want to store your JSON files:

```go
dir := "./"

db, err := New(dir)
if err != nil {
    fmt.Printf("Error! %v\n", err)
}
```

Once you have a Driver instance, you can use its `Write`, `Read`, `ReadAll`, and `Delete` methods to store and retrieve data from your JSON files.

## Write

The `Write` method takes three arguments: the table name, the key, and the value. The table name is the name of the JSON file in which the data will be stored, the key is the name of the JSON object that will be stored, and the value is the data to be stored.

```go
db.Write("users", value.Name, User{
    Name:    value.Name,
    Age:     value.Age,
    Contact: value.Contact,
    Address: value.Address,
})
```

## Read

The `Read` method takes three arguments: the table name, the key, and a pointer to the variable in which the data will be stored. The table name is the name of the JSON file from which the data will be retrieved, the key is the name of the JSON object that will be retrieved, and the pointer is the variable in which the data will be stored.

```go
var user User
db.Read("users", r.URL.Path[len("/users/"):], &user)
```

## ReadAll

The `ReadAll` method takes one argument: the table name. The table name is the name of the JSON file from which the data will be retrieved.

This method returns an array of interface{} values, each of which contains the data for one JSON object.

```go
records, err := db.ReadAll("users")
if err != nil {
    fmt.Printf("Error! %v", err)
}
fmt.Printf("%v\n", records)
```

## Delete

The `Delete` method takes two arguments: the table name and the key. The table name is the name of the JSON file from which the data will be deleted, and the key is the name of the JSON object that will be deleted.

```go
db.Delete("users", user.Name)
```

1. Enter the following command in your terminal: go run main.go 
2. The program will start running and you will be able to access it at http://localhost:8080
3. To add a new user, you need to send a POST request to http://localhost:8080/users with the user object in the request body.

For example, if you want to add a user named John, you can use the following curl command:

curl -X POST \
  http://localhost:8080/users \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "John",
  "age": "21",
  "contact": "123-456-7890",
  "address": {
    "city": "New York",
    "door_no": "123",
    "street": "Main Street",
    "country": "USA",
    "pincode": "12345",
    "state": "NY"
  }
}'
