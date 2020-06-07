package dao

type Category struct {
	ID   int    `json:"id" gorm:"primary"`
	Name string `json:"name"`
}

var categoryRepository *CategoryRepository

func init() {
	categoryRepository = &CategoryRepository{}
}

type CategoryRepository struct {
}

func GetCategoryRepository() *CategoryRepository {
	return categoryRepository
}

func (c *CategoryRepository) GetAllCategories() []Category {
	var categories []Category
	db.Find(&categories)
	return categories
}
