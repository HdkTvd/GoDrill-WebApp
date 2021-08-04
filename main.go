package main

import (
	"net/http"
	"os"
	"time"

	"github.com/godrill1/controller"
	"github.com/godrill1/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//create logger
	l := logrus.New()
	l.SetOutput(os.Stdout)

	//create database connection
	dsn := "root:Sequelp@ss@tcp(127.0.0.1:3306)/user_schema?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal("Error connecting database.")
	}
	db.AutoMigrate(&models.User{})

	//create handler
	uh := controller.NewUserHandler(l, db)

	//create a new server mux and register the handlers
	sm := http.NewServeMux()
	sm.HandleFunc("/users", uh.AddUsers)
	sm.HandleFunc("/", uh.GetUsers)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	l.Fatal(s.ListenAndServe())
}
