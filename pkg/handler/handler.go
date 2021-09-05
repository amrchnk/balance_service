package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/amrchnk/balance_service/pkg/service"
)

type Handler struct{
    services *service.Service
}

func NewHandler(services *service.Service) *Handler{
    return &Handler{services:services}
}

func (h *Handler) InitRoutes() *gin.Engine{
    router:=gin.New()

    b_api:=router.Group("/balance")
    {
        b_api.POST("/:type",h.changeUserBalance)
        b_api.GET("/:id",h.getBalanceById)
        b_api.POST("/transfer",h.transferMoney)
    }

    t_api:=router.Group("/transactions")
    {
        t_api.GET("/",h.getAllTransactions)
        t_api.GET("/:id",h.getTransactionByUserId)
    }

    return router
}