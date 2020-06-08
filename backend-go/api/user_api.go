package api

import (
	"backend-go/dao"
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *service.UserService
}

func userApiRegister(router *gin.Engine) {
	curd := UserController{service.GetUserService()}
	router.GET("/getAll", curd.getAll)
	router.GET("/getUsers", curd.getAllUsers)
	router.GET("/getUserNames", curd.getAllNames)
	router.POST("/userLogin", curd.userLogin)
	router.GET("/checkUserNameExist", curd.checkUserNameExist)
	router.GET("/checkUserEmailExist", curd.checkUserEmailExist)
	router.POST("/register", curd.register)
}

func (u *UserController) getAll(c *gin.Context) {
	users := u.userService.GetAll()
	c.JSON(http.StatusOK, users)
}

func (u *UserController) getAllNames(c *gin.Context) {
	userNames := u.userService.GetAllUserNames()
	c.JSON(http.StatusOK, userNames)
}

func (u *UserController) getAllUsers(c *gin.Context) {
	users := u.userService.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func (u *UserController) userLogin(c *gin.Context) {
	var user dao.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Name == "" || user.Password == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	id := u.userService.ValidateUser(user)
	c.JSON(http.StatusOK, id)
}

func (u *UserController) checkUserNameExist(c *gin.Context) {
	name := c.Query("name")
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil || id < 0 {
		c.Status(http.StatusBadRequest)
		return
	}
	var exist int
	if u.userService.ValidateUserName(name, id) {
		exist = 1
	} else {
		exist = 0
	}
	c.JSON(http.StatusOK, exist)
}

func (u *UserController) checkUserEmailExist(c *gin.Context) {
	email := c.Query("email")
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil || id < 0 {
		c.Status(http.StatusBadRequest)
		return
	}
	var exist int
	if u.userService.ValidateUserEmail(email, id) {
		exist = 1
	} else {
		exist = 0
	}
	c.JSON(http.StatusOK, exist)
}

func (u *UserController) register(c *gin.Context) {
	var user dao.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Name == "" || user.Password == "" || user.Email == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	status := u.userService.Register(user)
	c.JSON(http.StatusOK, status)
}
