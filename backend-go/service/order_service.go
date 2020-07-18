package service

import (
	"backend-go/dao"
	"backend-go/entity"
	"time"
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

func (o *OrderService) GetOrder(id int) *entity.DetailOrderWithProducts {
	order := o.orderRepository.GetOrderByID(id)
	products := o.orderRepository.GetProductsByOrderID(order.ID)
	detailOrder := &entity.DetailOrderWithProducts{Order: *order, Products: products}
	return detailOrder
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

func (o *OrderService) CancelOrder(order *entity.Order) int {
	err := o.orderRepository.UpdateOrderStatusToCancel(order)
	if err != nil {
		return 1
	}
	return 0
}

func (o *OrderService) PayOrder(order *entity.Order) int {
	err := o.orderRepository.UpdateOrderStatusToPaid(order)
	if err != nil {
		return 1
	}
	return 0
}

func (o *OrderService) NewOrder(orderProducts *entity.DetailOrderWithProducts) (status, id int) {
	var err error
	if err, id = o.orderRepository.InsertOrderAndProducts(orderProducts); err != nil {
		status = 1
	} else {
		status = 0
	}
	return status, id
}

func (o *OrderService) FinishOrder(order *entity.Order) int {
	err := o.orderRepository.UpdateOrderStatusToFinish(order)
	if err != nil {
		return 1
	}
	return 0
}

func (o *OrderService) GetOrders(page, pageSize int) ([]entity.OrderWithUser, int, int) {
	status := 0
	num := 0
	orders, err := o.orderRepository.GetOrdersWithUserByPage(page, pageSize)
	if err != nil {
		status = 1
	} else {
		for i, order := range orders {
			name, err := userService.userRepository.GetUserName(order.UserID)
			if err != nil {
				status = 1
				break
			} else {
				orders[i].Name = name
			}
		}
		num = o.orderRepository.GetOrderNumber()
	}
	return orders, num, status
}

func (o *OrderService) GetOrderNumber() int {
	return o.orderRepository.GetOrderNumber()
}

func (o *OrderService) GetOrderTotal() float32 {
	return o.orderRepository.GetFinishOrderSum()
}

func (o *OrderService) GetRecentData() ([]string, []int, []float32) {
	//data := make([]entity.OrderCount, 7)
	date := make([]string, 7)
	num := make([]int, 7)
	sum := make([]float32, 7)
	t := time.Now()
	for i := 0; i < 7; i++ {
		data := o.orderRepository.GetOrderDataByDate(t.AddDate(0, 0, i-7))
		date[i] = data.Date
		num[i] = data.Num
		sum[i] = data.Total
	}
	return date, num, sum
}