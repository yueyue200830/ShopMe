package dao

import "backend-go/entity"

var bannerRepository *BannerRepository

type BannerRepository struct {
}

func init() {
	bannerRepository = &BannerRepository{}
}

func GetBannerRepository() *BannerRepository {
	return bannerRepository
}

func (b *BannerRepository) GetNewBanners(num int) []entity.Banner {
	var banners []entity.Banner
	db.Order("id desc").Limit(num).Find(&banners)
	return banners
}
