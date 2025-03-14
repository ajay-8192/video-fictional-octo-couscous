package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName" gorm:"index:member"`
	LastName  string `json:"lastName" gorm:"index:member"`
	Email     string `json:"email" gorm:"unique;index"`
	Password  string `json:"-"`
	Username  string `json:"user" gorm:"unique;index"`
	Active    string `json:"active"`
}
