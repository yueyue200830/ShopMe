package api

import (
	"backend-go/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {

}

func UserApiRegister() {
	curd := UserController{}
	router.GET("/user/all", curd.GetAll)
}

func (userController *UserController) GetAll(c *gin.Context) {
	users := dao.GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}