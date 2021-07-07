package service

import (
	"github.com/p12s/ispring-todo-list-api"
	"github.com/p12s/ispring-todo-list-api/pkg/repository"
)

// Authorization - интерфейс для регистрации/авторизации
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

// TodoList - интерфейст для работы со списками задач
type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

// TodoItem - интейрфейс для работы с элементами списка задач
type TodoItem interface {
	Create(userId, listId int, list todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
	GetAllCompletedItems(userId int) ([]todo.TodoItem, error)
}

// UserAction - интерфейс для удаления пользователей
type UserAction interface {
	Delete(userId int) error
}

// Service - сервис
type Service struct {
	Authorization
	TodoItem
	TodoList
	UserAction
}

// NewService - конструктор
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
		UserAction:      NewUserService(repos.UserAction),
	}
}
