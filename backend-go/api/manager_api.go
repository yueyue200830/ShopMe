package api

import (
	"backend-go/entity"
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ManagerController struct {
	managerService *service.ManagerService
}

func managerApiRegister(router *gin.Engine) {
	curd := ManagerController{service.GetManagerService()}
	router.POST("/manager/login", curd.managerLogin)
	//router.GET("/manager/info", curd.getManagerInfo)
	router.POST("/manager/logout", curd.managerLogout)
}

func (m *ManagerController) managerLogin(c *gin.Context) {
	var manager *entity.Manager
	status := 0
	token := -1
	if err := c.ShouldBindJSON(&manager); err != nil {
		status = 1
	} else {
		token = m.managerService.Login(manager)
		if token == -1 {
			status = 1
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": gin.H{
			"token": token,
		},
	})
}

func (m *ManagerController) managerLogout(c *gin.Context) {
	c.JSON(http.StatusOK, 0)
}

//func (m *ManagerController) getManagerInfo(c *gin.Context) {
//
//}