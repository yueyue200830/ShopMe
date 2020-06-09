package api

import (
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
	categoryService *service.CategoryService
}

func categoryApiRegister(router *gin.Engine) {
	curd := CategoryController{service.GetCategoryService()}
	router.GET("/categories", curd.getAllCategories)
}

func (c *CategoryController) getAllCategories(context *gin.Context) {
	categories := c.categoryService.GetAllCategories()
	context.JSON(http.StatusOK, categories)
}