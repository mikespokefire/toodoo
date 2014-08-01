package toodoo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"text/template"
)

const todoTemplate = `[{{ with .Complete }}✔{{ else }} {{ end }}] {{ .Name }}
`

type TodoList []*Todo

func New() *TodoList {
	todos := new(TodoList)
	return todos
}

func (todos *TodoList) Add(name string) {
	latest := *todos
	latest = append(latest, newTodo(name))
	*todos = latest
}

func (todos *TodoList) Remove(index int64) {
	latest := *todos
	latest = append(latest[:index], latest[index+1:]...)
	*todos = latest
}

func (todos TodoList) Find(index int64) *Todo {
	return todos[index]
}

func (todos *TodoList) List() {
	tmpl, err := template.New("todo").Parse(todoTemplate)
	if err != nil {
		panic(err)
	}

	for _, todo := range *todos {
		err := tmpl.Execute(os.Stdout, todo)
		if err != nil {
			panic(err)
		}
	}
}

func newTodo(name string) *Todo {
	return &Todo{name, false}
}

func (todos *TodoList) Read() {
	file, err := ioutil.ReadFile(saveFileLocation())
	if err == nil {
		json_err := json.Unmarshal(file, todos)
		if json_err != nil {
			log.Fatal(json_err)
		}
	}
}

func (todos *TodoList) Save() {
	buffer, marshal_err := json.MarshalIndent(todos, "", "  ")
	if marshal_err != nil {
		log.Fatal(marshal_err)
	}

	err := ioutil.WriteFile(saveFileLocation(), buffer, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func saveFileLocation() string {
	usr, _ := user.Current()
	home_dir := usr.HomeDir
	return home_dir + "/.toodoos.json"
}
