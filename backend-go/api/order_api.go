package api

import (
	"backend-go/entity"
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderService *service.OrderService
}

func orderApiRegister(router *gin.Engine) {
	curd := OrderController{service.GetOrderService()}
	router.GET("/order", curd.getOrder)
	router.GET("/orderNumber", curd.getUserOrderNumber)
	router.GET("/orders", curd.getUserOrdersByPage)
	router.POST("/cancelOrder", curd.cancelOrder)
	router.POST("/payOrder", curd.payOrder)
	router.POST("/order", curd.newOrder)
	router.POST("/order/finish", curd.finishOrder)
}

func (o *OrderController) getOrder(c *gin.Context) {
	status := 0
	var order *entity.DetailOrderWithProducts
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id < 1 {
		status = 1
	} else {
		order = o.orderService.GetOrder(id)
		if order == nil || order.ID == 0 {
			status = 1
		}
	}

	c.JSON(http.StatusOK, gin.H {
		"code": status,
		"data": order,
	})
}

func (o *OrderController) getUserOrderNumber(c *gin.Context) {
	number := 0
	userID, err := strconv.Atoi(c.Query("userID"))
	if err == nil && userID > 0 {
		number = o.orderService.GetUserOrderNumber(userID)
	}
	c.JSON(http.StatusOK, number)
}

func (o *OrderController) getUserOrdersByPage(c *gin.Context) {
	status := 0
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		status = 1
	}
	pageSize, err := strconv.Atoi(c.Query("size"))
	if err != nil || pageSize < 1 {
		status = 1
	}
	userID, err := strconv.Atoi(c.Query("userID"))
	if err != nil || userID <= 0 {
		status = 1
	}

	var orders []entity.SimpleOrderWithProducts
	if status == 0 {
		orders = o.orderService.GetUserOrdersByPage(page, pageSize, userID)
	}
	if orders == nil || len(orders) == 0 {
		status = 1
	}

	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": orders,
	})
}

func (o *OrderController) cancelOrder(c *gin.Context) {
	var order *entity.Order
	status := 0
	if err := c.ShouldBindJSON(&order); err != nil {
		status = 1
	} else {
		status = o.orderService.CancelOrder(order)
	}
	c.JSON(http.StatusOK, status)
}

func (o *OrderController) payOrder(c *gin.Context) {
	var order *entity.Order
	status := 0
	if err := c.ShouldBindJSON(&order); err != nil {
		status = 1
	} else {
		status = o.orderService.PayOrder(order)
	}
	c.JSON(http.StatusOK, status)
}

func (o *OrderController) newOrder(c *gin.Context) {
	var orderProducts *entity.DetailOrderWithProducts
	status := 0
	var id int
	if err := c.ShouldBindJSON(&orderProducts); err != nil {
		status = 1
	} else {
		status, id = o.orderService.NewOrder(orderProducts)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": id,
	})
}

func (o *OrderController) finishOrder(c *gin.Context) {
	var order *entity.Order
	status := 0
	if err := c.ShouldBindJSON(&order); err != nil {
		status = 1
	} else {
		status = o.orderService.FinishOrder(order)
	}
	c.JSON(http.StatusOK, status)
}