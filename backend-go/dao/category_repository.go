package dao

import "backend-go/entity"

var categoryRepository *CategoryRepository

func init() {
	categoryRepository = &CategoryRepository{}
}

type CategoryRepository struct {
}

func GetCategoryRepository() *CategoryRepository {
	return categoryRepository
}

func (c *CategoryRepository) GetAllCategories() []entity.Category {
	var categories []entity.Category
	db.Find(&categories)
	return categories
}
