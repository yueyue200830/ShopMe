package service

import (
	"backend-go/dao"
	"backend-go/entity"
)

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

func (p *ProductService) GetAllProducts() []entity.Product {
	return p.productRepository.GetAllProducts()
}

func (p *ProductService) GetPromoteProducts() []entity.PromoteSection {
	num := 8
	promoteSections := make([]entity.PromoteSection, 0)
	promoteSections = append(promoteSections, entity.PromoteSection{Title: "新品上市", Products: p.productRepository.GetNewProducts(num)})
	categories := dao.GetCategoryRepository().GetAllCategories()
	for _, category := range categories {
		promoteSections = append(promoteSections, entity.PromoteSection{Title: category.Name, Category: category.ID, Products: p.productRepository.GetNewProductsOnCategory(num, category.ID)})
	}
	return promoteSections
}

func (p *ProductService) GetProductPath(image string) string {
	return "../images/products/" + image
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

func (p *ProductService) GetProductsByPage(page, pageSize, categoryID int) []entity.Product {
	if categoryID == 0 {
		return p.productRepository.GetProductsByPage(page, pageSize)
	} else {
		return p.productRepository.GetProductsByPageAndCategory(page, pageSize, categoryID)
	}
}

func (p *ProductService) GetProductInfo(id int) entity.Product {
	return p.productRepository.GetProductByID(id)
}