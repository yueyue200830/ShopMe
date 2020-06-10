package dao

import "backend-go/entity"

var orderRepository *OrderRepository

type OrderRepository struct {
}

func init() {
	orderRepository = &OrderRepository{}
}

func GetOrderRepository() *OrderRepository {
	return orderRepository
}

func (o *OrderRepository) GetOne() entity.Order {
	var order entity.Order
	db.First(&order)
	return order
}
