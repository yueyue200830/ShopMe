// Web api
package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()
	userApiRegister(router)
	bannerApiRegister(router)
	productApiRegister(router)
	categoryApiRegister(router)
	err := router.Run(":8000") // listen and serve on 0.0.0.0:8080
	if err != nil {
		fmt.Println("Start server failed!")
	}
}