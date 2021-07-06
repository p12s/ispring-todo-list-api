package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p12s/ispring-todo-list-api"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item todo.TodoItem) (int, error) {

	return 1, nil
}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem

	return items, nil
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem

	return item, nil
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {

	return nil
}

func (r *TodoItemPostgres) Update(userId, itemId int, input todo.UpdateItemInput) error {

	return nil
}