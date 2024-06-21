package main

type TodoServiceInterface interface {
	GetAllTodos() ([]Todo, error)
	GetTodo(id string) (Todo, error)
	CreateTodo(todo Todo) (Todo, error)
	UpdateTodo(id string, todo Todo) (Todo, error)
	DeleteTodo(id string) error
}

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
