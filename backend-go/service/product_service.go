package service

import (
	"backend-go/dao"
	"backend-go/entity"
	"backend-go/utils"
	"strings"
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

func (p *ProductService) GetProductImagePath(image string) string {
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

func (p *ProductService) GetProductsNames() []string {
	return p.productRepository.GetAllProductNames()
}

func (p *ProductService) AddProduct(product *entity.Product) int {
	image := strings.Split(product.Image, "/")
	product.Image = image[len(image) - 1]

	err := p.productRepository.AddProduct(product)
	if err != nil {
		return 1
	}
	return 0
}

func (p *ProductService) UpdateProduct(product *entity.Product) int {
	image := strings.Split(product.Image, "/")
	product.Image = image[len(image) - 1]

	err := p.productRepository.UpdateProduct(product)
	if err != nil {
		return 1
	}
	return 0
}

func (p *ProductService) DeleteProduct(id int) int {
	// todo: delete product image & product detail as well
	// todo: consider those product in the order
	err := p.productRepository.DeleteProduct(id)
	if err != nil {
		return 1
	}
	return 0
}

func (p *ProductService) GenerateRandomImageName(fileType string) (name string, status int) {
	status = 0
	s := strings.Split(fileType, "/")

	if len(s) != 2 || s[0] != "image" {
		status = 1
	} else {
		name = utils.GenerateRandomString(30)
		name = name + "." + s[1]
	}

	return name, status
}