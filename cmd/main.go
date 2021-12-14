package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lyx0/dank/handlers"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

func main() {

	bh := handlers.CheckMessage(log)

	sm := mux.NewRouter()
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/message", bh)

}
