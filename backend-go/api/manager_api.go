package api

import (
	"backend-go/entity"
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ManagerController struct {
	managerService *service.ManagerService
}

func managerApiRegister(router *gin.Engine) {
	curd := ManagerController{service.GetManagerService()}
	router.POST("/manager/login", curd.managerLogin)
	router.GET("/manager/info", curd.getManagerInfo)
	router.POST("/manager/logout", curd.managerLogout)
	router.GET("/managers", curd.getManagers)
	router.PUT("/manager", curd.updateManager)
	router.POST("/manager", curd.createManager)
	router.DELETE("/manager", curd.deleteManager)
	router.PUT("/manager/password/reset", curd.resetPassword)
	router.GET("/manager/name/check", curd.checkManagerNameExist)
	router.GET("/manager/main", curd.getMainData)
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

func (m *ManagerController) getManagerInfo(c *gin.Context) {
	status := 0
	token := c.Query("token")
	name := ""
	if id, err := strconv.Atoi(token); err != nil {
		status = 1
	} else {
		name, status = m.managerService.GetManagerInfo(id)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": gin.H{
			"name":   name,
			"avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		},
	})
}

func (m *ManagerController) getManagers(c *gin.Context) {
	status := 0
	num := 0
	var managers []entity.Manager
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		status = 1
	}
	pageSize, err := strconv.Atoi(c.Query("size"))
	if err != nil || pageSize < 1 {
		status = 1
	}
	if status == 0 {
		managers, num, status = m.managerService.GetManagers(page, pageSize)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": gin.H{
			"num":  num,
			"list": managers,
		},
	})
}

func (m *ManagerController) updateManager(c *gin.Context) {
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	status := 0
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		status = 1
	} else {
		status = m.managerService.UpdateManager(id, oldPassword, newPassword)
	}
	c.JSON(http.StatusOK, status)
}

func (m *ManagerController) createManager(c *gin.Context) {
	status := 0
	password := ""
	var manager *entity.Manager
	if err := c.ShouldBindJSON(&manager); err != nil {
		status = 1
	} else {
		status, password = m.managerService.CreateManager(manager)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": password,
	})
}

func (m *ManagerController) deleteManager(c *gin.Context) {
	status := 0
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		status = 1
	} else {
		status = m.managerService.DeleteManager(id)
	}

	c.JSON(http.StatusOK, status)
}

func (m *ManagerController) resetPassword(c *gin.Context) {
	status := 0
	password := ""
	var manager *entity.Manager
	if err := c.ShouldBindJSON(&manager); err != nil {
		status = 1
	} else {
		status, password = m.managerService.ResetPassword(manager)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": password,
	})
}

func (m *ManagerController) checkManagerNameExist(c *gin.Context) {
	var exist int
	name := c.Query("name")
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil || id < 0 {
		exist = 0
	} else {
		if m.managerService.ValidateManagerName(name, id) {
			exist = 1
		} else {
			exist = 0
		}
	}
	c.JSON(http.StatusOK, exist)
}

func (m *ManagerController) getMainData(c *gin.Context) {
	user := service.GetUserService().GetUserNumber()
	product := service.GetProductService().GetProductNumber(0)
	order := service.GetOrderService().GetOrderNumber()
	sum := service.GetOrderService().GetOrderTotal()

	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"product": product,
		"order": order,
		"sum": sum,
	})
}
