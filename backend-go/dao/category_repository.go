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
	db.Order("id").Find(&categories)
	return categories
}

func (c *CategoryRepository) GetCategoriesByPage(page, pageSize int) (categories []entity.CategoryWithNum, err error) {
	offsetNum := (page - 1) * pageSize
	joined := db.Raw("select * from categories left join (select category_id, count(*) as num from products group by category_id) as p on id = category_id")
	err = joined.Order("id desc").Offset(offsetNum).Limit(pageSize).Find(&categories).Error
	return categories, err
}

func (c *CategoryRepository) GetCategoryNumber() (number int) {
	db.Model(&entity.Category{}).Count(&number)
	return number
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

func (c *CategoryRepository) GetCategoryIDByName(name string) (id int) {
	var category entity.Category
	db.Select("id").Where("name = ?", name).First(&category)
	return category.ID
}