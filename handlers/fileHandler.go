package handlers

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

func ReadCsv(filePath string) {

	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Opening file error", err)
		return
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("Error reading csv file", err)
		return
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
			fmt.Println("The user details are invalid:", err)
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
