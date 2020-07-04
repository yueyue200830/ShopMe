package service

import (
	"backend-go/dao"
	"backend-go/entity"
)

var orderService *OrderService

type OrderService struct {
	orderRepository *dao.OrderRepository
}

func init() {
	orderService = &OrderService{dao.GetOrderRepository()}
}

func GetOrderService() *OrderService {
	return orderService
}

func (o *OrderService) GetOrder(id int) *entity.Order {
	return o.orderRepository.GetOrderByID(id)
}

func (o *OrderService) GetUserOrders(userID int) []entity.Order {
	return o.orderRepository.GetOrdersByID(userID)
}

func (o *OrderService) GetUserOrderNumber(userID int) (number int) {
	return o.orderRepository.GetOrderNumberByID(userID)
}

func (o *OrderService) GetUserOrdersByPage(page, pageSize, userID int) []entity.SimpleOrderWithProducts {
	orders := o.orderRepository.GetOrdersByUserIDAndPage(page, pageSize, userID)
	ordersWithProducts := make([]entity.SimpleOrderWithProducts, 0)
	for _, order := range orders {
		products := o.orderRepository.GetProductsByOrderID(order.ID)
		date := order.OrderTime.Format("2006-01-02")
		curOrder := entity.SimpleOrderWithProducts{ID: order.ID, Date: date, Sum: order.Sum, Status: order.Status, Products: products}
		ordersWithProducts = append(ordersWithProducts, curOrder)
	}
	return ordersWithProducts
}
