package data

type User struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	IsActive    string `json:"isActive"`
}

var userList []*User

type Users []*User

func GetUser() Users {
	return userList
}
