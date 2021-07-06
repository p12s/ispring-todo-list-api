package todo

import "errors"

type User struct {
	Id 			int 	`json:"-" db:"id"`
	Name 		string 	`json:"name" binding:"required"`
	Username 	string 	`json:"username" binding:"required"`
	Password 	string 	`json:"password" binding:"required"`
}

type TodoList struct {
	Id int `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id int
	UserId int
	ListId int
}
type TodoItem struct {
	Id int `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Done bool `json:"done" db:"done"`
}

type ListItem struct {
	Id int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}