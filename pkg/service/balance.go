package service

import (
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/amrchnk/balance_service/pkg/repository"
    "fmt"
)

type BalanceService struct{
    repo *repository.Repository
}

func NewBalanceService(repo *repository.Repository)*BalanceService{
    return &BalanceService{repo:repo}
}

func (s *BalanceService)ChangeUserBalance(balance models.Balance,tr_type string)(string,error){
    t:=models.Transaction{
        UserId:balance.UserId,
        Type: tr_type+" in balance",
        Amount: balance.Balance,
    }

    res,err:=s.repo.ChangeUserBalance(balance,tr_type)
    if err!=nil{
        return s.repo.ChangeUserBalance(balance,tr_type)
    }

    if str,er:=s.repo.CreateTransaction(t);er!=nil{
        return str,er
    }
    return res,err
}

func (s *BalanceService)GetBalanceById(id int)(models.Balance,error){
    return s.repo.GetBalanceById(id)
}

func (s *BalanceService)TransferMoney(senderId,receiverId int, sum float64)([]float64,error){
    var mas []float64
    from:=models.Transaction{
        UserId:senderId,
        Type: "outgoing transfer",
        Amount: sum,
        Description: fmt.Sprint("money transfer to user with id=",receiverId),
    }
    to:=models.Transaction{
        UserId:receiverId,
        Type: "incoming transfer",
        Amount: sum,
        Description: fmt.Sprint("money transfer from user with id=",senderId),
    }
    res,err:=s.repo.TransferMoney(senderId,receiverId,sum)

    if err!=nil{
        return s.repo.TransferMoney(senderId,receiverId,sum)
    }

    if _,er:=s.repo.CreateTransaction(to);er!=nil{
        return mas,er
    }

    if _,er:=s.repo.CreateTransaction(from);er!=nil{
        return mas,er
    }
    return res,err
}