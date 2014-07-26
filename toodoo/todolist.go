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
	todos := &TodoList{}
	return todos
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
