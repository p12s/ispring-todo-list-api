package service

import (
	"github.com/p12s/ispring-todo-list-api"
	"github.com/p12s/ispring-todo-list-api/pkg/repository"
)

// TodoListService - сервис для таблицы СПИСКОВ
type TodoListService struct {
	repo repository.TodoList
}

// NewTodoListService - конструктор
func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

// Create - создание
func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

// GetAll - получение всех
func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

// GetById - получение по id
func (s *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

// Delete - удаление
func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

// Update - обновление
func (s *TodoListService) Update(userId, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
