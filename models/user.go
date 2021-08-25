package models

// User developer register and signin
type User struct {
	DevID       int64  `json:"dev_id"`
	DevEmail    string `json:"dev_email"`
	DevPassword string `json:"dev_password"`
}
