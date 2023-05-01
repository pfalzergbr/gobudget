package main

type Storage interface {
	GetBudget(int) error
	CreateBudget(*Budget) error
	DeleteBudgget(int) error
	UpdateBudget(int, *Budget) error
}
