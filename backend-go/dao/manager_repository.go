package dao

import "backend-go/entity"

var managerRepository *ManagerRepository

type ManagerRepository struct {
}

func init() {
	managerRepository = &ManagerRepository{}
}

func GetManagerRepository() *ManagerRepository {
	return managerRepository
}

func (m *ManagerRepository) GetManagerIDByNameAndPassword(manager *entity.Manager) error {
	return db.Where(&manager).First(&manager).Error
}

func (m *ManagerRepository) GetManagerInfoByID(id int) *entity.Manager {
	manager := &entity.Manager{}
	db.First(&manager, id)
	return manager
}

func (m *ManagerRepository) GetManagersByPage(page, size int) (managers []entity.Manager, err error) {
	offsetNum := (page - 1) * size
	dbManagers := db.Table("managers").Order("id desc").Offset(offsetNum).Limit(size)
	err = dbManagers.Select("id, name").Scan(&managers).Error
	return managers, err
}

func (m *ManagerRepository) GetManagerNumber() (number int) {
	db.Table("managers").Count(&number)
	return number
}

func (m *ManagerRepository) DeleteManagerByID(id int) error {
	return db.Where("id = ?", id).Delete(&entity.Manager{}).Error
}

func (m *ManagerRepository) UpdateManager(manager *entity.Manager) error {
	return db.Save(&manager).Error
}

func (m *ManagerRepository) CreateManager(manager *entity.Manager) error {
	return db.Create(&manager).Error
}
