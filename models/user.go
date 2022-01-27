package models

import (
	"github.com/jinzhu/gorm"
)

//User struct declaration
type User struct {
	gorm.Model

	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Gender   string `json:"Gender"`
	Age      int
	Password string `json:"Password"`
	Symptoms []Symptom
}

type Symptom struct {
	gorm.Model

	Notes  string
	UserID uint
}
