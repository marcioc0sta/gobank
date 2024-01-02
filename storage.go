package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStorage) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key, 
		first_name varchar(50),
		last_name varchar(50), 
		number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateAccount(acc *Account) error {
	q := `insert into account 
	(first_name, last_name, number, balance, created_at) 
	values ($1, $2, $3, $4, $5)`

	resp, err := s.db.Query(
		q,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
	)

	if err != nil {
		return err
	}

	fmt.Printf("resp: %v", resp)

	return nil
}

func (s *PostgresStorage) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	_, err := s.db.Query("delete from account where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) GetAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query("select * from account where id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanToAccount(rows)
	}
	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStorage) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		acc, err := scanToAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}

	return accounts, nil
}

func scanToAccount(rows *sql.Rows) (*Account, error) {
	acc := &Account{}
	err := rows.Scan(
		&acc.ID,
		&acc.FirstName,
		&acc.LastName,
		&acc.Number,
		&acc.Balance,
		&acc.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return acc, nil
}
