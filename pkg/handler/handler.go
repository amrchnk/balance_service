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

    api:=router.Group("/balance")
    {
        api.POST("/:type",h.changeUserBalance)
        api.GET("/:id",h.getBalanceById)
        api.POST("/transfer",h.transferMoney)
    }

    return router
}