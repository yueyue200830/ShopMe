package service

import (
	"backend-go/dao"
	"backend-go/entity"
	"backend-go/utils"
	"strings"
)

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

func (p *ProductDetailService) GetDetailsByID(id int) []entity.ProductDetail {
	return p.productDetailRepository.GetDetailsByID(id)
}

func (p *ProductDetailService) GetDetailPath(name string) string {
	return "../images/details/" + name
}

func (p *ProductDetailService) GenerateRandomImageName(fileType string) (name string, status int) {
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

func (p *ProductDetailService) UpdateDetail(details []entity.ProductDetail) int {
	if len(details) == 0 {
		return 1
	}
	id := details[0].ProductID
	l, err := p.productDetailRepository.GetDetailNumber(id)
	if err != nil {
		return 1
	}

	for i, detail := range details {
		detail.Order = i + 1
		if i < l {
			if err := p.productDetailRepository.UpdateDetail(detail); err != nil {
				return 1
			}
		} else {
			if err := p.productDetailRepository.CreateDetail(detail); err != nil {
				return 1
			}
		}
	}
	if len(details) < l {
		for i := len(details); i < l; i++ {
			if err := p.productDetailRepository.DeleteDetail(id, i+1); err != nil {
				return 1
			}
		}
	}

	return 0
}
