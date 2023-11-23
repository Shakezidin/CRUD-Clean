package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string `json:"username"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phonenumber"`
}
