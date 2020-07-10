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

func (b *BannerRepository) GetBannerProductsByPage(page, size int) (list []entity.BannerProduct) {
	offsetNum := (page - 1) * size
	db.Table("banners").Select("banners.id as id, banner, product_id, price, title, image").Order("id desc").Offset(offsetNum).Limit(size).Joins("left join products on banners.product_id = products.id").Scan(&list)
	return list
}

func (b *BannerRepository) GetBannerNumber() (number int) {
	db.Table("banners").Count(&number)
	return number
}

func (b *BannerRepository) DeleteBannerByID(id int) error {
	return db.Where("id = ?", id).Delete(&entity.Banner{}).Error
}

func (b *BannerRepository) UpdateBanner(banner entity.Banner) error {
	return db.Save(&banner).Error
}

func (b *BannerRepository) CreateBanner(banner entity.Banner) error {
	return db.Create(&banner).Error
}