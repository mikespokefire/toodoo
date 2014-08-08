package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mikespokefire/toodoo/toodoo"
)

func main() {
	listName := flag.String("list", "default", "The todo list that you want to manipulate")
	flag.Parse()

	cmd := flag.Arg(0)

	switch cmd {
	case "list":
		list(*listName)
	case "add":
		arguments := flag.Args()[1:]
		name := strings.Join(arguments, " ")
		add(*listName, name)
	case "remove":
		index, err := strconv.ParseInt(flag.Arg(1), 10, 0)
		if err != nil {
			panic(err)
		}
		remove(*listName, index)
	case "help":
		usage()
	case "complete":
		index, err := strconv.ParseInt(flag.Arg(1), 10, 0)
		if err != nil {
			panic(err)
		}
		complete(*listName, index)
	case "incomplete":
		index, err := strconv.ParseInt(flag.Arg(1), 10, 0)
		if err != nil {
			panic(err)
		}
		incomplete(*listName, index)
	case "":
		usage()
		os.Exit(1)
	default:
		arguments := flag.Args()[0:]
		name := strings.Join(arguments, " ")
		add(*listName, name)
	}
}

func usage() {
	fmt.Println(`Toodoo is a tool for storing todos on your computer.

Usage:

	toodoo [command|arguments]

The commands are:

	list            list your todos
	add             add a todo (add is not necessary)
	complete        mark a todo as complete
	incomplete      mark a todo as incomplete
	remove          remove a todo
	help            this help message
	version         print the version number`)
}

func list(listName string) {
	todos := toodoo.TodoList{
		Name: listName,
	}
	todos.Read()
	todos.List()
}

func add(listName string, name string) {
	todos := toodoo.TodoList{
		Name: listName,
	}
	todos.Read()
	todos.Add(name)
	todos.Save()
}

func remove(listName string, index int64) {
	todos := toodoo.TodoList{
		Name: listName,
	}
	todos.Read()
	todos.Remove(index)
	todos.Save()
}

func complete(listName string, index int64) {
	todos := toodoo.TodoList{
		Name: listName,
	}
	todos.Read()
	todos.Find(index).MarkAsComplete()
	todos.Save()
}

func incomplete(listName string, index int64) {
	todos := toodoo.TodoList{
		Name: listName,
	}
	todos.Read()
	todos.Find(index).MarkAsIncomplete()
	todos.Save()
}
