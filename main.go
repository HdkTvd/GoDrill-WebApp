package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/godrill1/handlers"
)

func getUser(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		userList := handlers.GetUserMap()
		e := json.NewEncoder(rw)
		err := e.Encode(userList)
		if err != nil {
			http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
			return
		}
		// fmt.Fprint(rw, string(data))
	}

	if r.Method == http.MethodPost {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "No Data found", http.StatusBadRequest)
			return
		}
		handlers.ReadCsv(string(data))
	}
}

func main() {

	// handlers.WriteFromCSVToStruct()
	// handlers.WriteFromStructToJSON()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Welcome!")
	})
	http.HandleFunc("/users", getUser)

	http.ListenAndServe(":9090", nil)

}
