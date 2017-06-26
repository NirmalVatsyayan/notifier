package controllers

import (
	"gopkg.in/mgo.v2"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"../models"
	"log"
	"fmt"
)

type NotificationController struct {
	session *mgo.Session
}

func NewNotificationController(session *mgo.Session) *NotificationController {
	return &NotificationController{session}
}


/*
For OPTIONS (CORS) handling, not required when webservers on domain/ sub-domain will be used
 */
func (controller NotificationController) OptionsNotification(writer http.ResponseWriter, req *http.Request) {
	log.Println("Inside OptionsNotification")
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin","*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST")
	//must be equal to Access-Control-Request-Headers
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	writer.WriteHeader(http.StatusOK) // 200
}

func (controller NotificationController) PostNotification(writer http.ResponseWriter, req *http.Request) {

	// read request body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writer.WriteHeader(404)
		fmt.Fprintf(writer, "%s\n", "Error parsing request body")
		return
		//panic(err)
	}

	//unmarshall the form data
	var notification_obj models.Notification
	err = json.Unmarshal(body, &notification_obj)
	if err != nil {
		writer.WriteHeader(404)
		fmt.Fprintf(writer, "%s\n", "Error in request data")
		return
		//panic(err)
	}
	//log.Println(notification_obj)

	notification_json, _ := json.Marshal(notification_obj)

	if len(notification_obj.Header) < 20 || len(notification_obj.Header) > 150 {
		//fmt.Fprintf(writer, "%s\n", "Invalid header data length")
		writer.WriteHeader(404)
		fmt.Fprintf(writer, "%s\n", "Invalid header data length")
		return
	}

	if len(notification_obj.Payload) < 20 || len(notification_obj.Payload) > 300 {
		writer.WriteHeader(404)
		fmt.Fprintf(writer, "%s\n", "Invalid payload data length")
		return
	}

	// Insert notification object
	if err := controller.session.DB("notify").C("notification").Insert(notification_obj); err != nil {
		writer.WriteHeader(404)
		fmt.Fprintf(writer, "%s\n", "Error saving data in DB")
		return
	}


	// return the object inserted in response
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin","*")
	writer.WriteHeader(http.StatusOK) // 200
	writer.Write(notification_json)
}
