package api

import (
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService *service.ProductService
}

func productApiRegister(router *gin.Engine) {
	curd := ProductController{service.GetProductService()}
	router.GET("/getAllProducts", curd.getAllProducts)
	router.GET("/getPromoteProducts", curd.getPromoteProducts)
	router.GET("/product/image/:image", curd.getProductImage)
	router.GET("/getProducts", curd.getProductsByPage)
	router.GET("/getProductNumber", curd.getProductNumber)
}

func (p *ProductController) getAllProducts(c *gin.Context) {
	products := p.productService.GetAllProducts()
	c.JSON(http.StatusOK, products)
}

func (p *ProductController) getPromoteProducts(c *gin.Context) {
	promotes := p.productService.GetPromoteProducts()
	c.JSON(http.StatusOK, promotes)
}

func (p *ProductController) getProductImage(c *gin.Context) {
	imagePath := p.productService.GetProductPath(c.Param("image"))
	c.File(imagePath)
}

func (p *ProductController) getProductsByPage(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		c.Status(http.StatusBadRequest)
		return
	}
	pageSize, err := strconv.Atoi(c.Query("size"))
	if err != nil || pageSize < 1 {
		c.Status(http.StatusBadRequest)
		return
	}
	// category default to be 0 which is for all products
	categoryID, err := strconv.Atoi(c.DefaultQuery("category", "0"))
	if err != nil || categoryID < 0 {
		c.Status(http.StatusBadRequest)
		return
	}
	products := p.productService.GetProductsByPage(page, pageSize, categoryID)
	c.JSON(http.StatusOK, products)
}

func (p *ProductController) getProductNumber(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.DefaultQuery("category", "0"))
	if err != nil || categoryID < 0 {
		c.Status(http.StatusBadRequest)
		return
	}
	number := p.productService.GetProductNumber(categoryID)
	c.JSON(http.StatusOK, number)
}

