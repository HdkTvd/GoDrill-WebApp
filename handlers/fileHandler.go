package handlers

import (
	"encoding/csv"
	"os"
	"strconv"

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

		phNo, err := strconv.Atoi(lines[3])
		if err != nil {
			l.Error("Unable to convert String type to Int type: ", err)
			continue
		}

		isAct, err := strconv.ParseBool(lines[4])
		if err != nil {
			l.Error("Unable to Parse bool: ", err)
			continue
		}

		uuid, err := uuid.FromString(lines[0])
		if err != nil {
			l.Error("Unable to convert UUID from string type: ", err)
			continue
		}

		user := &User{
			UUID:        uuid,
			Name:        lines[1],
			Email:       lines[2],
			PhoneNumber: phNo,
			IsActive:    isAct,
		}

		err = user.IsValid()
		if err != nil {
			l.Error(err)
		} else {
			res, err := AddUser(user, db)
			if err != nil {
				l.Error(err)
			} else {
				l.Info("Rows Affected:", res)
			}

		}
	}

}
