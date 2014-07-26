package main

import (
	"flag"
	"fmt"
	"github.com/mikespokefire/toodoo/toodoo"
)

func main() {
	flag.Parse()

	cmd := flag.Arg(0)

	switch cmd {
	case "list":
		list()
	case "add":
		name := flag.Arg(1)
		add(name)
	default:
		usage()
	}
}

func usage() {
	fmt.Println(`Toodoo is a tool for storing todos on your computer.

Usage:

	toodoo command [arguments]

The commands are:

	list            list your todos
	add             add a todo
	complete        mark a todo as complete
	remove          remove a todo
	version         print the version number`)
}

func list() {
	todos := toodoo.New()
	todos.Read()
	todos.List()
}

func add(name string) {
	todos := toodoo.New()
	todos.Read()
	todos.Add(name)
	todos.Save()
}
