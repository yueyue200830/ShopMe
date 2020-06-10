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
	productDetailApiRegister(router)
	cartApiRegister(router)
	orderApiRegister(router)
	err := router.Run(":8000")
	if err != nil {
		fmt.Println("Start server failed!")
	}
}