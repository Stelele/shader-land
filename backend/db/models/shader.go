package models

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"net/url"

	"github.com/google/uuid"
	"github.com/mattn/go-sqlite3"
)

type Shader struct {
	Id           int64  `json:"id"`
	UserId       int64  `json:"userId"`
	Url          string `json:"url"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Code         string `json:"code"`
	CreationDate int    `json:"creationDate"`
}

type ShaderRequest struct {
	UserId       int64  `json:"userId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Code         string `json:"code"`
	CreationDate int    `json:"creationDate"`
}

type ShaderInterface interface {
	Migrate() error
	Create(shader ShaderRequest) (*Shader, error)
	All() ([]Shader, error)
	GetByName(name string) ([]Shader, error)
	GetByUrl(url string) (*Shader, error)
	Update(id int64, updated Shader) (*Shader, error)
	Delete(id int64) error
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
		userId INTEGER NOT NULL,
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
		INSERT INTO shaders(userId, name, description, code, creationDate, url)
			VALUES(?,?,?,?,?,?)
	`

	urlId := uuid.New()
	shaderUrl := base64.StdEncoding.EncodeToString(urlId[:])
	shaderUrl = url.PathEscape(shaderUrl[:len(shaderUrl)-2])

	res, err := r.db.Exec(command, shader.UserId, shader.Name, shader.Description, shader.Code, shader.CreationDate, shaderUrl)
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

func (r *ShaderRepository) GetByName(name string) ([]Shader, error) {
	shaders := make([]Shader, 0)
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
			url,
			name,
			description,
			code,
			creationDate
		FROM shaders
		WHERE url = ?
		ORDER BY name
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

func (r *ShaderRepository) Update(id int64, updated Shader) (*Shader, error) {
	return nil, errors.New("")
}

func (r *ShaderRepository) Delete(id int64) error {
	return errors.New("")
}
