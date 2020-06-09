package service

import (
	"backend-go/dao"
)

type PromoteSection struct {
	Title    string        `json:"title"`
	Category int           `json:"category"`
	Products []dao.Product `json:"products"`
}

var productService *ProductService

type ProductService struct {
	productRepository *dao.ProductRepository
}

func init() {
	productService = &ProductService{dao.GetProductRepository()}
}

func GetProductService() *ProductService {
	return productService
}

func (p *ProductService) GetAllProducts() []dao.Product {
	return p.productRepository.GetAllProducts()
}

func (p *ProductService) GetPromoteProducts() []PromoteSection {
	num := 8
	promoteSections := make([]PromoteSection, 0)
	promoteSections = append(promoteSections, PromoteSection{"新品上市", 0, p.productRepository.GetNewProducts(num)})
	categories := dao.GetCategoryRepository().GetAllCategories()
	for _, category := range categories {
		promoteSections = append(promoteSections, PromoteSection{category.Name, category.ID, p.productRepository.GetNewProductsOnCategory(num, category.ID)})
	}
	return promoteSections
}

func (p *ProductService) GetProductPath(image string) string {
	return "../images/products/" + image
}

func (p *ProductService) GetProductPages(pageSize, categoryID int) int {
	var number int
	if categoryID == 0 {
		number = p.productRepository.GetProductNumber()
	} else {
		number = p.productRepository.GetProductNumberByCategory(categoryID)
	}
	pages := number / pageSize
	if number % pageSize != 0 {
		pages = pages + 1
	}
	return pages
}

func (p *ProductService) GetProductNumber(categoryID int) int {
	var number int
	if categoryID == 0 {
		number = p.productRepository.GetProductNumber()
	} else {
		number = p.productRepository.GetProductNumberByCategory(categoryID)
	}
	return number
}

func (p *ProductService) GetProductsByPage(page, pageSize, categoryID int) []dao.Product {
	if categoryID == 0 {
		return p.productRepository.GetProductsByPage(page, pageSize)
	} else {
		return p.productRepository.GetProductsByPageAndCategory(page, pageSize, categoryID)
	}
}

func (p *ProductService) GetProductInfo(id int) dao.Product {
	return p.productRepository.GetProductByID(id)
}