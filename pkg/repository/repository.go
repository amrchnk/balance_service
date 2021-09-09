package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/amrchnk/balance_service/pkg/models"
)

type Balance interface{
	ChangeUserBalance(input models.Balance, tr_type string)(string,error)
	GetBalanceById(id int)(float64,error)
	TransferMoney(input models.TransferQuery)([]float64,error)
}

type Transactions interface{
    CreateTransaction(input models.Transaction)(string,error)
    GetAllTransactions(input models.AddressReq)([]models.Transaction,error)
    GetTransactionByUserId(id int,input models.AddressReq)([]models.Transaction,error)
}

type Repository struct {
	Balance
	Transactions
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Balance: NewBalancePostgres(db),
		Transactions: NewTransactionsPostgres(db),
	}
}
