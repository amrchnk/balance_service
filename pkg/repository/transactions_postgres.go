package repository

import (
    "github.com/amrchnk/balance_service/pkg/models"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type TransactionsPostgres struct{
	db *sqlx.DB
}

func NewTransactionsPostgres(db *sqlx.DB) *TransactionsPostgres{
	return &TransactionsPostgres{db:db}
}

func (r *TransactionsPostgres) CreateTransaction(input models.Transaction)(string,error){
    //Transaction start
	tx,err:=r.db.Begin()
	if err!=nil{
		return "problem",err
	}
    createQuery:=fmt.Sprintf("INSERT INTO %s (user_id,type_t,amount,description) VALUES ($1,$2,$3,$4)",transactionsTable)

	_,err=tx.Exec(createQuery,input.UserId,input.Type,input.Amount,input.Description)
	if err!=nil{
		tx.Rollback()
		return "problem",err
	}
	return "Ok",tx.Commit()
}

func (r *TransactionsPostgres) GetAllTransactions(input models.AddressReq)([]models.Transaction,error){
    var tr_mas []models.Transaction
    s:=""
    switch {
        case input.Direction=="up":
            s+=" ASC"
        default:
            s+=" DESC"
    }
    query:=fmt.Sprintf("SELECT user_id,type_t,amount,description,created FROM %s ORDER BY $1"+s, transactionsTable)
    fmt.Println(query)
    fmt.Println(input.Sort)
    if err:=r.db.Select(&tr_mas,query,input.Sort); err!=nil{
        fmt.Println(err)
        return nil,err
    }

    return tr_mas,nil
}