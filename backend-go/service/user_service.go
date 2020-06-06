// Service for api
package service

import (
	"backend-go/dao"
)

func GetAll() []dao.User {
	return dao.GetAll()
}

func GetAllUserNames() []string {
	userNames := dao.GetAllUserNames()
	var names []string
	for _, user := range userNames {
		names = append(names, user.Name)
	}
	return names
}

func GetAllUsers() []dao.User {
	return dao.GetAllUsers()
}
