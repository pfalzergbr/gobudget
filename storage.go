package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface {
	GetBudget(id int) error
	CreateBudget(budget *Budget) error
	DeleteBudgget(id int) error
	UpdateBudget(id int, budget *Budget) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStorage, error) {
	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/gobudget?sslmode=disable")

	if err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}
