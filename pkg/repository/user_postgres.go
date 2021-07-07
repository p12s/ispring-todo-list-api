package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// UserPostgres - репозиторий
type UserPostgres struct {
	db *sqlx.DB
}

// NewUserPostgres - конструктор объекта репозитория
func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

// DeleteUser - удаление пользователя
func (r *UserPostgres) Delete(userId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, usersTable)
	_, err := r.db.Exec(query, userId)
	return err
}
