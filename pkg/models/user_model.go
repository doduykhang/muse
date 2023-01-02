package models

import "time"

type User struct {
	BaseModel
	Firstname string
	Lastname  string
	Birthdate  time.Time
	Gender    bool
	AccountID uint
}
