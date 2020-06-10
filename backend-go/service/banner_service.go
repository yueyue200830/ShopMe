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
