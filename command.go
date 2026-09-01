package main

import (
	"flag"
	"fmt"
)

type CmdFlags struct {
	Add string
	Del int
	Edit string
	Toggle int
	List bool
}

func NewCmdFlags() *CmdFlags {
	cf := &CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by index")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index (format: index:new title)")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle the completion status of a todo by index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return cf
}

func (cf *CmdFlags) Execute(todos *Todos){
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Edit != "":
		var index int
		var newTitle string
		n, err := fmt.Sscanf(cf.Edit, "%d:%s", &index, &newTitle)	
	if err != nil || n != 2 {
		fmt.Println("Invalid format for edit. Use index:new_title")
		return
	
	}
	todos.edit(index, newTitle)
default:
	fmt.Println("Invalid command!!")
}
}