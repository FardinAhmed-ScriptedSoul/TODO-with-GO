package main

// import (
// 	"fmt"
// )
func main() {
	todos := Todos{}//this is a slice of Todo structs
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	todos.print()
	storage.Save([]Todos{todos})
}

