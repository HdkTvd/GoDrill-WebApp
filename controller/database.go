package controller

import (
	"github.com/godrill1/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	Username = "root"
	Password = "Sequelp@ss"
	Hostname = "127.0.0.1"
	Port     = "3306"
	Schema   = "user_schema"
)

func ConnectToDB(log *logrus.Logger) *gorm.DB {
	dsn := Username + ":" + Password + "@tcp(" + Hostname + ":" + Port + ")/" + Schema + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database.")
		return nil
	} else {
		log.Info("Connected to Database.")
		createTable(log, db)
		return db
	}
}

func createTable(log *logrus.Logger, db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Error("[ERROR] Creating User Table", err)
	}
}
