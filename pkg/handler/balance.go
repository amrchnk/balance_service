package handler

import (
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
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

    fmt.Println("user id: ",input.UserId)
    fmt.Println("balance: ",input.Balance)
    tr_type:=c.Param("type")
    fmt.Println("type: ",tr_type)
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
                "error": err,
        })
        return
    }

    c.JSON(http.StatusOK,map[string]interface{}{
        "id":input.UserId,
        "status":http.StatusOK,
        "message":"Current balance in rubles: "+str,
    })
}
