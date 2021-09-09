package service

import (
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/amrchnk/balance_service/pkg/repository"
    "fmt"
    "net/http"
    "os"
    "encoding/json"
    "io/ioutil"
    "errors"
    "math"
)

const defaultCur="RUB"

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

func (s *BalanceService)GetBalanceById(input models.UserBalanceQuery)(models.UserBalanceResponse,error){
    var res models.UserBalanceResponse
    balance,err:=s.repo.GetBalanceById(input.UserId)
    if err!=nil{
        return res,err
    }
    res.UserId,res.Balance=input.UserId,balance
    if (input.Currency!=""){
        var data models.MessageAPI
        url:=fmt.Sprintf("http://api.exchangeratesapi.io/v1/latest?access_key=%s&format=1",os.Getenv("API_KEY"))
        resp, err := http.Get(url)
        if err != nil {
            return res,err
        }
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return res,err
        }
        err = json.Unmarshal(body, &data)
        if err != nil {
            return res,err
        }
        if _,ex:=data.Rates[input.Currency];!ex{
            return res,errors.New("Incorrect currency")
        }
        res.Currency=input.Currency
        res.Balance=math.Round(balance/(data.Rates[defaultCur]/data.Rates[input.Currency])*100)/100
        return res,nil
    }
    res.Currency="RUB"
    return res,nil
}

func (s *BalanceService)TransferMoney(input models.TransferQuery)([]float64,error){
    var mas []float64
    from:=models.Transaction{
        UserId:input.SenderId,
        Type: "outgoing transfer",
        Amount: input.Sum,
        Description: fmt.Sprint("money transfer to user with id=",input.ReceiverId),
    }
    to:=models.Transaction{
        UserId: input.ReceiverId,
        Type: "incoming transfer",
        Amount: input.Sum,
        Description: fmt.Sprint("money transfer from user with id=",input.SenderId),
    }
    res,err:=s.repo.TransferMoney(input)

    if err!=nil{
        return s.repo.TransferMoney(input)
    }

    if _,er:=s.repo.CreateTransaction(to);er!=nil{
        return mas,er
    }

    if _,er:=s.repo.CreateTransaction(from);er!=nil{
        return mas,er
    }
    return res,err
}