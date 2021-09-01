package service

import (
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/amrchnk/balance_service/pkg/repository"
)

type BalanceService struct{
    repo *repository.Repository
}

func NewBalanceService(repo *repository.Repository)*BalanceService{
    return &BalanceService{repo:repo}
}

func (s *BalanceService)ChangeUserBalance(balance models.Balance,tr_type string)(string,error){
    return s.repo.ChangeUserBalance(balance,tr_type)
}

func (s *BalanceService)GetBalanceById(id int)(models.Balance,error){
    return s.repo.GetBalanceById(id)
}