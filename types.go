package main

type Budget struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
	OwnerId int    `json:"ownerId"`
}

func NewBudget(id int, name string, balance int, ownerId int) *Budget {
	return &Budget{
		ID:      id,
		Name:    name,
		Balance: balance,
		OwnerId: ownerId,
	}
}
