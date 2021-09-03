package handler

import (
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/gin-gonic/gin"
    "net/http"
//     "fmt"
    "strconv"
    "math"
)

func (h *Handler) getAllTransactions(c *gin.Context){
    input:=models.AddressReq{
        Sort: c.Query("sort"),
        Direction: c.Query("direction"),
    }
    page,err:=strconv.Atoi(c.Query("page"))
    records,err2:=strconv.Atoi(c.Query("records"))
    if (err!=nil||err2!=nil){
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
                "id":-1,
                "status":http.StatusBadRequest,
                "message": "Invalid type of data",
        })
        return
    }
    input.Page,input.Records=page,records

    ////validation

    res,err:=h.services.Transactions.GetAllTransactions(input)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
                "id":-1,
                "status":http.StatusInternalServerError,
                "message": err,
        })
        return
    }

    if (int(math.Ceil(float64(len(res))/float64(input.Records)))<input.Page){
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "id":-1,
            "status":http.StatusBadRequest,
            "message":"Page number out of range",
        })
        return
    }
    l:=(input.Page-1)*input.Records
    r:=input.Page*input.Records
    if(len(res)<r){
        res=res[l:]
    }else{
        res=res[l:r]
    }

    c.JSON(http.StatusOK,res)

}