package dao

import "time"

type OrderStatus string

const (
	Ordered  OrderStatus = "ordered"
	Paid     OrderStatus = "paid"
	Canceled OrderStatus = "canceled"
)

type Order struct {
	ID         int         `json:"id"`
	UserID     int         `json:"userID"`
	Status     OrderStatus `json:"status" gorm:"type:enum"`
	Sum        int         `json:"sum"`
	OrderTime  *time.Time  `json:"orderTime"`
	PaidTime   *time.Time  `json:"paidTime"`
	CancelTime *time.Time  `json:"cancelTime"`
}

var orderRepository *OrderRepository

type OrderRepository struct {
}

func init() {
	orderRepository = &OrderRepository{}
}

func GetOrderRepository() *OrderRepository {
	return orderRepository
}

func (o *OrderRepository) GetOne() Order {
	var order Order
	db.First(&order)
	return order
}
