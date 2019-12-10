package models

// User model
type User struct {
	Model
	Username  string
	Email     string
	Password  string
	Firstname string
	Lastname  string
}
