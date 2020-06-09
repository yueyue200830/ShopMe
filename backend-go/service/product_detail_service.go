package service

import "backend-go/dao"

var productDetailService *ProductDetailService

type ProductDetailService struct {
	productDetailRepository *dao.ProductDetailRepository
}

func init() {
	productDetailService = &ProductDetailService{dao.GetProductDetailRepository()}
}

func GetProductDetailService() *ProductDetailService {
	return productDetailService
}

func (p *ProductDetailService) GetDetailsByID(id int) []dao.ProductDetail {
	return p.productDetailRepository.GetDetailsByID(id)
}

func (p *ProductDetailService) GetDetailPath(name string) string {
	return "../images/details/" + name
}