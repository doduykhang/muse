package models

import "time"

type Song struct {
	BaseModel
	Title       string
	Image       string
	Audio       string
	Duration    uint
	ReleaseDate time.Time
}
