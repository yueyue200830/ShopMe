package service

import (
	"backend-go/dao"
	"backend-go/entity"
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