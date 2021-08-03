package main

import (
	"net/http"
	"os"
	"time"

	"github.com/godrill1/handlers"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//create logger
	l := logrus.New()
	l.SetOutput(os.Stdout)

	//create database connection
	dsn := "user:password@tcp(127.0.0.1:3306)/dbase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal("Error connecting database.")
	}
	db.AutoMigrate(&handlers.User{})

	//create handler
	uh := handlers.NewUserHandler(l, db)

	//create a new server mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/users", uh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	l.Fatal(s.ListenAndServe())
}
