package handler

import (
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/gin-gonic/gin"
    "net/http"
//     "fmt"
    "strconv"
)

func (h *Handler) changeUserBalance(c *gin.Context){
    var input models.Balance
    //var transact models.Transaction

    if err:=c.BindJSON(&input);err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
                "id":-1,
                "status":http.StatusBadRequest,
                "error": "error in input",
        })
        return
    }

    tr_type:=c.Param("type")
    if err:=ValidateType(tr_type);!err{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "id":-1,
            "status":http.StatusBadRequest,
            "error": "Unknown type of transaction",
        })
        return
    }

    str,err:=h.services.Balance.ChangeUserBalance(input,tr_type)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
                "id":-1,
                "status":http.StatusInternalServerError,
                "message": err,
        })
        return
    }

    c.JSON(http.StatusOK,map[string]interface{}{
        "id":input.UserId,
        "status":http.StatusOK,
        "message":"Current balance in rubles: "+str,
    })
}

// func (h *Handler) getBalanceById(c *gin.Context){
//     var balance models.Balance
//     id,err:=strconv.Atoi(c.Param("id"))
//     if err!=nil{
//         c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
//             "id":-1,
//             "status":http.StatusBadRequest,
//             "message":"Invalid id",
//         })
//         return
//     }
//
//     balance,err=h.services.Balance.GetBalanceById(id)
//     if err!=nil{
//         c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
//             "id":-1,
//             "status":http.StatusInternalServerError,
//             "message":err,
//         })
//         return
//     }
//     if currency=="USD"{
//
//     }
//     c.JSON(http.StatusOK,map[string]interface{}{
//         "id":balance.UserId,
//         "balance (rub)":balance.Balance,
//     })
// }

func (h *Handler) getBalanceById(c *gin.Context){
    var req models.UserBalanceQuery
    id,err:=strconv.Atoi(c.Param("id"))
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "id":-1,
            "status":http.StatusBadRequest,
            "message":"Invalid id",
        })
        return
    }
    req.UserId,req.Currency=id,c.Query("currency")
//     fmt.Printf(req.Currency)
    res,err:=h.services.Balance.GetBalanceById(req)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
            "id":-1,
            "status":http.StatusInternalServerError,
            "message":err,
        })
        return
    }
    c.JSON(http.StatusOK,res)
}

func (h *Handler) transferMoney(c *gin.Context){
    var input models.TransferQuery
    var balances [] float64
    if err:=c.BindJSON(&input);err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
                "id":-1,
                "status":http.StatusBadRequest,
                "error": "error in input",
        })
        return
    }

    balances,err:=h.services.Balance.TransferMoney(input.SenderId,input.ReceiverId,input.Sum)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
            "id":-1,
            "status":http.StatusInternalServerError,
            "message":err,
        })
        return
    }
    res:=models.TransferResponse{
        SenderId: input.SenderId,
        SenderSum: balances[0],
        ReceiverId: input.ReceiverId,
        ReceiverSum: balances[1],
    }
    c.JSON(http.StatusOK,res)
}