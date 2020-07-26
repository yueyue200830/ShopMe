// Web api
package api

import (
	"backend-go/conf"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	gin.SetMode(conf.RunMode)

	router := gin.Default()
	userApiRegister(router)
	bannerApiRegister(router)
	productApiRegister(router)
	categoryApiRegister(router)
	productDetailApiRegister(router)
	cartApiRegister(router)
	orderApiRegister(router)
	managerApiRegister(router)

	err := router.Run(fmt.Sprintf(":%d", conf.HTTPPort))
	if err != nil {
		fmt.Println("Start server failed!")
	}
}
