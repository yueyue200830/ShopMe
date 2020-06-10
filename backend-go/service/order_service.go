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

func (o *OrderService) GetOne() entity.Order {
	return o.orderRepository.GetOne()
}