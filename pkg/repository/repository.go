package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/amrchnk/balance_service/pkg/models"
)

type Balance interface{
	ChangeUserBalance(input models.Balance, tr_type string)(string,error)
}

type Repository struct {
	Balance
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Balance: NewBalancePostgres(db),
	}
}
