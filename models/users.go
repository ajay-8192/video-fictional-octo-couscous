package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string    `json:"firstName" gorm:"type:varchar(255);index:member"`
	LastName     string    `json:"lastName" gorm:"type:varchar(255);index:member"`
	Email        string    `json:"email" gorm:"type:varchar(255);unique;index"`
	Password     string    `json:"-"`
	Username     string    `json:"user" gorm:"type:varchar(255);unique;index"`
	Active       bool      `json:"active"`
	LastLoggedIn time.Time `json:"lastLoggedIn"`
}
