package api

import (
	"backend-go/entity"
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CartController struct {
	cartService *service.CartService
}

func cartApiRegister(router *gin.Engine) {
	curd := CartController{service.GetCartService()}
	router.GET("/cardProducts", curd.getCardProducts)
	router.POST("/cardProduct", curd.addProduct)
	router.PUT("/cardProduct", curd.modifyProductNumber)
	router.DELETE("/cardProduct", curd.deleteProduct)
}

func (c *CartController) getCardProducts(context *gin.Context) {
	status := 0
	var products []entity.CartProduct
	id, err := strconv.Atoi(context.Query("id"))
	if err != nil || id < 0 {
		status = 1
	} else {
		products = c.cartService.GetCartProducts(id)
	}
	context.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": products,
	})
}

func (c *CartController) addProduct(context *gin.Context) {
	status := 0
	var cart *entity.Cart
	if err := context.ShouldBindJSON(&cart); err != nil {
		status = 1
	} else {
		status = c.cartService.AddProduct(cart)
	}
	context.JSON(http.StatusOK, status)
}

func (c *CartController) modifyProductNumber(context *gin.Context) {
	status := 0
	var cart *entity.Cart
	if err := context.ShouldBindJSON(&cart); err != nil {
		status = 1
	} else {
		status = c.cartService.ModifyProduct(cart)
	}
	context.JSON(http.StatusOK, status)
}

func (c *CartController) deleteProduct(context *gin.Context) {
	status := 0
	productID, err := strconv.Atoi(context.Query("productID"))
	if err != nil || productID < 1 {
		status = 1
	}
	userID, err := strconv.Atoi(context.Query("userID"))
	if err != nil || userID < 1 {
		status = 1
	}
	if status == 0 {
		status = c.cartService.DeleteProduct(productID, userID)
	}

	context.JSON(http.StatusOK, status)
}
