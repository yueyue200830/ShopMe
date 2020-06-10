// Service for api
package service

import (
	"backend-go/dao"
	"backend-go/entity"
	"fmt"
	"regexp"
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

func (u *UserService) GetAll() []entity.User {
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

func (u *UserService) GetAllUsers() []entity.User {
	return u.userRepository.GetAllUsers()
}

// Return 0 for not found
func (u *UserService) ValidateUser(user entity.User) (id int) {
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

// Check whether user email exists expect id itself
func (u *UserService) ValidateUserEmail(email string, id int) bool {
	userID := u.userRepository.GetUserIDByEmail(email)
	if userID > 0 && id != userID {
		return true
	} else {
		return false
	}
}

// 0 -> ok, 1 -> nameError, 2 -> emailError, 3 -> passwordError, 4 -> other
func (u *UserService) Register(user entity.User) int {
	match, err := regexp.MatchString("^[a-zA-Z0-9_\u4e00-\u9fa5]{4,20}$", user.Name)
	if err != nil {
		fmt.Println(err)
		return 4
	} else if !match || u.ValidateUserName(user.Name, 0){
		return 1
	}
	match, err = regexp.MatchString("^[a-zA-Z0-9]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$", user.Email)
	if err != nil {
		fmt.Println(err)
		return 4
	} else if !match || u.ValidateUserEmail(user.Email, 0){
		return 2
	}
	// check whether there is alphabet & number
	match, err = regexp.MatchString("(([a-zA-Z].*[0-9])|([0-9].*[a-zA-Z]))", user.Password)
	if err != nil {
		fmt.Println(err)
		return 4
	} else if !match || len(user.Password) > 16 || len(user.Password) < 6 {
		return 3
	}
	return u.userRepository.AddUser(user)
}