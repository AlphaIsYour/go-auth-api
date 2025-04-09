package models

type User struct {
	Username string
	Password string // hashed password
}

// Fake database pakai map
var Users = map[string]User{}
