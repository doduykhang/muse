package models

type Account struct {
	BaseModel
	Email    string
	Password string
	Active   bool
	Banned   bool
	Role string
}
