package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
)

var (
	RepoErr             = errors.New("Unable to handle Repo Request")
	ErrIdNotFound       = errors.New("Id not found")
	ErrPhonenumNotFound = errors.New("Phone num is not found")
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) (Repository, error) {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "mongodb"),
	}, nil
}

func (repo *repo) CreateCustomer(ctx context.Context, customer Customer) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO Customer(customerid, email, phone) VALUES (?, ?, ?)", customer.Customerid, customer.Email, customer.Phone)
	if err != nil {
		fmt.Println("Error occured inside CreateCustomer in repo")
		return err
	} else {
		fmt.Println("User Created:", customer.Email)
	}
	return nil
}
func (repo *repo) GetCustomerById(ctx context.Context, id string) (interface{}, error) {
	customer := Customer{}

	err := repo.db.QueryRowContext(ctx, "SELECT c.customerid,c.email,c.phone FROM Customer as c where c.customerid = ?", id).Scan(&customer.Customerid, &customer.Email, &customer.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return customer, ErrIdNotFound
		}
		return customer, err
	}
	return customer, nil
}
func (repo *repo) GetAllCustomers(ctx context.Context) (interface{}, error) {
	customer := Customer{}
	var res []interface{}
	rows, err := repo.db.QueryContext(ctx, "SELECT c.customerid,c.email,c.phone FROM Customer as c ")
	if err != nil {
		if err == sql.ErrNoRows {
			return customer, ErrIdNotFound
		}
		return customer, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&customer.Customerid, &customer.Email, &customer.Phone)
		res = append([]interface{}{customer}, res...)
	}
	return res, nil
}
func (repo *repo) DeleteCustomer(ctx context.Context, id string) (string, error) {
	res, err := repo.db.ExecContext(ctx, "DELETE FROM Customer WHERE customerid = ? ", id)
	if err != nil {
		return "", err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	} else if rowCnt == 0 {
		return "", ErrIdNotFound
	}
	return "Successfully deleted ", nil
}
func (repo *repo) UpdateCustomer(ctx context.Context, customer Customer) (string, error) {
	res, err := repo.db.ExecContext(ctx, "UPDATE Customer as c SET c.Email=? , c.Phone = ? WHERE c.customerid = ?", customer.Email, customer.Phone, customer.Customerid)
	if err != nil {
		return "", err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowCnt == 0 {
		return "", ErrIdNotFound
	}

	return "successfully updated", err
}
