package main

import (
	"fmt"
)
func main() {
	todos := Todos{}//this is a slice of Todo structs
	todos.add("Buy groceries")
	todos.add("Walk the dog")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("%+v\n", todos)
}

