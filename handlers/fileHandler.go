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

func ReadCsv(filePath string, l *logrus.Logger, db *gorm.DB) {

	csvFile, err := os.Open(filePath)
	if err != nil {
		l.Error("Cannot open file.")
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		l.Error("Unable to read file.")
	}

	for _, lines := range csvLines {

		phNo, _ := strconv.Atoi(lines[3])
		// if errPh != nil {
		// 	l.Error("Unable to convert String type to Int type: ", err)
		// 	continue
		// }

		isAct, _ := strconv.ParseBool(lines[4])
		// if errAct != nil {
		// 	l.Error("Unable to Parse bool: ", err)
		// 	continue
		// }

		uid, _ := uuid.FromString(lines[0])
		// if errId != nil {
		// 	l.Error("Unable to convert UUID from string type: ", err)
		// 	continue
		// }

		user := &models.User{
			UUID:        uid,
			Name:        lines[1],
			Email:       lines[2],
			PhoneNumber: phNo,
			IsActive:    isAct,
		}

		err := user.IsValid()
		if err != nil {
			l.Error("User details invalid:", err)
		} else {
			res, err := models.AddUser(user, db)
			if err != nil {
				l.Error(err)
			} else {
				l.Info("Rows Affected:", res)
			}
		}

	}

}
