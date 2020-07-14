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

func (c *CategoryRepository) DeleteCategoryByID(id int) error {
	return db.Where("id = ?", id).Delete(&entity.Category{}).Error
}

func (c *CategoryRepository) UpdateCategory(category *entity.Category) error {
	return db.Save(&category).Error
}

func (c *CategoryRepository) CreateCategory(category *entity.Category) error {
	return db.Create(&category).Error
}