package dao

import "fmt"

type User struct {
	ID       int    `json:"id" gorm:"AUTO_INCREMENT"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   int    `json:"avatarPath" gorm:"default:NULL"`
}

var userRepository *UserRepository

type UserRepository struct {
}

func init() {
	userRepository = &UserRepository{}
}

func GetUserRepository() *UserRepository {
	return userRepository
}

func (u *UserRepository) GetAll() []User {
	var users []User
	db.Find(&users)
	return users
}

func (u *UserRepository) GetAllUserNames() []User {
	var users []User
	db.Select("name").Find(&users)
	return users
}

func (u *UserRepository) GetAllUsers() []User {
	var users []User
	db.Select("id, name, email, avatar").Find(&users)
	return users
}

func (u *UserRepository) GetUserIDByNameAndPassword(user User) int {
	db.Where(&user).First(&user)
	return user.ID
}

func (u *UserRepository) GetUserIDByName(name string) int {
	var user User
	db.Select("id").Where("name = ?", name).First(&user)
	return user.ID
}

func (u *UserRepository) GetUserIDByEmail(email string) int {
	var user User
	db.Select("id").Where("email = ?", email).First(&user)
	return user.ID
}

func (u *UserRepository) AddUser(user User) int {
	if err := db.Create(&user).Error; err != nil {
		fmt.Println(err)
		return 4
	}
	return 0
}