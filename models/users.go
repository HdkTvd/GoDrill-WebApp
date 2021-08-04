package models

import (
	"strconv"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	UUID        uuid.UUID `json:"uuid" gorm:"primary_key"`
	Name        string    `json:"name" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	PhoneNumber int       `json:"phone_number"`
	IsActive    bool      `json:"isActive"`
}

var UserList []User

func AddUser(u *User, db *gorm.DB) (int64, error) {
	result := db.Create(*u)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (u *User) IsValid(log *logrus.Logger) bool {
	if u.UUID == uuid.Nil {
		log.Warn("UUID not provided.")
		log.Info("Creating a new User UUID...")
		u.UUID = uuid.NewV4()
	}
	if u.Name == "" {
		log.Error("Name field empty.")
		return false
	}
	if u.Email == "" {
		log.Error("Email field empty.")
		return false
	}
	if len(strconv.Itoa(u.PhoneNumber)) != 10 {
		log.Error("Invalid Phone Number.")
		return false
	}

	return true
}
