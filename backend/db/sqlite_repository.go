package db

import (
	"ShaderLand/db/models"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	Shaders *models.ShaderRepository
}

var DbRepo *Repository

const dbFileName = "sqlite.db"

func newRepository(db *sql.DB) *Repository {
	return &Repository{
		Shaders: models.NewShaderRepository(db),
	}
}

func (r *Repository) migrate() []error {
	errors := make([]error, 0)

	err1 := r.Shaders.Migrate()

	if err1 != nil {
		errors = append(errors, err1)
	}

	return errors
}

func InitDb() {
	log.Printf("Initializing Database")
	conn, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	DbRepo = newRepository(conn)
	errs := DbRepo.migrate()
	if len(errs) > 0 {
		log.Fatal(errs)
	}
	log.Printf("Done Initializing Database")
}
