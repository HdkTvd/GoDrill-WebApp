package data

import uuid "github.com/satori/go.uuid"

type UUID interface{}

type User struct {
	UUID        UUID   `json:"uuid"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	IsActive    string `json:"isActive"`
}

var userList []*User

type Users []*User

func GetUserList() Users {
	return userList
}

func AddUser(u *User) {
	u.UUID = getNextID()
	userList = append(userList, u)
}

func getNextID() UUID {
	uid := uuid.NewV4()
	return uid
}
