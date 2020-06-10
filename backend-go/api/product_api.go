package api

import (
	"backend-go/entity"
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
	router.GET("/allProducts", curd.getAllProducts)
	router.GET("/promoteProducts", curd.getPromoteProducts)
	router.GET("/productImage/:image", curd.getProductImage)
	router.GET("/products", curd.getProductsByPage)
	router.GET("/productNumber", curd.getProductNumber)
	router.GET("/product/:id", curd.getProductInfo)
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
	status := 0
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		status = 1
	}
	pageSize, err := strconv.Atoi(c.Query("size"))
	if err != nil || pageSize < 1 {
		status = 1
	}
	// category default to be 0 which is for all products
	categoryID, err := strconv.Atoi(c.DefaultQuery("category", "0"))
	if err != nil || categoryID < 0 {
		status = 1
	}
	var products []entity.Product
	if status == 0 {
		products = p.productService.GetProductsByPage(page, pageSize, categoryID)
	}
	if products == nil || len(products) == 0 {
		status = 1
	}

	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": products,
	})
}

func (p *ProductController) getProductNumber(c *gin.Context) {
	number := 0
	categoryID, err := strconv.Atoi(c.DefaultQuery("category", "0"))
	if err == nil && categoryID >= 0 {
		number = p.productService.GetProductNumber(categoryID)
	}

	c.JSON(http.StatusOK, number)
}

// Get one product by id
func (p *ProductController) getProductInfo(c *gin.Context) {
	status := 0
	var product entity.Product
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		status = 1
	} else {
		product = p.productService.GetProductInfo(id)
		if product.ID == 0 {
			status = 1
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": product,
	})
}

