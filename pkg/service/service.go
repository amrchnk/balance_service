package service

import (
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/amrchnk/balance_service/pkg/repository"
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=mocks/mock.go

type Balance interface{
    ChangeUserBalance(balance models.Balance,tr_type string)(string,error)
    GetBalanceById(input models.UserBalanceQuery)(models.UserBalanceResponse,error)
    TransferMoney(input models.TransferQuery)([]float64,error)
}

type Transactions interface{
    GetAllTransactions(input models.AddressReq)([]models.Transaction,error)
    GetTransactionByUserId(id int,input models.AddressReq)([]models.Transaction,error)
}

type Service struct{
    Balance
    Transactions
}

func NewService(repos *repository.Repository) *Service{
    return &Service{
        Balance: NewBalanceService(repos),
        Transactions: NewTransactionsService(repos),
    }
}