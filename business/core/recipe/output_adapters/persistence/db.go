package persistence

import (
	"context"
	"fmt"
	"github.com/baransonmez/coff.app/business/core/recipe"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // for postgres
	"log"
)

type store struct {
	db *sqlx.DB
}

var schema = `CREATE TABLE IF NOT EXISTS recipe (
    id uuid  PRIMARY KEY,
    coffee_id uuid,
    user_id uuid,
    description text,
    date_created timestamp,
    date_updated timestamp
    );

CREATE TABLE IF NOT EXISTS step (
    recipe_id uuid,
    step_order integer,
    description text,
    duration_in_seconds integer,
    date_created timestamp,
    date_updated timestamp
    );
`

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

func (s store) Create(_ context.Context, recipe recipe.Recipe) error {
	tx := s.db.MustBegin()
	tx.MustExec("INSERT INTO recipe (id, coffee_id, user_id, description, date_created, date_updated) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		recipe.ID, recipe.CoffeeID, recipe.UserID, recipe.Description, recipe.DateCreated, recipe.DateUpdated)

	for i, s := range recipe.Steps {
		tx.MustExec("INSERT INTO step (recipe_id, step_order, description, duration_in_seconds, date_created, date_updated) "+
			"VALUES ($1, $2, $3, $4, $5, $6)",
			recipe.ID, i+1, s.Description, s.DurationInSeconds, recipe.DateCreated, recipe.DateUpdated)
	}

	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s store) Get(id recipe.ID) (*recipe.Recipe, error) {
	row := s.db.QueryRowx("SELECT * FROM recipe WHERE id=$1", id)
	var r Recipe
	err := row.StructScan(&r)
	if err != nil {
		return nil, err
	}

	steps := []Step{}
	err = s.db.Select(&steps, "SELECT * FROM step WHERE recipe_id=$1", id)
	if err != nil {
		return nil, err
	}
	r.Steps = steps
	return r.ToRecipe(), nil
}
