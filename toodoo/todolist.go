package toodoo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"text/template"
)

const todoTemplate = `{{ range .Todos }}[{{ with .Complete }}✔{{ else }} {{ end }}] {{ .Name }}
{{ end }}`

type TodoList struct {
	Todos []Todo
}

func (list *TodoList) Add(name string) {
	list.Todos = append(list.Todos, *newTodo(name))
}

func (list *TodoList) Remove(index int64) {
	list.Todos = append(list.Todos[:index], list.Todos[index+1:]...)
}

func (list *TodoList) Find(index int64) *Todo {
	return &list.Todos[index]
}

func (list *TodoList) List() {
	tmpl, err := template.New("todo").Parse(todoTemplate)
	if err != nil {
		log.Fatal(err)
	}

	template_err := tmpl.Execute(os.Stdout, list)
	if template_err != nil {
		log.Fatal(template_err)
	}
}

func newTodo(name string) *Todo {
	return &Todo{name, false}
}

func (list *TodoList) Read() {
	file, err := ioutil.ReadFile(saveFileLocation())
	todos := new([]Todo)
	if err == nil {
		json_err := json.Unmarshal(file, todos)
		if json_err != nil {
			log.Fatal(json_err)
		}

		list.Todos = *todos
	}
}

func (list *TodoList) Save() {
	buffer, marshal_err := json.MarshalIndent(list.Todos, "", "  ")
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
