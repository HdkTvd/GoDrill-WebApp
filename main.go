package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	csvFile, err := os.Open("./data/users.csv")
	if err != nil {
		fmt.Println("Opening file error", err)
		return
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("Error reading csv file", err)
	}
	defer csvFile.Close()

	//Validation and changes

	//

	for _, lines := range csvLines {

		users := user{
			UUID:        lines[0],
			Name:        lines[1],
			Email:       lines[2],
			PhoneNumber: lines[3],
			IsActive:    lines[4],
		}
		userList = append(userList, users)
	}

	data_json, err := json.Marshal(userList)
	if err != nil {
		fmt.Println("Error Marshaling csv data to json", err)
		return
	}

	err = ioutil.WriteFile("convData.json", data_json, 0644)
	if err != nil {
		fmt.Println("Error writing to json file", err)
	}

	fmt.Println("Contents of file in JSON:", string(data_json))

}
