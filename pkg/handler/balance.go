package handler

import (
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func (h *Handler) changeUserBalance(c *gin.Context){
    var input models.Balance

    if err:=c.BindJSON(&input);err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "status":http.StatusBadRequest,
            "error": "invalid data in body",
        })
        return
    }

    tr_type:=c.Param("type")
    if err:=ValidateType(tr_type);!err{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "status":http.StatusBadRequest,
            "error": "Unknown type of operation",
        })
        return
    }

    str,err:=h.services.Balance.ChangeUserBalance(input,tr_type)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
            "status":http.StatusInternalServerError,
            "message": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK,map[string]interface{}{
        "id":input.UserId,
        "status":http.StatusOK,
        "message":"Current balance in rubles: "+str,
    })
}

func (h *Handler) getBalanceById(c *gin.Context){
    var req models.UserBalanceQuery
    id,err:=strconv.Atoi(c.Param("id"))
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "status":http.StatusBadRequest,
            "message":"Invalid id",
        })
        return
    }
    req.UserId,req.Currency=id,c.Query("currency")
    res,err:=h.services.Balance.GetBalanceById(req)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
            "status":http.StatusInternalServerError,
            "message":err.Error(),
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
            "status":http.StatusBadRequest,
            "message": "invalid data in body",
        })
        return
    }

    balances,err:=h.services.Balance.TransferMoney(input)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
            "status":http.StatusInternalServerError,
            "message":err.Error(),
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