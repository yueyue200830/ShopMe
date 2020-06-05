package dao

import "fmt"

type User struct {
	ID           int
	Name         string
	Age          int
}

type Product struct {
	Code string `gorm:"primary_key"`
	Price int
}

func Insert() {

	db.Table("table1").Create(&Product{Code: "L11123", Price: 1000})
}

func GetAll() {
	var products []Product
	db.Table("table1").Find(&products)
	for _, p := range products {
		fmt.Println(p.Code, p.Price)
	}
}

func GetOne() {
	var product Product
	err := db.Table("table1").First(&product).Error
	fmt.Println("err: ", err)
	fmt.Println("price: ", product.Price)
}

func AddUser() {
	db.Create(&User{Name: "Apple", Age: 20})
}

func GetUsers() []User {
	var users []User
	db.Find(&users)
	return users
}