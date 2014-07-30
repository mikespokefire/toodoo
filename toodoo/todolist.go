package toodoo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"text/template"
)

const todoTemplate = "[{{ with .Complete }}✔{{ else }} {{ end }}] {{ .Name }}\n"

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

func (todos *TodoList) List() {
	t, err := template.New("todoTemplate").Parse(todoTemplate)
	if err != nil {
		panic(err)
	}

	for _, todo := range *todos {
		err := t.Execute(os.Stdout, todo)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}

func newTodo(name string) *Todo {
	return &Todo{name, false}
}

func (todos *TodoList) Read() {
	usr, _ := user.Current()
	home_dir := usr.HomeDir

	file, err := ioutil.ReadFile(home_dir + "/.toodoos.json")
	if err != nil {
		log.Fatal(err)
	}

	json_err := json.Unmarshal(file, todos)
	if json_err != nil {
		log.Fatal(json_err)
	}
}

func (todos *TodoList) Save() {
	usr, _ := user.Current()
	home_dir := usr.HomeDir

	buffer, marshal_err := json.Marshal(todos)
	if marshal_err != nil {
		log.Fatal(marshal_err)
	}

	err := ioutil.WriteFile(home_dir+"/.toodoos.json", buffer, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
