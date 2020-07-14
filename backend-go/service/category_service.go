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

func (c *CategoryService) DeleteCategory(id int) (status int) {
	err := c.categoryService.DeleteCategoryByID(id)
	if err != nil {
		return 1
	}
	return 0
}

func (c *CategoryService) UpdateCategory(category *entity.Category) (status int) {
	err := c.categoryService.UpdateCategory(category)
	if err != nil {
		return 1
	}
	return 0
}

func (c *CategoryService) CreateCategory(category *entity.Category) (status int) {
	err := c.categoryService.CreateCategory(category)
	if err != nil {
		return 1
	}
	return 0
}