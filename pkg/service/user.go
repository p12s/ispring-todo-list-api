package service

import (
	"github.com/p12s/ispring-todo-list-api/pkg/repository"
)

// UserService - сервис для таблицы пользователей
type UserService struct {
	userRepo     repository.UserAction
}

// NewUserService - конструктор
func NewUserService(userRepo repository.UserAction) *UserService {
	return &UserService{userRepo: userRepo}
}

// Delete - удаление
func (s *UserService) Delete(userId int) error {
	return s.userRepo.Delete(userId)
}
