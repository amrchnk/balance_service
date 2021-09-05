package service

import (
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/amrchnk/balance_service/pkg/repository"
//     "fmt"
)

type TransactionsService struct{
    repo *repository.Repository
}

func NewTransactionsService(repo *repository.Repository)*TransactionsService{
    return &TransactionsService{repo:repo}
}

func (s *TransactionsService) GetAllTransactions(input models.AddressReq)([]models.Transaction,error){
    return s.repo.GetAllTransactions(input)
}

func (s *TransactionsService) GetTransactionByUserId(id int,input models.AddressReq)([]models.Transaction,error){
    return s.repo.GetTransactionByUserId(id,input)
}