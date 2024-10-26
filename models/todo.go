package models

type Todo struct {
	Id        int      `json:"i"`
	Titile    string   `json:"t"`
	Completed bool     `json:"c"`
}

type ITodoCrudAPI interface {
	GetTodoList() []Todo
	GetTodo(int) (Todo, error)
	CreateTodo(Todo) error
	UpdateTodo(Todo) error
	DeleteTodo(int) error
}