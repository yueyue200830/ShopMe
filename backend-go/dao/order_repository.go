package dao

import (
	"backend-go/entity"
	"errors"
	"time"
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

//func (o *OrderRepository) GetOrdersByID(userId int) []entity.Order {
//	var orders []entity.Order
//	db.Where("user_id = ?", userId).Find(&orders)
//	return orders
//}

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

func (o *OrderRepository) GetOrdersWithUserByPage(page, pageSize int) ([]entity.OrderWithUser, error) {
	var orders []entity.OrderWithUser
	offSetNum := (page - 1) * pageSize
	err := db.Table("orders").Order("id desc").Offset(offSetNum).Limit(pageSize).Scan(&orders).Error
	return orders, err
}

func (o *OrderRepository) GetOrderNumber() (number int) {
	db.Table("orders").Count(&number)
	return number
}

func (o *OrderRepository) GetProductsByOrderID(orderID int) []entity.ProductOfOrder {
	var products []entity.ProductOfOrder
	db.Raw("select * from (select * from order_products where order_id = ?) as order_products join products on product_id = id", orderID).Scan(&products)
	return products
}

func (o *OrderRepository) UpdateOrderStatusToCancel(order *entity.Order) error {
	updateMap := map[string]interface{}{"status": entity.Canceled, "cancel_time": time.Now()}
	return db.Model(&order).UpdateColumns(updateMap).Error
}

func (o *OrderRepository) UpdateOrderStatusToPaid(order *entity.Order) error {
	updateMap := map[string]interface{}{"status": entity.Paid, "pay_time": time.Now()}
	return db.Model(&order).UpdateColumns(updateMap).Error
}

func (o *OrderRepository) UpdateOrderStatusToFinish(order *entity.Order) error {
	updateMap := map[string]interface{}{"status": entity.Finished, "finish_time": time.Now()}
	return db.Model(&order).UpdateColumns(updateMap).Error
}

func (o *OrderRepository) InsertOrderAndProducts(orderProducts *entity.DetailOrderWithProducts) (error, int) {
	tx := db.Begin()
	sum := float32(0)

	// Check stock and delete
	for _, product := range orderProducts.Products {
		var p entity.Product
		if err := tx.First(&p, product.ID).Error; err != nil {
			tx.Rollback()
			return err, -1
		}
		if p.Stock < product.Num {
			tx.Rollback()
			return errors.New("stock is not enough"), -1
		}
		if err := tx.Model(&p).UpdateColumn("stock", p.Stock-product.Num).Error; err != nil {
			tx.Rollback()
			return err, -1
		}
		sum = sum + float32(product.Num)*product.Price

	}

	// create order and get order id
	order := orderProducts.Order
	order.Sum = sum
	order.Status = entity.Ordered
	order.OrderTime = time.Now()
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err, -1
	}

	// create order product
	orderID := order.ID
	for _, product := range orderProducts.Products {
		orderProduct := entity.OrderProduct{OrderID: orderID, ProductID: product.ID, Num: product.Num, Price: product.Price}
		if err := tx.Create(&orderProduct).Error; err != nil {
			tx.Rollback()
			return err, -1
		}
	}

	// delete shopping cart
	for _, product := range orderProducts.Products {
		cart := &entity.Cart{UserID: order.UserID, ProductID: product.ID, Num: product.Num}
		if err := tx.Delete(&cart).Error; err != nil {
			tx.Rollback()
			return err, -1
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err, -1
	}

	return nil, orderID
}

func (o *OrderRepository) GetFinishOrderSum() float32 {
	var s []float32
	db.Table("orders").Where("status = ?", entity.Finished).Select("sum(sum) as total").Limit(1).Pluck("total", &s)
	return s[0]
}

func (o *OrderRepository) GetOrderDataByDate(t time.Time) (val entity.OrderCount) {
	val.Date = t.Format("2006-01-02")
	db.Table("orders").Where("datediff ( finish_time, ? ) = 0", val.Date).Select("sum(sum) as total, count(*) as num").Limit(1).Scan(&val)
	return val
}