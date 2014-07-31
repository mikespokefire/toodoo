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
	flag.Parse()

	cmd := flag.Arg(0)

	switch cmd {
	case "list":
		list()
	case "add":
		arguments := flag.Args()[1:]
		name := strings.Join(arguments, " ")
		add(name)
	case "remove":
		index, err := strconv.ParseInt(flag.Arg(1), 10, 0)
		if err != nil {
			panic(err)
		}
		remove(index)
	case "help":
		usage()
	default:
		usage()
		os.Exit(1)
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
	help            this help message
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
