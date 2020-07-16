package service

import (
	"backend-go/dao"
	"backend-go/entity"
)

var cartService *CartService

type CartService struct {
	cartRepository *dao.CartRepository
}

func init() {
	cartService = &CartService{dao.GetCartRepository()}
}

func GetCartService() *CartService {
	return cartService
}

func (c *CartService) GetCartProducts(id int) []entity.CartProduct {
	return c.cartRepository.GetUserProducts(id)
}

func (c *CartService) AddProduct(cart *entity.Cart) (status int) {
	num := cart.Num
	if num == 0 {
		num = 1
	}
	cart.Num = 0
	c.cartRepository.GetCartByID(cart)

	var err error
	if cart.Num == 0 {
		cart.Num = num
		err = c.cartRepository.AddCart(cart)
	} else {
		err = c.cartRepository.UpdateCartNum(cart, cart.Num+num)
	}

	if err != nil {
		return 1
	} else {
		return 0
	}
}

func (c *CartService) ModifyProduct(cart *entity.Cart) (status int) {
	err := c.cartRepository.UpdateCartNum(cart, cart.Num)
	if err != nil {
		return 1
	} else {
		return 0
	}
}

func (c *CartService) DeleteProduct(productID, userID int) (status int) {
	cart := &entity.Cart{UserID: userID, ProductID: productID}
	err := c.cartRepository.DeleteProduct(cart)
	if err != nil {
		return 1
	} else {
		return 0
	}
}
