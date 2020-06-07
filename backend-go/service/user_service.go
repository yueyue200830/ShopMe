// Service for api
package service

import (
	"backend-go/dao"
)

var userService *UserService

type UserService struct {
	userRepository *dao.UserRepository
}

func init() {
	userService = &UserService{dao.GetUserRepository()}
}

func GetUserService() *UserService {
	return userService
}

func (u *UserService) GetAll() []dao.User {
	return u.userRepository.GetAll()
}

func (u *UserService) GetAllUserNames() []string {
	userNames := u.userRepository.GetAllUserNames()
	var names []string
	for _, user := range userNames {
		names = append(names, user.Name)
	}
	return names
}

func (u *UserService) GetAllUsers() []dao.User {
	return u.userRepository.GetAllUsers()
}

// Return 0 for not found
func (u *UserService) ValidateUser(user dao.User) (id int) {
	return u.userRepository.GetUserIDByNameAndPassword(user)
}

// Check whether user name exists expect id itself
func (u *UserService) ValidateUserName(name string, id int) bool {
	userID := u.userRepository.GetUserIDByName(name)
	if userID > 0 && id != userID {
		return true
 	} else {
 		return false
	}
}