package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"text/template"
)

const todoTemplate = "[{{ with .Complete }}✔{{ else }} {{ end }}] {{ .Name }}\n"

type Todo struct {
	Name     string `json:"name"`
	Complete bool   `json:"is_complete"`
}

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
	usr, _ := user.Current()
	home_dir := usr.HomeDir

	file, err := ioutil.ReadFile(home_dir + "/.toodoos.json")
	if err != nil {
		log.Fatal(err)
	}

	var todos []Todo
	json_err := json.Unmarshal(file, &todos)
	if json_err != nil {
		log.Fatal(json_err)
	}

	t, err := template.New("todoTemplate").Parse(todoTemplate)
	if err != nil {
		panic(err)
	}

	for _, todo := range todos {
		err := t.Execute(os.Stdout, todo)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}

func add() {
	fmt.Println("Need to add some stuff")
}
