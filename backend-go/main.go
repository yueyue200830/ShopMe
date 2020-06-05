package main

import (
	"backend-go/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		//dao.GetAll()
		//dao.GetOne()
		//dao.AddUser()
		users := dao.GetUsers()
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	defer dao.Close()
}
