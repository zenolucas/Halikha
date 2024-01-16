package models

// User represents a user in the system
type User struct {
	ID       int64
	Username string
	Usertype string
	Password string
}