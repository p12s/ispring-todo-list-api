package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p12s/ispring-todo-list-api"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	return 1, nil
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	return lists, nil
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	return list, nil
}

func (r *TodoListPostgres) Delete(userId, listId int) error {

	return nil
}

func (r *TodoListPostgres) Update(userId, listId int, input todo.UpdateListInput) error {

	return nil
}
