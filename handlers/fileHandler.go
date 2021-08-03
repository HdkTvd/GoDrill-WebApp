package handlers

import (
	"encoding/csv"
	"os"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func ReadCsv(filePath string, l *logrus.Logger) {

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
		isAct, _ := strconv.ParseBool(lines[4])
		uuid, _ := uuid.FromString(lines[0])

		user := &User{
			UUID:        uuid,
			Name:        lines[1],
			Email:       lines[2],
			PhoneNumber: phNo,
			IsActive:    isAct,
		}

		err := user.IsValid()
		if err != nil {
			l.Error(err)
		} else {
			AddUser(user)
		}
	}

}

// func WriteFromStructToJSON() {

// 	userList := GetUserMap()

// 	data_json, err := json.Marshal(userList)
// 	if err != nil {
// 		fmt.Println("Error Marshaling csv data to json", err)
// 		return
// 	}

// 	err = ioutil.WriteFile("data/convData.json", data_json, 0644)
// 	if err != nil {
// 		fmt.Println("Error writing to json file", err)
// 		return
// 	}
// }
