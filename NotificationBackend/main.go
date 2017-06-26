package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"./controllers"
)

func main() {
	//get db session for controller
	notification_controller := controllers.NewNotificationController(getSession())

	//route requests
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/notification", notification_controller.PostNotification).Methods("POST")
	router.HandleFunc("/notification", notification_controller.OptionsNotification).Methods("OPTIONS")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getSession() *mgo.Session {
	// Connect to local mongo
	session, err := mgo.Dial("mongodb://192.168.56.101")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return session
}
