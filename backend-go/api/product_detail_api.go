package api

import (
	"backend-go/entity"
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductDetailController struct {
	productDetailService *service.ProductDetailService
}

func productDetailApiRegister(router *gin.Engine) {
	curd := ProductDetailController{service.GetProductDetailService()}
	router.GET("/productDetails", curd.getDetailsByID)
	router.GET("/productDetail/:name", curd.getDetailByName)
}

func (p *ProductDetailController) getDetailsByID(c *gin.Context) {
	status := 0
	var details []entity.ProductDetail
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id < 1 {
		status = 1
	} else {
		details = p.productDetailService.GetDetailsByID(id)
		if len(details) == 0 {
			status = 1
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": details,
	})
}

func (p *ProductDetailController) getDetailByName(c *gin.Context) {
	path := p.productDetailService.GetDetailPath(c.Param("name"))
	c.File(path)
}