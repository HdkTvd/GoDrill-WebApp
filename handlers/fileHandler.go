package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/godrill1/data"
)

func ReadFromCSV() {

	csvFile, err := os.Open("./data/usersList.csv")
	if err != nil {
		fmt.Println("Opening file error", err)
		return
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("Error reading csv file", err)
		return
	}

	defer csvFile.Close()

	for _, lines := range csvLines {
		user := &data.User{
			UUID:        lines[0],
			Name:        lines[1],
			Email:       lines[2],
			PhoneNumber: lines[3],
			IsActive:    lines[4],
		}
		err := user.IsValid()
		if err != nil {
			fmt.Println("The user details are invalid", err)
		} else {
			data.AddUser(user)
		}
	}

}

func WriteToJSON() {
	userList := data.GetUserList()
	data_json, err := json.Marshal(userList)

	if err != nil {
		fmt.Println("Error Marshaling csv data to json", err)
		return
	}
	err = ioutil.WriteFile("data/convData.json", data_json, 0644)

	if err != nil {
		fmt.Println("Error writing to json file", err)
		return
	}
}
