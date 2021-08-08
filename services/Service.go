package services

import (
	"encoding/json"

	"github.com/godrill1/models"
	"gorm.io/gorm"
)

type Service interface {
	GetUsers(db *gorm.DB) ([]byte, error)
	AddUsers(us *models.User) (int64, error)
}

type ServiceImplementation struct{}

func (si *ServiceImplementation) GetUsers(db *gorm.DB) ([]byte, error) {
	var userList []models.User
	db.Find(&userList)
	data, err := json.Marshal(userList)

	return data, err
}

func (si *ServiceImplementation) AddUser(us *models.User, db *gorm.DB) (int64, error) {
	result := db.Create(*us)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
