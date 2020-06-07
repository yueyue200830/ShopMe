package service

import (
	"backend-go/dao"
)

var categoryService *CategoryService

type CategoryService struct {
	categoryService *dao.CategoryRepository
}

func init() {
	categoryService = &CategoryService{dao.GetCategoryRepository()}
}

func GetCategoryService() *CategoryService {
	return categoryService
}

func (c *CategoryService) GetAllCategories() []dao.Category {
	return c.categoryService.GetAllCategories()
}

