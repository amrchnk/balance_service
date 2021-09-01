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
    //fmt.Println(input.Balance, input.UserId)
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
// 	fmt.Println("result ",row)

    s:=fmt.Sprintf("%.2f", balance)
    return s,tx.Commit()
}

func (r *BalancePostgres)GetBalanceById(id int)(models.Balance,error){
    var balance models.Balance

    query:=fmt.Sprintf("SELECT user_id,balance FROM %s WHERE user_id=$1",balanceTable)
    err:=r.db.Get(&balance,query,id)
//     fmt.Println(err)
    return balance,err
}

func (r *BalancePostgres)TransferMoney(senderId,receiverId int, sum float64)([]float64,error){
    var balances [2] float64
    tx,err:=r.db.Begin()
	if err!=nil{
		return err
	}

	decQuery:=fmt.Sprintf("UPDATE balance SET balance=(SELECT balance FROM balance where user_id=$1)-$2 WHERE user_id=$1 RETURNING balance")
    row:=tx.QueryRow(decQuery,senderId,sum)
    err=row.Scan(&balance)
    if err!=nil{
        tx.Rollback()
        return err
    }

	incQuery:=fmt.Sprintf("INSERT INTO balance (user_id,balance) VALUES ($1,$2) ON CONFLICT (user_id) DO UPDATE SET balance=(select balance from balance where user_id = $1) + $2")
    _,err=tx.Exec(incQuery,receiverId,sum)
    if err!=nil{
        tx.Rollback()
        return err
    }

    return tx.Commit()
 }