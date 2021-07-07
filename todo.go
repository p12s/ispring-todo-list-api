package todo

import "errors"

// User - пользователь
type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TodoList - список задач
type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

// UsersList - привязка списка к пользователю
type UsersList struct {
	Id     int
	UserId int
	ListId int
}

// TodoItem - элемент списка
type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

// ListItem - привязка элемента к списку
type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

// UpdateListInput - данные для обновление списка
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// Validate - валидация при обновлении списка
func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

// UpdateItemInput - данные для обновления элемента
type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

// Validate - валидация при обновлении элемента
func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
