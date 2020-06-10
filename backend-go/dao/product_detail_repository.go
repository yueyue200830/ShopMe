package dao

type ProductDetail struct {
	ProductID     int    `json:"productID" gorm:"primary_key;auto_increment:false"`
	Detail string `json:"detailPath"`
	Order  int    `json:"order" gorm:"primary_key;unique;auto_increment:false"`
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

