package handler

import (
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/gin-gonic/gin"
    "net/http"
//     "fmt"
    "strconv"
)

func (h *Handler) getAllTransactions(c *gin.Context){
    page:=c.Get("page")
    sort:=c.Get("sort")
}