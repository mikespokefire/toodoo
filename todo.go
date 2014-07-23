package toodoo

type Todo struct {
	Name     string `json:"name"`
	Complete bool   `json:"is_complete"`
}
