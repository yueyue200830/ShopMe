package api

import (
	"backend-go/entity"
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryController struct {
	categoryService *service.CategoryService
}

func categoryApiRegister(router *gin.Engine) {
	curd := CategoryController{service.GetCategoryService()}
	router.GET("/categories", curd.getAllCategories)
	router.DELETE("/category", curd.deleteCategory)
	router.PUT("/category", curd.updateCategory)
	router.POST("/category", curd.createCategory)
}

func (c *CategoryController) getAllCategories(context *gin.Context) {
	categories := c.categoryService.GetAllCategories()
	context.JSON(http.StatusOK, categories)
}

func (c *CategoryController) deleteCategory(context *gin.Context) {
	status := 0
	id, err := strconv.Atoi(context.Query("id"))
	if err != nil {
		status = 1
	} else {
		status = c.categoryService.DeleteCategory(id)
	}

	context.JSON(http.StatusOK, status)
}

func (c *CategoryController) updateCategory(context *gin.Context) {
	status := 0
	var category *entity.Category
	if err := context.ShouldBindJSON(&category); err != nil {
		status = 1
	} else {
		status = c.categoryService.UpdateCategory(category)
	}
	context.JSON(http.StatusOK, status)
}

func (c *CategoryController) createCategory(context *gin.Context) {
	status := 0
	var category *entity.Category
	if err := context.ShouldBindJSON(&category); err != nil {
		status = 1
	} else {
		status = c.categoryService.CreateCategory(category)
	}
	context.JSON(http.StatusOK, status)
}

