package service

import (
	"backend-go/dao"
	"backend-go/entity"
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
		banners[i].Banner.Banner = "/banner/" + banners[i].Banner.Banner
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