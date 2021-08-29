package repository

import (
    "github.com/amrchnk/balance_service/pkg/models"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type BalancePostgres struct{
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres{
	return &BalancePostgres{db:db}
}

func (r *BalancePostgres) ChangeUserBalance(input models.Balance, tr_type string)(string,error){
    var balance float64
    fmt.Println(input.Balance, input.UserId)
    //Transaction start
	tx,err:=r.db.Begin()
	if err!=nil{
	    s:=fmt.Sprintf("%f", balance)
		return s,err
	}
	//СДЕЛАТЬ ПРОВЕРКУ tr_type....
	Query:=""
    if tr_type=="increase"{
        Query=fmt.Sprint("INSERT INTO balance (user_id,balance) VALUES ($1,$2) ON CONFLICT (user_id) DO UPDATE SET balance=(select balance from balance where user_id = $1) + $2 RETURNING balance")
    }else if tr_type=="decrease"{
        Query=fmt.Sprint("INSERT INTO balance (user_id,balance) VALUES ($1,$2) ON CONFLICT (user_id) DO UPDATE SET balance=(select balance from balance where user_id = $1) - $2 RETURNING balance")
    }

    row:=tx.QueryRow(Query,input.UserId,input.Balance)
	err=row.Scan(&balance)

	if err!=nil{
	    tx.Rollback()
	    s:=fmt.Sprintf("%f", balance)
        return s,err
	}
    s:=fmt.Sprintf("%f", balance)
    return s,nil
}