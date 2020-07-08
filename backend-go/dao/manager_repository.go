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