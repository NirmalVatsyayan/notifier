package models

import "time"

type Notification struct {
	Header string `json:"header"`
	Payload string `json:"payload"`
	ImageUrl string `json:"imageUrl"`
	UserQuery string `json:"userQuery"`
	Notification_time time.Time `json:"notification_time"`
}
