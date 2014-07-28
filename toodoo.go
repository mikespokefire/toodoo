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
		add()
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

func add() {
	fmt.Println("Need to add some stuff")
}
