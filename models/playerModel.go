package models

import "github.com/jinzhu/gorm"

// User represents the user model
type User struct {
    gorm.Model
    ID    string `gorm:"primary_key" json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
