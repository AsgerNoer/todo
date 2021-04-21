package models

type TodoList struct {
	Items []Item
}

type TodoListStore interface {
	GetTodoList() (TodoList, error)
	ClearTodoList() error
}
