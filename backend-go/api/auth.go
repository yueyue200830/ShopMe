package api

import (
	"backend-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func authApiRegister(router, authRouter *gin.RouterGroup) {
	router.GET("/auth", getAuth)
}

func getAuth(c *gin.Context) {
	status := 0
	token, err := utils.GenerateToken("name", 32)
	if err != nil {
		status = 1
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : status,
		"data" : token,
	})
}