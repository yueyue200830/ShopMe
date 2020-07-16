package service

import (
	"backend-go/dao"
	"backend-go/entity"
	"backend-go/utils"
)

var managerService *ManagerService

type ManagerService struct {
	managerRepository *dao.ManagerRepository
}

func init() {
	managerService = &ManagerService{dao.GetManagerRepository()}
}

func GetManagerService() *ManagerService {
	return managerService
}

func (m *ManagerService) Login(manager *entity.Manager) int {
	if err := m.managerRepository.GetManagerIDByNameAndPassword(manager); err != nil {
		return -1
	} else {
		return manager.ID
	}
}

func (m *ManagerService) GetManagerInfo(id int) (name string, status int) {
	manager := m.managerRepository.GetManagerInfoByID(id)
	status = 0
	if name = manager.Name; name == "" {
		status = 1
	}
	return name, status
}

func (m *ManagerService) GetManagers(page, size int) ([]entity.Manager, int, int) {
	status := 0
	users, err := m.managerRepository.GetManagersByPage(page, size)
	if err != nil {
		status = 1
	}
	num := m.managerRepository.GetManagerNumber()
	return users, num, status
}

func (m *ManagerService) DeleteManager(id int) (status int) {
	if id == 0 {
		return 1
	}
	if err := m.managerRepository.DeleteManagerByID(id); err != nil {
		return 1
	}
	return 0
}

func (m *ManagerService) UpdateManager(manager *entity.Manager) (status int) {
	if err := m.managerRepository.UpdateManager(manager); err != nil {
		return 1
	}
	return 0
}

func (m *ManagerService) CreateManager(manager *entity.Manager) (status int, password string) {
	status = 0
	password = ""
	if len(manager.Name) == 0 {
		status = 1
	} else {
		manager.Password = utils.GenerateRandomString(12)
		if err := m.managerRepository.CreateManager(manager); err != nil {
			status = 1
		} else {
			password = manager.Password
		}
	}
	return status, password
}

func (m *ManagerService) ResetPassword(manager *entity.Manager) (status int, password string) {
	status = 0
	password = ""
	if manager.ID == 0 {
		status = 1
	} else {
		manager.Password = utils.GenerateRandomString(12)
		if err := m.managerRepository.UpdateManager(manager); err != nil {
			status = 1
		} else {
			password = manager.Password
		}
	}
	return status, password
}
