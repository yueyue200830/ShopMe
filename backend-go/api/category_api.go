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
	router.GET("/allCategories", curd.getAllCategories)
	router.GET("/categories", curd.getCategories)
	router.DELETE("/category", curd.deleteCategory)
	router.PUT("/category", curd.updateCategory)
	router.POST("/category", curd.createCategory)
}

func (c *CategoryController) getAllCategories(context *gin.Context) {
	categories := c.categoryService.GetAllCategories()
	context.JSON(http.StatusOK, categories)
}

func (c *CategoryController) getCategories(context *gin.Context) {
	status := 0
	num := 0
	var categories []entity.CategoryWithNum
	page, err := strconv.Atoi(context.Query("page"))
	if err != nil || page < 1 {
		status = 1
	}
	pageSize, err := strconv.Atoi(context.Query("size"))
	if err != nil || pageSize < 1 {
		status = 1
	}
	if status == 0 {
		categories, num, status = c.categoryService.GetCategories(page, pageSize)
	}

	context.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": gin.H{
			"num": num,
			"list": categories,
		},
	})
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

