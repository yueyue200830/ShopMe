package dao

type Banner struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Banner    string `json:"bannerPath"`
	ProductID int    `json:"productID"`
}

var bannerRepository *BannerRepository

type BannerRepository struct {
}

func init() {
	bannerRepository = &BannerRepository{}
}

func GetBannerRepository() *BannerRepository {
	return bannerRepository
}

func (b *BannerRepository) GetNewBanners(num int) []Banner {
	var banners []Banner
	db.Order("id desc").Limit(num).Find(&banners)
	return banners
}
