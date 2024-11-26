package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add string
	Delete int
	Edit string
	Toggle int
	List bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "Add", "", "Add a new Todo")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit a todo item by index")
	flag.IntVar(&cf.Delete, "Delete", -1, "Delete a todo item")
	flag.IntVar(&cf.Toggle, "Toggle", -1, "Toggle a todo item finished or unfinished")
	flag.BoolVar(&cf.List, "List", false, "List all todo items")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List: 
		todos.print()

	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit, Please use id:new_title")
			os.Exit(1)
		}
		
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Delete != -1:
		todos.delete(cf.Delete)

	default:
		fmt.Println("Invalid Command")
	}
}