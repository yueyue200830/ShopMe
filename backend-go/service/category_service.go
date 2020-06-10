package service

import (
	"backend-go/dao"
	"backend-go/entity"
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

func (c *CategoryService) GetAllCategories() []entity.Category {
	return c.categoryService.GetAllCategories()
}

