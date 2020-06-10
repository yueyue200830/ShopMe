package dao

import (
	"backend-go/entity"
)

var orderRepository *OrderRepository

type OrderRepository struct {
}

func init() {
	orderRepository = &OrderRepository{}
}

func GetOrderRepository() *OrderRepository {
	return orderRepository
}

func (o *OrderRepository) GetOrderByID(id int) *entity.Order {
	order := &entity.Order{}
	db.First(&order, id)
	return order
}

func (o *OrderRepository) GetOrdersByID(userId int) []entity.Order {
	var orders []entity.Order
	db.Where("user_id = ?", userId).Find(&orders)
	return orders
}

func (o *OrderRepository) GetOrderNumberByID(userID int) (number int) {
	db.Model(&entity.Order{}).Where("user_id = ?", userID).Count(&number)
	return number
}

func (o *OrderRepository) GetOrdersByUserIDAndPage(page, pageSize, userID int) []entity.Order {
	var orders []entity.Order
	offSetNum := (page - 1) * pageSize
	db.Where("user_id = ?", userID).Order("order_time desc").Offset(offSetNum).Limit(pageSize).Find(&orders)
	return orders
}