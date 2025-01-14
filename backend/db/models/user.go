package models

import (
	"database/sql"
	"errors"
)

type User struct {
	Id       int64
	UserName string
}

type UserInteface interface {
	Migrate() error
	Create(user User) (*User, error)
	All() ([]User, error)
	GetByName(name string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetById(id int64) (*User, error)
	Update(id int64, updated User) (*User, error)
	Delete(id int64) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Migrate() error {
	query := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			userName TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`

	_, err := r.db.Exec(query)
	return err
}

func (r *UserRepository) Create(user User) (*User, error) {
	return nil, errors.New("Not implemented")
}

func (r *UserRepository) All() ([]User, error) {
	return nil, errors.New("Not implemented")
}

func (r *UserRepository) GetByName(name string) (*User, error) {
	return nil, errors.New("Not implemented")
}

func (r *UserRepository) GetByEmail(email string) (*User, error) {
	return nil, errors.New("Not implemented")
}

func (r *UserRepository) GetById(id int64) (*User, error) {
	return nil, errors.New("Not implemented")
}

func (r *UserRepository) Update(id int64, updated User) (*User, error) {
	return nil, errors.New("Not implemented")
}

func (r *UserRepository) Delete(id int64) error {
	return errors.New("Not implemented")
}
