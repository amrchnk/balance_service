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

    if err:=c.BindJSON(&input);err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
                "id":-1,
                "status":http.StatusBadRequest,
                "error": "error in input",
        })
        return
    }

    tr_type:=c.Param("type")
//     fmt.Println("type: ",tr_type)
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

func (h *Handler) getBalanceById(c *gin.Context){
    var balance models.Balance
    id,err:=strconv.Atoi(c.Param("id"))
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "id":-1,
            "status":http.StatusBadRequest,
            "message":"Invalid id",
        })
        return
    }

    balance,err=h.services.Balance.GetBalanceById(id)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
            "id":-1,
            "status":http.StatusInternalServerError,
            "message":err,
        })
        return
    }

    c.JSON(http.StatusOK,map[string]interface{}{
        "id":balance.UserId,
        "balance (rub)":balance.Balance,
    })
}
