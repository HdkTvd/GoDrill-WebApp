package main

import (
	"net/http"
	"os"
	"time"

	"github.com/godrill1/controller"
	"github.com/godrill1/services"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	//create a new logger
	log := logrus.New()
	log.SetOutput(os.Stdout)

	//connect to MYSQL database
	db := controller.ConnectToDB(log)

	//Interface for services
	var sv *services.ServiceImplementation

	uh := controller.NewUserController(sv, log, db)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/GET/users", uh.GetUsersController)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/POST/users", uh.AddUsersController)

	//create a new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
