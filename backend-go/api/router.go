// Web api
package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitServer() {
	router = gin.Default()
	UserApiRegister()
	err := router.Run(":8000") // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Println("Start server failed!")
	}
}
