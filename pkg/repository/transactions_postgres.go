package repository

import (
    "github.com/amrchnk/balance_service/pkg/models"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type TransactionsPostgres struct{
	db *sqlx.DB
}

func NewTransactionsPostgres(db *sqlx.DB) *BalancePostgres{
	return &TransactionsPostgres{db:db}
}

func (r *BalancePostgres) CreateTransaction(input models.Transaction)error{
    //Transaction start
	tx,err:=r.db.Begin()
	if err!=nil{
		return err
	}
    createQuery:=fmt.Sprintf("INSERT INTO %s (user_id,type_t,amount,description) VALUES ($1,$2,$3,$4)",transactionsTable)

	_,err:=tx.Exec(createQuery,input.Title,advert.Description,advert.Price)
	if err!=nil{
		tx.Rollback()
		return 0,err
	}

	return id,tx.Commit()
}