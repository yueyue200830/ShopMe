package dao

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	AvatarId int    `json:"avatar_id" gorm:"column:avatar_id"`
}

func GetAll() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetAllUserNames() []User {
	var users []User
	db.Select("name").Find(&users)
	return users
}

func GetAllUsers() []User {
	var users []User
	db.Select("id, name, email, avatar_id").Find(&users)
	return users
}
