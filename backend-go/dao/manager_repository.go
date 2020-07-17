package dao

import (
	"backend-go/entity"
	"errors"
)

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

func (m *ManagerRepository) UpdateManagerPassword(id int, oldPassword, newPassword string) error {
	tx := db.Begin()

	manager := &entity.Manager{ID: id}
	if err := tx.First(&manager).Error; err != nil {
		tx.Rollback()
		return err
	}
	if manager.Password != oldPassword {
		tx.Rollback()
		return errors.New("wrong password")
	}
	if err := tx.Model(&manager).Update("password", newPassword).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (m *ManagerRepository) CreateManager(manager *entity.Manager) error {
	return db.Create(&manager).Error
}

func (m *ManagerRepository) GetManagerIDByName(name string) int {
	var manager entity.Manager
	db.Select("id").Where("name = ?", name).First(&manager)
	return manager.ID
}
