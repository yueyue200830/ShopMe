package service

import (
	"backend-go/dao"
	"backend-go/entity"
	"backend-go/utils"
	"strings"
)

var bannerService *BannerService

type BannerService struct {
	bannerRepository *dao.BannerRepository
}

func init() {
	bannerService = &BannerService{dao.GetBannerRepository()}
}

func GetBannerService() *BannerService {
	return bannerService
}

func (b *BannerService) GetMainBanners() []entity.Banner {
	return b.bannerRepository.GetNewBanners(4)
}

func (b *BannerService) GetBannerPath(name string) string {
	return "../images/banners/" + name
}

func (b *BannerService) GetBanners(page, size int) ([]entity.BannerProduct, int) {
	banners := b.bannerRepository.GetBannerProductsByPage(page, size)
	number := b.bannerRepository.GetBannerNumber()
	for i, _ := range banners {
		banners[i].Banner.Banner = "/banner/image/" + banners[i].Banner.Banner
		banners[i].Image = "/productImage/" + banners[i].Image
	}
	return banners, number
}

func (b *BannerService) DeleteBanner(id int) (status int) {
	err := b.bannerRepository.DeleteBannerByID(id)
	if err == nil {
		return 0
	}
	return 1
}

func (b *BannerService) GenerateRandomImageName(fileType string) (name string, status int) {
	status = 0
	s := strings.Split(fileType, "/")

	if len(s) != 2 || s[0] != "image" {
		status = 1
	} else {
		name = utils.GenerateRandomString(30)
		name = name + "." + s[1]
	}

	return name, status
}

func (b *BannerService) UpdateBanner(bannerProduct *entity.BannerProduct) (status int) {
	banner := bannerProduct.Banner
	// get banner name
	name := strings.Split(banner.Banner, "/")
	banner.Banner = name[len(name) - 1]

	err := b.bannerRepository.UpdateBanner(banner)
	if err != nil {
		return 1
	}
	return 0
}

func (b *BannerService) CreateBanner(bannerProduct *entity.BannerProduct) (status int) {
	banner := bannerProduct.Banner
	// get banner name
	name := strings.Split(banner.Banner, "/")
	banner.Banner = name[len(name) - 1]

	err := b.bannerRepository.CreateBanner(banner)
	if err != nil {
		return 1
	}
	return 0
}