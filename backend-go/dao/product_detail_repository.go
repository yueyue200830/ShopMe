package dao

import (
	"backend-go/entity"
)

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

func (p *ProductDetailRepository) GetDetailNumber(id int) (number int, err error) {
	err = db.Table("product_details").Where("product_id = ?", id).Count(&number).Error
	return number, err
}

func (p *ProductDetailRepository) UpdateDetail(detail entity.ProductDetail) error {
	return db.Save(&detail).Error
}

func (p *ProductDetailRepository) CreateDetail(detail entity.ProductDetail) error {
	return db.Create(detail).Error
}

func (p *ProductDetailRepository) DeleteDetail(id, order int) error {
	return db.Delete(entity.ProductDetail{Order: order, ProductID: id}).Error
}
