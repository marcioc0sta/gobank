package main

import (
	"math/rand"
	"time"
)

type CreateAccountReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    int64     `json:"number"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(10000)),
		CreatedAt: time.Now().UTC(),
	}
}

type TransferRequest struct {
	ToAccountID int     `json:"to_account"`
	Amount      float64 `json:"amount"`
}
