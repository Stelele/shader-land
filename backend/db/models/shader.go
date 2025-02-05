package models

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/mattn/go-sqlite3"
)

type Shader struct {
	Id           int64  `json:"id"`
	UserId       string `json:"userId"`
	UserName     string `json:"userName"`
	Url          string `json:"url"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Code         string `json:"code"`
	CreationDate int    `json:"creationDate"`
}

type ShaderRequest struct {
	UserId       string `json:"userId"`
	UserName     string `json:"userName"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Code         string `json:"code"`
	CreationDate int    `json:"creationDate"`
}

type ShaderRepository struct {
	db *sql.DB
}

func NewShaderRepository(db *sql.DB) *ShaderRepository {
	return &ShaderRepository{db: db}
}

func (r *ShaderRepository) Migrate() error {
	command := `
	CREATE TABLE IF NOT EXISTS shaders(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId TEXT NOT NULL,
		userName TEXT NOT NULL,
		url TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		description TEXT,
		code TEXT NOT NULL,
		creationDate INTEGER,
		FOREIGN KEY (userId)
			REFERENCES users(id)
				ON UPDATE CASCADE
       			ON DELETE CASCADE
	);
	`

	_, err := r.db.Exec(command)
	return err
}

func (r *ShaderRepository) Create(shader ShaderRequest) (*Shader, error) {
	command := `
		INSERT INTO shaders(userId, userName, name, description, code, creationDate, url)
			VALUES(?,?,?,?,?,?,?)
	`

	urlId := uuid.New()
	shaderUrl := base64.StdEncoding.EncodeToString(urlId[:])
	shaderUrl = url.PathEscape(shaderUrl[:len(shaderUrl)-2])

	res, err := r.db.Exec(
		command,
		shader.UserId,
		shader.UserName,
		shader.Name,
		shader.Description,
		shader.Code,
		shader.CreationDate,
		shaderUrl)

	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}

		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	shaderResponse := Shader{
		Id:           id,
		UserId:       shader.UserId,
		UserName:     shader.UserName,
		Url:          shaderUrl,
		Name:         shader.Name,
		Description:  shader.Description,
		Code:         shader.Code,
		CreationDate: shader.CreationDate,
	}
	return &shaderResponse, nil
}

func (r *ShaderRepository) All() ([]Shader, error) {
	shaders := make([]Shader, 0)
	query := `
		SELECT
			id,
			userId,
			userName,
			url,
			name,
			description,
			code,
			creationDate
		FROM shaders
		ORDER BY LOWER(name)
	`

	rows, err := r.db.Query(query)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr, sqlite3.ErrNotFound) {
				return shaders, nil
			}
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var shader Shader
		err := rows.Scan(
			&shader.Id,
			&shader.UserId,
			&shader.UserName,
			&shader.Url,
			&shader.Name,
			&shader.Description,
			&shader.Code,
			&shader.CreationDate,
		)
		if err != nil {
			return nil, err
		}

		shaders = append(shaders, shader)
	}

	return shaders, nil
}

func (r *ShaderRepository) GetById(id int64) (*Shader, error) {
	query := `
		SELECT
			id,
			userId,
			url,
			name,
			description,
			code,
			creationDate
		FROM shaders
		WHERE id = ?
	`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	found := rows.Next()
	if !found {
		return nil, ErrNotExists
	}

	var shader Shader
	err2 := rows.Scan(
		&shader.Id,
		&shader.UserId,
		&shader.Url,
		&shader.Name,
		&shader.Description,
		&shader.Code,
		&shader.CreationDate,
	)
	if err2 != nil {
		return nil, err
	}

	return &shader, nil
}

func (r *ShaderRepository) GetByName(name string) ([]Shader, error) {
	shaders := make([]Shader, 0)
	query := `
		SELECT
			id,
			userId,
			userName,
			url,
			name,
			description,
			code,
			creationDate
		FROM shaders
		WHERE name LIKE ?
		ORDER BY name
	`

	rows, err := r.db.Query(query, "%"+name+"%")
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr, sqlite3.ErrNotFound) {
				return shaders, nil
			}
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var shader Shader
		err := rows.Scan(
			&shader.Id,
			&shader.UserId,
			&shader.UserName,
			&shader.Url,
			&shader.Name,
			&shader.Description,
			&shader.Code,
			&shader.CreationDate,
		)
		if err != nil {
			return nil, err
		}

		shaders = append(shaders, shader)
	}

	return shaders, nil
}

func (r *ShaderRepository) GetByUrl(url string) (*Shader, error) {
	query := `
		SELECT
			id,
			userId,
			userName,
			url,
			name,
			description,
			code,
			creationDate
		FROM shaders
		WHERE url = ?
	`

	rows, err := r.db.Query(query, url)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	found := rows.Next()
	if !found {
		return nil, ErrNotExists
	}

	var shader Shader
	err2 := rows.Scan(
		&shader.Id,
		&shader.UserId,
		&shader.UserId,
		&shader.Url,
		&shader.Name,
		&shader.Description,
		&shader.Code,
		&shader.CreationDate,
	)
	if err2 != nil {
		return nil, err
	}

	return &shader, nil
}

func (r *ShaderRepository) Update(id int64, updated ShaderRequest) (*Shader, error) {
	command := "UPDATE shaders SET "
	updates := make([]any, 0, 4)
	cols := make([]string, 0, 4)

	if updated.Code != "" {
		cols = append(cols, "code")
		updates = append(updates, updated.Code)
	}
	if updated.CreationDate > 0 {
		cols = append(cols, "creationDate")
		updates = append(updates, updated.CreationDate)
	}
	if updated.Description != "" {
		cols = append(cols, "description")
		updates = append(updates, updated.Description)
	}
	if updated.Name != "" {
		cols = append(cols, "name")
		updates = append(updates, updated.Name)
	}

	if len(updates) == 0 {
		resp, err := r.GetById(id)
		return resp, err
	}

	for i := 0; i < len(updates); i++ {
		if i == 0 {
			command += fmt.Sprintf("%s = ?", cols[i])
			continue
		}
		command += fmt.Sprintf(",%s = ? ", cols[i])
	}
	command += "WHERE id = ?"
	updates = append(updates, id)

	res, err := r.db.Exec(command, updates...)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	shader, _ := r.GetById(id)
	return shader, nil
}

func (r *ShaderRepository) Delete(id int64) error {
	command := `DELETE FROM shaders WHERE id = ?`

	_, err := r.db.Exec(command, id)
	if err != nil {
		return err
	}

	return nil
}
