package dao

import "backend-go/entity"

var productRepository *ProductRepository

type ProductRepository struct {
}

func init() {
	productRepository = &ProductRepository{}
}

func GetProductRepository() *ProductRepository {
	return productRepository
}

func (p *ProductRepository) GetAllProducts() []entity.Product {
	var products []entity.Product
	db.Find(&products)
	return products
}

func (p *ProductRepository) GetNewProducts(num int) []entity.Product {
	var products []entity.Product
	db.Limit(num).Find(&products)
	return products
}

func (p *ProductRepository) GetNewProductsOnCategory(num, categoryID int) []entity.Product {
	var products []entity.Product
	db.Where("category_id = ?", categoryID).Limit(num).Find(&products)
	return products
}

func (p *ProductRepository) GetProductNumber() (number int) {
	db.Model(&entity.Product{}).Count(&number)
	return number
}

func (p *ProductRepository) GetProductNumberByCategory(categoryID int) (number int) {
	db.Model(&entity.Product{}).Where("category_id = ?", categoryID).Count(&number)
	return number
}

func (p *ProductRepository) GetProductsByPage(page, pageSize int) []entity.Product {
	var products []entity.Product
	db.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&products)
	return products
}

func (p *ProductRepository) GetProductsByPageAndCategory(page, pageSize, categoryID int) []entity.Product {
	var products []entity.Product
	offsetNum := (page - 1) * pageSize
	db.Where("category_id = ?", categoryID).Order("id desc").Offset(offsetNum).Limit(pageSize).Find(&products)
	return products
}

func (p *ProductRepository) GetProductByID(id int) entity.Product {
	var product entity.Product
	db.First(&product, id)
	return product
}

func (p *ProductRepository) GetAllProductNames() (names []string) {
	db.Table("products").Pluck("title", &names)
	return names
}

func (p *ProductRepository) AddProduct(product *entity.Product) error {
	return db.Create(&product).Error
}

func (p *ProductRepository) UpdateProduct(product *entity.Product) error {
	return db.Save(&product).Error
}

func (p *ProductRepository) DeleteProduct(id int) error {
	return db.Where("id = ?", id).Delete(&entity.Product{}).Error
}