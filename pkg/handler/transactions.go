package handler

import (
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "math"
//     "fmt"
)

func (h *Handler) getAllTransactions(c *gin.Context){
    input:=models.AddressReq{
        Sort: c.Query("sort"),
        Direction: c.Query("direction"),
    }
    if input.Sort==""{
        input.Sort="amount"
    }
    if input.Direction==""{
        input.Direction="up"
    }
    page,err:=strconv.Atoi(c.Query("page"))
    records,err2:=strconv.Atoi(c.Query("records"))
    if (err!=nil||err2!=nil){
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "status":http.StatusBadRequest,
            "message": "Invalid type of data for parameters page and records",
        })
        return
    }
    input.Page,input.Records=page,records

    res,err:=h.services.Transactions.GetAllTransactions(input)
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
            "status":http.StatusInternalServerError,
            "message": err.Error(),
        })
        return
    }

    if ((int(math.Ceil(float64(len(res))/float64(input.Records)))<input.Page)||input.Page==0){
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
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

func (h *Handler) getTransactionByUserId (c *gin.Context){
    input:=models.AddressReq{
        Sort: c.Query("sort"),
        Direction: c.Query("direction"),
    }
    if input.Sort==""{
        input.Sort="amount"
    }
    if input.Direction==""{
        input.Direction="up"
    }
    page,err:=strconv.Atoi(c.Query("page"))
    records,err2:=strconv.Atoi(c.Query("records"))

    if (err!=nil||err2!=nil){
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "status":http.StatusBadRequest,
            "message": "Invalid type of data for parameters page and records",
        })
        return
    }
    input.Page,input.Records=page,records

    id,err:=strconv.Atoi(c.Param("id"))
    if err!=nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
            "status":http.StatusBadRequest,
            "message":"Invalid id",
        })
        return
    }

    res,err:=h.services.Transactions.GetTransactionByUserId(id,input)

    if err!=nil{
        c.AbortWithStatusJSON(http.StatusInternalServerError,map[string]interface{}{
            "status":http.StatusInternalServerError,
            "message": err.Error(),
        })
        return
    }

    if ((int(math.Ceil(float64(len(res))/float64(input.Records)))<input.Page)||input.Page==0){
        c.AbortWithStatusJSON(http.StatusBadRequest,map[string]interface{}{
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