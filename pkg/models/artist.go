package models

import "time"

type Aritst struct {
	BaseModel
	Name        string
	Description string
	Image       string
	Birtdate    time.Time
	Gender      bool
}
