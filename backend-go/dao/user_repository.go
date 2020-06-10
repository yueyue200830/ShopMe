package dao

import (
	"backend-go/entity"
	"fmt"
)

var userRepository *UserRepository

type UserRepository struct {
}

func init() {
	userRepository = &UserRepository{}
}

func GetUserRepository() *UserRepository {
	return userRepository
}

func (u *UserRepository) GetAll() []entity.User {
	var users []entity.User
	db.Find(&users)
	return users
}

func (u *UserRepository) GetAllUserNames() []entity.User {
	var users []entity.User
	db.Select("name").Find(&users)
	return users
}

func (u *UserRepository) GetAllUsers() []entity.User {
	var users []entity.User
	db.Select("id, name, email, avatar").Find(&users)
	return users
}

func (u *UserRepository) GetUserIDByNameAndPassword(user entity.User) int {
	db.Where(&user).First(&user)
	return user.ID
}

func (u *UserRepository) GetUserIDByName(name string) int {
	var user entity.User
	db.Select("id").Where("name = ?", name).First(&user)
	return user.ID
}

func (u *UserRepository) GetUserIDByEmail(email string) int {
	var user entity.User
	db.Select("id").Where("email = ?", email).First(&user)
	return user.ID
}

func (u *UserRepository) AddUser(user entity.User) int {
	if err := db.Create(&user).Error; err != nil {
		fmt.Println(err)
		return 4
	}
	return 0
}