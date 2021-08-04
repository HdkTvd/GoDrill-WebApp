package handlers

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/godrill1/models"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ReadCsvFile(filePath string, log *logrus.Logger, db *gorm.DB) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Error("Cannot open .csv file.")
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Error("Unable to read .csv file.")
	}

	for _, lines := range csvLines {

		phNo, _ := strconv.Atoi(lines[3])
		isAct, _ := strconv.ParseBool(lines[4])
		uid, _ := uuid.FromString(lines[0])

		user := &models.User{
			UUID:        uid,
			Name:        lines[1],
			Email:       lines[2],
			PhoneNumber: phNo,
			IsActive:    isAct,
		}

		if user.IsValid(log) {
			res, err := models.AddUser(user, db)
			if err != nil {
				log.Error(err)
			} else {
				log.Info("Rows Affected: ", res)
			}
		} else {
			log.Error("Unable to create User with UUID " + user.UUID.String())
		}
	}
}
