package api

import (
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
	orderService *service.OrderService
}

func orderApiRegister(router *gin.Engine) {
	curd := OrderController{service.GetOrderService()}
	router.GET("/order", curd.GetOne)
}

func (o OrderController) GetOne(c *gin.Context) {
	order := o.orderService.GetOne()
	c.JSON(http.StatusOK, order)
}