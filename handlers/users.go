package handlers

import (
	"fmt"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	UUID        uuid.UUID `json:"uuid" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"not null"`
	Email       string    `json:"email" gorm:"not null"`
	PhoneNumber int       `json:"phone_number" gorm:"size:10"`
	IsActive    bool      `json:"isActive" gorm:"default:false"`
}

var userMap = make(map[uuid.UUID]User)

func GetUsers(db *gorm.DB) *gorm.DB {
	result := db.Find(&User{})
	return result
}

func AddUser(u *User, db *gorm.DB) (int64, error) {
	userMap[u.UUID] = *u
	result := db.Create(*u)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (u *User) IsValid() error {

	var ErrNameNotFound = fmt.Errorf("Name is required")
	var ErrEmailNotFound = fmt.Errorf("Email is required")
	var ErrInvalidPhoneNumber = fmt.Errorf("Invalid Phone Number")
	var ErrDuplicateUUID = fmt.Errorf("Duplicate UUID found")

	if u.UUID == uuid.Nil {
		u.UUID = uuid.NewV4()
	}

	if u.Name == "" {
		return ErrNameNotFound
	}

	if u.Email == "" {
		return ErrEmailNotFound
	}

	if _, isDuplicate := userMap[u.UUID]; isDuplicate {
		return ErrDuplicateUUID
	}

	if len(strconv.Itoa(u.PhoneNumber)) != 10 {
		return ErrInvalidPhoneNumber
	}

	return nil
}
