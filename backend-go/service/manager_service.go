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

func (m *ManagerService) GetManagerInfo(id int) (name string, status int){
	manager := m.managerRepository.GetManagerInfoByID(id)
	status = 0
	if name = manager.Name; name == "" {
		status = 1
	}
	return name, status
}