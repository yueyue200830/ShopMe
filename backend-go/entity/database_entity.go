package entity

import "time"

type Banner struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Banner    string `json:"bannerPath"`
	ProductID int    `json:"productID"`
}

type Cart struct {
	UserID    int `json:"userID" gorm:"primary_key;auto_increment:false"`
	ProductID int `json:"productID" gorm:"primary_key;auto_increment:false"`
	Num       int `json:"num"`
}

type Category struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type OrderStatus string

const (
	Ordered  OrderStatus = "ordered"
	Paid     OrderStatus = "paid"
	Finished OrderStatus = "finished"
	Canceled OrderStatus = "canceled"
)

type Order struct {
	ID         int         `json:"id"`
	UserID     int         `json:"userID"`
	Status     OrderStatus `json:"status" gorm:"type:enum"`
	Sum        float32     `json:"sum"`
	OrderTime  time.Time   `json:"orderTime"`
	PayTime    *time.Time  `json:"payTime"`
	FinishTime *time.Time  `json:"finishTime"`
	CancelTime *time.Time  `json:"cancelTime"`
}

type OrderProduct struct {
	OrderID   int     `json:"orderID" gorm:"primary_key;auto_increment:false"`
	ProductID int     `json:"productID" gorm:"primary_key;auto_increment:false"`
	Num       int     `json:"num"`
	Price     float32 `json:"price"`
}

type ProductDetail struct {
	ProductID int    `json:"productID" gorm:"primary_key;auto_increment:false"`
	Detail    string `json:"detailPath"`
	Order     int    `json:"order" gorm:"primary_key;unique;auto_increment:false"`
}

type Product struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	Stock      int     `json:"stock"`
	Price      float32 `json:"price"`
	Image      string  `json:"image"`
	CategoryID int     `json:"categoryID"`
}

type User struct {
	ID       int    `json:"id" gorm:"AUTO_INCREMENT"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatarPath" gorm:"default:NULL"`
}

type Manager struct {
	ID       int    `json:"id" gorm:"AUTO_INCREMENT"`
	Name     string `json:"userName"`
	Password string `json:"password"`
}

type CategoryWithNum struct {
	Category
	Num int `json:"num"`
}
