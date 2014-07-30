package toodoo

type Todo struct {
	Name     string `json:"name"`
	Complete bool   `json:"is_complete"`
}

func (todo *Todo) MarkAsComplete() {
	todo.Complete = true
}

func (todo *Todo) MarkAsIncomplete() {
	todo.Complete = false
}
