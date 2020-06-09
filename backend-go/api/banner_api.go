package api

import (
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BannerController struct {
	bannerService *service.BannerService
}

func bannerApiRegister(router *gin.Engine) {
	curd := BannerController{service.GetBannerService()}
	router.GET("/mainBanners", curd.getMainBanners)
	router.GET("/banner/:banner", curd.getBannerImage)
}

func (b *BannerController) getMainBanners(c *gin.Context) {
	banners := b.bannerService.GetMainBanners()
	c.JSON(http.StatusOK, banners)
}

func (b *BannerController) getBannerImage(c *gin.Context) {
	bannerPath := b.bannerService.GetBannerPath(c.Param("banner"))
	c.File(bannerPath)
}
