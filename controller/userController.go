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

type UserController struct {
	log *logrus.Logger
	db  *gorm.DB
}

func NewUserController(log *logrus.Logger, db *gorm.DB) *UserController {
	return &UserController{log, db}
}

func (uc *UserController) GetUsers(rw http.ResponseWriter, r *http.Request) {
	uc.db.Find(&models.UserList)
	e := json.NewEncoder(rw)
	err := e.Encode(models.UserList)
	if err != nil {
		uc.log.Error("Unable to convert to json")
	}
	uc.log.Info("Users Get request")
	return
}

func (uc *UserController) AddUsers(rw http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		uc.log.Error("Unable to read POST data ", http.StatusBadRequest)
		return
	}
	handlers.ReadCsvFile(string(data), uc.log, uc.db)
	return
}
