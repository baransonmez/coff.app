package persistence

import (
	"context"
	"fmt"
	"github.com/baransonmez/coff.app/business/core/coffee"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // for postgres
	"log"
)

type store struct {
	db *sqlx.DB
}

var schema = `CREATE TABLE IF NOT EXISTS coffee_bean (
    id uuid  PRIMARY KEY,
    name text,
    roaster text,
    origin text,
    price integer,
    roast_date timestamp,
    date_created timestamp,
    date_updated timestamp
    );`

// execute a query on the server
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

func (s store) Create(_ context.Context, bean coffee.Bean) error {
	tx := s.db.MustBegin()
	tx.MustExec("INSERT INTO coffee_bean (id, name, roaster, origin, price, roast_date, date_created, date_updated) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		bean.ID, bean.Name, bean.Roaster, bean.Origin, bean.Price, bean.RoastDate, bean.DateCreated, bean.DateUpdated)

	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s store) Get(id coffee.ID) (*coffee.Bean, error) {
	row := s.db.QueryRowx("SELECT * FROM coffee_bean WHERE id=$1", id)
	var bean Bean
	err := row.StructScan(&bean)
	if err != nil {
		return nil, err
	}
	return bean.ToBean(), nil
}
