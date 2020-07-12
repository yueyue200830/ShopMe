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
	router.POST("/productImage", curd.uploadProductImage)
	router.GET("/products", curd.getProductsByPage)
	router.GET("/productNumber", curd.getProductNumber)
	router.GET("/product/:id", curd.getProductInfo)
	router.POST("/product", curd.addProduct)
	router.PUT("/product", curd.updateProduct)
	router.DELETE("/product", curd.deleteProduct)
	router.GET("/productNames", curd.getProductNames)
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
	imagePath := p.productService.GetProductImagePath(c.Param("image"))
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

func (p *ProductController) getProductNames(c *gin.Context) {
	names := p.productService.GetProductsNames()
	status := 0
	if names == nil || len(names) == 0 {
		status = 1
	}
	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": names,
	})
}

func (p *ProductController) addProduct(c *gin.Context) {
	status := 0
	var product *entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		status = 1
	} else {
		status = p.productService.AddProduct(product)
	}
	c.JSON(http.StatusOK, status)
}

func (p *ProductController) updateProduct(c *gin.Context) {
	status := 0
	var product *entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		status = 1
	} else {
		status = p.productService.UpdateProduct(product)
	}
	c.JSON(http.StatusOK, status)
}

func (p *ProductController) deleteProduct(c *gin.Context) {
	status := 0
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		status = 1
	} else {
		status = p.productService.DeleteProduct(id)
	}

	c.JSON(http.StatusOK, status)
}

func (p *ProductController) uploadProductImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileType := file.Header.Get("Content-Type")
	status := 0
	name := ""

	name, status = p.productService.GenerateRandomImageName(fileType)
	if status == 0 {
		err := c.SaveUploadedFile(file, "../images/products/" + name)
		if err != nil {
			status = 1
		}
	}

	url := ""
	if status == 0 {
		url = "/productImage/" + name
	}
	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"url": url,
	})
}
