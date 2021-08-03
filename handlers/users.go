package handlers

import (
	"fmt"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber int       `json:"phone_number"`
	IsActive    bool      `json:"isActive"`
}

var userMap = make(map[string]User)

func GetUserMap() map[string]User {
	return userMap
}

func AddUser(u *User) {
	userMap[u.UUID.String()] = *u
	fmt.Println(*u)
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

	if _, isDuplicate := userMap[u.UUID.String()]; isDuplicate {
		return ErrDuplicateUUID
	}

	if len(strconv.Itoa(u.PhoneNumber)) != 10 {
		return ErrInvalidPhoneNumber
	}

	return nil
}
