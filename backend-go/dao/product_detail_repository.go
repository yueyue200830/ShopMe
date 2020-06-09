package dao

type ProductDetail struct {
	ProductID     int    `json:"productID" gorm:"primary"`
	Detail string `json:"detailPath"`
	Order  int    `json:"order" gorm:"primary;unique"`
}

var productDetailRepository *ProductDetailRepository

type ProductDetailRepository struct {
}

func init() {
	productDetailRepository = &ProductDetailRepository{}
}

func GetProductDetailRepository() *ProductDetailRepository {
	return productDetailRepository
}

func (p *ProductDetailRepository) GetDetailsByID(id int) []ProductDetail {
	var productDetails []ProductDetail
	db.Where("product_id = ?", id).Order("order").Find(&productDetails)
	return productDetails
}

