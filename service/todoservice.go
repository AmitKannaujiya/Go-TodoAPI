package service

import (
	"errors"
	m "go-todo/models"
	"sync"
)

type TodoService struct{
	todos []m.Todo
	todosMap map[int]*m.Todo
}

var Singleton *TodoService

var once sync.Once

func GetTodoService() *TodoService {
	if Singleton ==  nil {
		once.Do(func() {
			Singleton = &TodoService{
				todos: make([]m.Todo, 0),
				todosMap: make(map[int]*m.Todo),
			}
		})
	}
	return Singleton
}
// GetTodoList() []Todo
// GetTodo(int) error
// CreateTodo(Todo) error
// UpdateTodo(Todo) error
// DeleteTodo(Todo) error
func (ts *TodoService) GetTodoList() []m.Todo {
	return ts.todos
}

func (ts *TodoService) GetTodo(id int) (m.Todo, error) {
	if todo, exists := ts.todosMap[id]; exists {
		return *todo, nil
	}
	var t m.Todo
	return t, errors.New("TODO NOT FOUND")
}

func (ts *TodoService) CreateTodo(todo m.Todo) error {
	id := len(ts.todos) + 1 
	todo.Id = id
	ts.todosMap[id] = &todo
	ts.todos = append(ts.todos, todo)
	return nil
}

func (ts *TodoService) UpdateTodo(todo m.Todo) error {
	if td, exists := ts.todosMap[todo.Id]; exists {
		td.Completed = todo.Completed
		td.Titile = todo.Titile
		ts.todosMap[todo.Id] = td
		for i, t := range ts.todos {
			if t.Id == todo.Id {
				ts.todos[i] = *td
				break
			}
		}
		return nil
	}
	return errors.New("TODO NOT FOUND")
}

func (ts *TodoService) DeleteTodo(id int) error {
	if td, exists := ts.todosMap[id]; exists {
		for i, t:= range ts.todos {
			if t.Id == td.Id {
				ts.todos = append(ts.todos[:i], ts.todos[i+1:]...)
				break
			}
		}
		return nil
	}
	return errors.New("TODO NOT FOUND")
}


