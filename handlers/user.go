package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	l *logrus.Logger
}

func NewUserHandler(l *logrus.Logger) *UserHandler {
	return &UserHandler{l}
}

func (uh *UserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		userList := GetUserMap()
		e := json.NewEncoder(rw)
		err := e.Encode(userList)
		if err != nil {
			uh.l.Error("Unable to convert to json")
		}
		uh.l.Info("Users Get request")
	}

	if r.Method == http.MethodPost {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			uh.l.Error("Unable to convert to json")
		}
		ReadCsv(string(data), uh.l)
		uh.l.Info("Some user details were added!")
	}
}