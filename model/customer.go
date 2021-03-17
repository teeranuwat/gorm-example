package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Serial    string `gorm:"not null;unique"`
	Firstname string `gorm:"not null"`
	Lastname  string `gorm:"not null"`
	ID        string `gorm:"not null"`
	Tel       string `gorm:"not null"`
	Email     string `gorm:"not null"`
}
