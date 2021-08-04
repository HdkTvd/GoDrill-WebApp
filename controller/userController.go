package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/godrill1/handlers"
	"github.com/godrill1/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserHandler struct {
	l  *logrus.Logger
	db *gorm.DB
}

func NewUserHandler(l *logrus.Logger, db *gorm.DB) *UserHandler {
	return &UserHandler{l, db}
}

func (uh *UserHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		//fetch all users from usertable in sql database
		uh.db.Find(&models.UserList)

		//convert table into json
		e := json.NewEncoder(rw)
		err := e.Encode(models.UserList)
		if err != nil {
			uh.l.Error("Unable to convert to json")
		}
		uh.l.Info("Users Get request")
		return
	} else {
		http.Error(rw, "Unable to fulfill request ", http.StatusBadRequest)
	}
}

func (uh *UserHandler) AddUsers(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			uh.l.Error("Unable to read POST data ", http.StatusBadRequest)
			return
		}
		handlers.ReadCsv(string(data), uh.l, uh.db)
		return
	} else {
		uh.l.Error(rw, "Unable to fulfill request ", http.StatusBadRequest)
	}

}

// func (uh *UserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		uh.GetUsers(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		uh.AddUsers(rw, r)
// 		return
// 	}
// }
