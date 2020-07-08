package api

import (
	"backend-go/entity"
	"backend-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BannerController struct {
	bannerService *service.BannerService
}

func bannerApiRegister(router *gin.Engine) {
	curd := BannerController{service.GetBannerService()}
	router.GET("/mainBanners", curd.getMainBanners)
	router.GET("/banner/:banner", curd.getBannerImage)
	router.GET("/banners", curd.getBanners)
	router.DELETE("/banner", curd.deleteBanner)
}

func (b *BannerController) getMainBanners(c *gin.Context) {
	banners := b.bannerService.GetMainBanners()
	c.JSON(http.StatusOK, banners)
}

func (b *BannerController) getBannerImage(c *gin.Context) {
	bannerPath := b.bannerService.GetBannerPath(c.Param("banner"))
	c.File(bannerPath)
}

func (b *BannerController) getBanners(c *gin.Context) {
	status := 0
	var list []entity.BannerProduct
	number := 0
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		status = 1
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil || size < 1 {
		status = 1
	}
	if status == 0 {
		list, number = b.bannerService.GetBanners(page, size)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"data": gin.H{
			"items": list,
			"total": number,
		},
	})
}

func (b *BannerController) deleteBanner(c *gin.Context) {
	status := 0
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		status = 1
	} else {
		status = b.bannerService.DeleteBanner(id)
	}

	c.JSON(http.StatusOK, status)
}
