package api

import (
	"backend-go/entity"
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
	//router.GET("/getAll", curd.getAll)
	//router.GET("/getUsers", curd.getAllUsers)
	//router.GET("/getUserNames", curd.getAllNames)
	router.GET("/users", curd.getUsers)
	router.POST("/userLogin", curd.userLogin)
	router.GET("/checkUserNameExist", curd.checkUserNameExist)
	router.GET("/checkUserEmailExist", curd.checkUserEmailExist)
	router.POST("/register", curd.register)
	router.GET("/avatar/:image", curd.getAvatar)
	router.PUT("/userPassword", curd.updatePassword)
	router.GET("/user", curd.getUserInfo)
	router.PUT("/user", curd.updateUserInfo)
	router.GET("/user/number", curd.getUserNumber)
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

// return 0 for not found or other error input cases
func (u *UserController) userLogin(c *gin.Context) {
	var user entity.User
	id := -1
	if err := c.ShouldBindJSON(&user); err != nil {
		id = 0
	}
	if user.Name == "" || user.Password == "" {
		id = 0
	}
	if id == -1 {
		id = u.userService.ValidateUser(user)
	}

	c.JSON(http.StatusOK, id)
}

// Return 1 when name exist, 0 otherwise
func (u *UserController) checkUserNameExist(c *gin.Context) {
	var exist int
	name := c.Query("name")
	id, err := strconv.Atoi(c.DefaultQuery("id", "0"))
	if err != nil || id < 0 {
		exist = 0
	} else {
		if u.userService.ValidateUserName(name, id) {
			exist = 1
		} else {
			exist = 0
		}
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
	var user entity.User
	status := -1
	if err := c.ShouldBindJSON(&user); err != nil {
		status = 4
	} else if user.Name == "" || user.Password == "" || user.Email == "" {
		status = 4
	} else {
		status = u.userService.Register(user)
	}

	c.JSON(http.StatusOK, status)
}

func (u *UserController) getAvatar(c *gin.Context) {
	avatarPath := u.userService.GetAvatarPath(c.Param("image"))
	c.File(avatarPath)
}

func (u *UserController) updatePassword(c *gin.Context) {
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	status := 0
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		status = 1
	} else {
		status = u.userService.UpdatePassword(id, oldPassword, newPassword)
	}
	c.JSON(http.StatusOK, status)
}

func (u *UserController) getUserInfo(c *gin.Context) {
	status := 0
	var info map[string]interface{}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		status = 1
	} else {
		info = u.userService.GetUserInfo(id)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": info,
	})
}

func (u *UserController) updateUserInfo(c *gin.Context) {
	var user *entity.User
	status := 0
	if err := c.ShouldBindJSON(&user); err != nil {
		status = 1
	} else {
		status = u.userService.UpdateUserInfo(user)
	}
	c.JSON(http.StatusOK, status)
}

func (u *UserController) getUsers(c *gin.Context) {
	status := 0
	num := 0
	var users []entity.User
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		status = 1
	}
	pageSize, err := strconv.Atoi(c.Query("size"))
	if err != nil || pageSize < 1 {
		status = 1
	}
	if status == 0 {
		users, num, status = u.userService.GetUsers(page, pageSize)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": gin.H{
			"num":  num,
			"list": users,
		},
	})
}

func (u *UserController) getUserNumber(c *gin.Context) {
	num := u.userService.GetUserNumber()
	c.JSON(http.StatusOK, num)
}
