// Web api
package api

import (
	"backend-go/conf"
	"backend-go/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	gin.SetMode(conf.RunMode)
	router := gin.Default()
	authRouter := router.Group("")
	authRouter.Use(middleware.ValidateJWT())
	unAuthRouter := router.Group("")

	userApiRegister(unAuthRouter, authRouter)
	bannerApiRegister(unAuthRouter, authRouter)
	productApiRegister(unAuthRouter, authRouter)
	categoryApiRegister(unAuthRouter, authRouter)
	productDetailApiRegister(unAuthRouter, authRouter)
	cartApiRegister(unAuthRouter, authRouter)
	orderApiRegister(unAuthRouter, authRouter)
	managerApiRegister(unAuthRouter, authRouter)
	authApiRegister(unAuthRouter, authRouter)

	err := router.Run(fmt.Sprintf(":%d", conf.HTTPPort))
	if err != nil {
		fmt.Println("Start server failed!")
	}
}
