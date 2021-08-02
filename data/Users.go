package data

import (
	"fmt"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

type UUID interface{}

type User struct {
	UUID        UUID   `json:"uuid"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phone_number"`
	IsActive    bool   `json:"isActive"`
}

var userList []*User

type Users []*User

func GetUserList() Users {
	return userList
}

func GetUserNames() []string {
	var names []string

	for _, u := range userList {
		names = append(names, u.Name)
	}

	return names
}

func AddUser(u *User) {
	userList = append(userList, u)
}

func getNewID() UUID {
	id := uuid.NewV4()
	return id
}

func (u *User) IsValid() error {

	var ErrNameNotFound = fmt.Errorf("Name is required")
	var ErrEmailNotFound = fmt.Errorf("Email is required")
	var ErrInvalidPhoneNumber = fmt.Errorf("Invalid Phone Number")

	if u.UUID == "" {
		u.UUID = getNewID()
	}

	if u.Name == "" {
		return ErrNameNotFound
	}

	if u.Email == "" {
		return ErrEmailNotFound
	}

	if len(strconv.Itoa(u.PhoneNumber)) != 10 {
		return ErrInvalidPhoneNumber
	}

	return nil
}
