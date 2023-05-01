package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface {
	GetBudget(id int) error
	CreateBudget(budget *Budget) error
	DeleteBudget(id int) error
	UpdateBudget(id int, budget *Budget) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStorage, error) {
	db, err := sql.Open("postgres", "postgres://postgres:1234@localhost:5432/?sslmode=disable")

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) Init() error {
	// return s.db.Close()
	s.createBudgetTable()

	return nil
}

func (s *PostgresStorage) createBudgetTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS budgets (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			balance INT NOT NULL,
			owner_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		);
	`)

	return err
}

func (s *PostgresStorage) GetBudget(id int) (*Budget, error) {
	return nil, nil
}

func (s *PostgresStorage) CreateBudget(budget *Budget) (*Budget, error) {
	return nil, nil
}

func (s *PostgresStorage) DeleteBudget(id int) error {
	return nil
}

func (s *PostgresStorage) UpdateBudget(id int, budget *Budget) error {
	return nil
}
