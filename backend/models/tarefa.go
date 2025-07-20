package models

type Tarefa struct {
	ID    int `json:"id,omitempty"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
