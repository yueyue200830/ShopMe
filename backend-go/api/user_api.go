package api

import (
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func UserApiRegister() {
	curd := UserController{}
	router.GET("/getAll", curd.GetAll)
	router.GET("/getUsers", curd.GetAllUsers)
	router.GET("/getUserNames", curd.GetAllNames)
}

func (userController *UserController) GetAll(c *gin.Context) {
	users := service.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (userController *UserController) GetAllNames(c *gin.Context) {
	userNames := service.GetAllUserNames()
	c.JSON(http.StatusOK, gin.H{
		"names": userNames,
	})
}

func (userController *UserController) GetAllUsers(c *gin.Context) {
	users := service.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
