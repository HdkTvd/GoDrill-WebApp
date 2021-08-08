package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/godrill1/handlers"
	"github.com/godrill1/services"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserController struct {
	sv  *services.ServiceImplementation
	log *logrus.Logger
	db  *gorm.DB
}

func NewUserController(sv *services.ServiceImplementation, log *logrus.Logger, db *gorm.DB) *UserController {
	return &UserController{sv, log, db}
}

func (uc *UserController) GetUsersController(rw http.ResponseWriter, r *http.Request) {

	userList, err := uc.sv.GetUsers(uc.db)
	if err != nil {
		uc.log.Error("Unable to convert to json")
		return
	} else {
		rw.Write(userList)
		uc.log.Info("Users Get request")
	}

	return
}

func (uc *UserController) AddUsersController(rw http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		uc.log.Error("Unable to read POST data ", http.StatusBadRequest)
		return
	}
	handlers.ReadCsvFile(string(data), uc.log, uc.db, uc.sv)
	return
}
