package persistence

import (
	"context"
	"fmt"
	"github.com/baransonmez/coff.app/business/core/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // for postgres
	"log"
)

type store struct {
	db *sqlx.DB
}

var schema = `CREATE TABLE IF NOT EXISTS users (
    id text  PRIMARY KEY,
    name text,
    date_created timestamp,
    date_updated timestamp
    );`

func NewStore() (store, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "coffapp")

	db, err := sqlx.Connect("postgres", psqlconn)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalln(err)
	}
	return store{db: db}, nil
}
func (s store) Create(_ context.Context, user user.User) error {
	tx := s.db.MustBegin()
	tx.MustExec("INSERT INTO users (id, name, date_created, date_updated) "+
		"VALUES ($1, $2, $3, $4)",
		user.ID, user.Name, user.DateCreated, user.DateUpdated)

	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s store) Get(id user.ID) (*user.User, error) {
	row := s.db.QueryRowx("SELECT * FROM users WHERE id=$1", id)
	var u User
	err := row.StructScan(&u)
	if err != nil {
		return nil, err
	}

	return u.ToUser(), nil
}
