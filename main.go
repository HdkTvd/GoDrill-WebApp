package main

import (
	"net/http"
	"os"
	"time"

	"github.com/godrill1/handlers"
	"github.com/sirupsen/logrus"
)

func main() {

	//create logger
	l := logrus.New()
	l.SetOutput(os.Stdout)

	//create handler
	uh := handlers.NewUserHandler(l)

	//create a new server mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/users", uh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	l.Fatal(s.ListenAndServe())
}
