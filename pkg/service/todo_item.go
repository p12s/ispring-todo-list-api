package service

import (
	"github.com/p12s/ispring-todo-list-api"
	"github.com/p12s/ispring-todo-list-api/pkg/repository"
)

// TodoItemService - сервис для таблицы ЭЛЕМЕНТОВ списка
type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

// NewTodoItemService - конструктор
func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

// Create - создание
func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

// GetAll - получение всех
func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

// GetById - получение по id
func (s *TodoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

// Delete - удаление
func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

// Update - обновление
func (s *TodoItemService) Update(userId, itemId int, input todo.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}

// GetAllCompletedItems - получение всех выполненных
func (s *TodoItemService) GetAllCompletedItems(userId int) ([]todo.TodoItem, error) {
	return s.repo.GetAllCompletedItems(userId)
}
