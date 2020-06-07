package dao

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   int    `json:"avatarPath"`
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

