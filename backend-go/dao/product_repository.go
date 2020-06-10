package dao

type Product struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	Stock      int     `json:"stock"`
	Price      float32 `json:"price"`
	Image      string  `json:"image"`
	CategoryID int     `json:"categoryID"`
}

var productRepository *ProductRepository

type ProductRepository struct {
}

func init() {
	productRepository = &ProductRepository{}
}

func GetProductRepository() *ProductRepository {
	return productRepository
}

func (p *ProductRepository) GetAllProducts() []Product {
	var products []Product
	db.Find(&products)
	return products
}

func (p *ProductRepository) GetNewProducts(num int) []Product {
	var products []Product
	db.Limit(num).Find(&products)
	return products
}

func (p *ProductRepository) GetNewProductsOnCategory(num, categoryID int) []Product {
	var products []Product
	db.Where("category_id = ?", categoryID).Limit(num).Find(&products)
	return products
}

func (p *ProductRepository) GetProductNumber() int {
	var number int
	db.Model(&Product{}).Count(&number)
	return number
}

func (p *ProductRepository) GetProductNumberByCategory(categoryID int) int {
	var number int
	db.Model(&Product{}).Where("category_id = ?", categoryID).Count(&number)
	return number
}

func (p *ProductRepository) GetProductsByPage(page, pageSize int) []Product {
	var products []Product
	db.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&products)
	return products
}

func (p *ProductRepository) GetProductsByPageAndCategory(page, pageSize, categoryID int) []Product {
	var products []Product
	db.Where("category_id = ?", categoryID).Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&products)
	return products
}

func (p *ProductRepository) GetProductByID(id int) Product {
	var product Product
	db.First(&product, id)
	return product
}
