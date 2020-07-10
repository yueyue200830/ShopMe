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
	router.GET("/banner/image/:banner", curd.getBannerImage)
	router.GET("/banners", curd.getBanners)
	router.DELETE("/banner", curd.deleteBanner)
	router.PUT("/banner", curd.updateBanner)
	router.POST("/banner", curd.createBanner)
	router.POST("/banner/image", curd.uploadBannerImage)
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

func (b *BannerController) uploadBannerImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileType := file.Header.Get("Content-Type")
	status := 0
	name := ""

	name, status = b.bannerService.GenerateRandomImageName(fileType)
	if status == 0 {
		err := c.SaveUploadedFile(file, "../images/banners/" + name)
		if err != nil {
			status = 1
		}
	}

	url := ""
	if status == 0 {
		url = "/banner/image/" + name
	}
	c.JSON(http.StatusOK, gin.H{
		"code": status,
		"url": url,
	})
}

func (b *BannerController) updateBanner(c *gin.Context) {
	status := 0
	var banner *entity.BannerProduct
	if err := c.ShouldBindJSON(&banner); err != nil {
		status = 1
	} else {
		status = b.bannerService.UpdateBanner(banner)
	}
	c.JSON(http.StatusOK, status)
}

func (b *BannerController) createBanner(c *gin.Context) {
	status := 0
	var banner *entity.BannerProduct
	if err := c.ShouldBindJSON(&banner); err != nil {
		status = 1
	} else {
		status = b.bannerService.CreateBanner(banner)
	}
	c.JSON(http.StatusOK, status)
}
