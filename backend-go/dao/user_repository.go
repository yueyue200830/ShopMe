package dao

import (
	"backend-go/entity"
	"errors"
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

func (u *UserRepository) GetUsersByPage(page, size int) (users []entity.User, err error) {
	offsetNum := (page - 1) * size
	dbUsers := db.Table("users").Order("id desc").Offset(offsetNum).Limit(size)
	err = dbUsers.Select("id, name, email, avatar").Scan(&users).Error
	return users, err
}

func (u *UserRepository) GetUserNumber() (number int) {
	db.Table("users").Count(&number)
	return number
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

func (u *UserRepository) UpdatePassword(id int, oldPassword, newPassword string) error {
	tx := db.Begin()

	user := &entity.User{ID: id}
	if err := tx.First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if user.Password != oldPassword {
		tx.Rollback()
		return errors.New("wrong password")
	}
	if err := tx.Model(&user).Update("password", newPassword).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserInfo(id int) *entity.User {
	user := &entity.User{ID: id}
	db.First(&user)
	return user
}

func (u *UserRepository) UpdateNameAndEmail(user *entity.User) int {
	status := 0
	info := map[string]interface{}{"name": user.Name, "email": user.Email}
	if err := db.Model(&user).Updates(info).Error; err != nil {
		status = 1
	}
	return status
}

func (u *UserRepository) GetUserName(id int) (string, error) {
	user := &entity.User{ID: id}
	err := db.Select("id, name").First(&user).Error
	return user.Name, err
}
