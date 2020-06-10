package service

import "backend-go/dao"

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

func (o *OrderService) GetOne() dao.Order {
	return o.orderRepository.GetOne()
}