package validator

import (
	"fmt"

	"github.com/godrill1/data"
	"github.com/godrill1/handlers"
)

func Validator() {

	var isEligible []bool
	userNames := data.GetUserNames()

	for i := 0; i < len(userNames); i++ {
		if userNames[i] == "" {
			isEligible[i] = true
		}
	}

	for i := 0; i < len(userNames); i++ {
		if userNames[i] == "" {
			fmt.Println("A user does not have name")
		} else {
			handlers.WriteToJSON()
		}
	}
}
