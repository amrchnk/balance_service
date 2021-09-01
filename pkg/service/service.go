package service

import (
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/amrchnk/balance_service/pkg/repository"
)

type Balance interface{
    ChangeUserBalance(balance models.Balance,tr_type string)(string,error)
    GetBalanceById(id int)(models.Balance,error)
}

type Service struct{
    Balance
}

func NewService(repos *repository.Repository) *Service{
    return &Service{
        Balance: NewBalanceService(repos),
    }
}