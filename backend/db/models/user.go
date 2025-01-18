package models

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"userName"`
}

type UserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UserInteface interface {
	Migrate() error
	Create(user UserRequest) (*User, error)
	All() ([]User, error)
	GetByUserName(userName string) (*User, error)
	Delete(id int64) error
	IsValidPassword(user UserRequest) bool
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

func (r *UserRepository) Create(user UserRequest) (*User, error) {
	command := "INSERT INTO users(userName, password) VALUES (?,?)"

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	resp, err := r.db.Exec(command, user.UserName, password)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}

	userResponse := &User{
		Id:       id,
		UserName: user.UserName,
	}

	return userResponse, nil
}

func (r *UserRepository) All() ([]User, error) {
	querry := "SELECT id, username FROM users"
	rows, err := r.db.Query(querry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var user User
		rows.Scan(
			&user.Id,
			&user.UserName,
		)

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetByUserName(userName string) (*User, error) {
	query := "SELECT id, userName FROM users WHERE userName = ?"
	rows, err := r.db.Query(query, userName)
	if err != nil {
		return nil, err
	}

	foundUser := rows.Next()
	if !foundUser {
		return nil, ErrNotExists
	}

	var user User
	rows.Scan(
		&user.Id,
		&user.UserName,
	)

	return &user, nil
}

func (r *UserRepository) Delete(id int64) error {
	command := "DELETE FROM users WHERE id = ?"

	_, err := r.db.Exec(command, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) IsValidPassword(user UserRequest) bool {
	query := "SELECT userName, password FROM users WHERE userName = ?"
	rows, err := r.db.Query(query, user.UserName)
	if err != nil {
		return false
	}

	isFound := rows.Next()
	if !isFound {
		return false
	}

	var foundUser UserRequest
	rows.Scan(
		&foundUser.UserName,
		&foundUser.Password,
	)

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	return err == nil
}
