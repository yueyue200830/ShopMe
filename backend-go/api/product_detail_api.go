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
	router.POST("/productDetail/image", curd.uploadDetailImage)
	router.POST("/productDetail", curd.updateDetail)
}

func (p *ProductDetailController) getDetailsByID(c *gin.Context) {
	status := 0
	var details []entity.ProductDetail
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id < 1 {
		status = 1
	} else {
		details = p.productDetailService.GetDetailsByID(id)
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

func (p *ProductDetailController) uploadDetailImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileType := file.Header.Get("Content-Type")
	status := 0
	name := ""

	name, status = p.productDetailService.GenerateRandomImageName(fileType)
	if status == 0 {
		err := c.SaveUploadedFile(file, "../images/details/" + name)
		if err != nil {
			status = 1
		}
	}

	url := ""
	if status == 0 {
		//url = "/productImage/" + name
		url = name
	}
	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"url": url,
	})
}

func (p *ProductDetailController) updateDetail(c *gin.Context) {
	status := 0
	var details []entity.ProductDetail
	if err := c.ShouldBindJSON(&details); err != nil {
		status = 1
	} else {
		status = p.productDetailService.UpdateDetail(details)
	}
	c.JSON(http.StatusOK, status)
}