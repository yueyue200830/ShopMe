package dao

import (
	"backend-go/entity"
	"github.com/jinzhu/gorm"
)

var cartRepository *CartRepository

type CartRepository struct {
}

func init() {
	cartRepository = &CartRepository{}
}

func GetCartRepository() *CartRepository {
	return cartRepository
}

func (c *CartRepository) GetUserProducts(id int) []entity.CartProduct {
	var products []entity.CartProduct
	db.Raw("select * from (select * from carts where user_id = ?) as cart join products on product_id = id", id).Scan(&products)
	return products
}

func (c *CartRepository) GetCartByID(cart *entity.Cart) {
	db.Where(&cart).First(&cart)
}

// Set cart number to 0 for finding the record, then update to num
func (c *CartRepository) UpdateCartNum(cart *entity.Cart, num int) error {
	cart.Num = 0
	return db.Model(&cart).Where(&cart).Update("num", gorm.Expr("?", num)).Error
}

func (c *CartRepository) AddCart(cart *entity.Cart) error {
	return db.Create(&cart).Error
}

func (c *CartRepository) DeleteProduct(cart *entity.Cart) error {
	return db.Model(&cart).Where(&cart).Delete(&cart).Error
}
