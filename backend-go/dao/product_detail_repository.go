package dao

import "backend-go/entity"

var productDetailRepository *ProductDetailRepository

type ProductDetailRepository struct {
}

func init() {
	productDetailRepository = &ProductDetailRepository{}
}

func GetProductDetailRepository() *ProductDetailRepository {
	return productDetailRepository
}

func (p *ProductDetailRepository) GetDetailsByID(id int) []entity.ProductDetail {
	var productDetails []entity.ProductDetail
	db.Where("product_id = ?", id).Order("order").Find(&productDetails)
	return productDetails
}

