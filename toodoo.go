package main

import (
	"flag"
	"fmt"
	"strconv"

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
	case "remove":
		index, err := strconv.ParseInt(flag.Arg(1), 10, 0)
		if err != nil {
			panic(err)
		}
		remove(index)
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

func remove(index int64) {
	todos := toodoo.New()
	todos.Read()
	todos.Remove(index)
	todos.Save()
}
