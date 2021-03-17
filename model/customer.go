package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Serial    string  `gorm:"not null;unique"`
	Firstname string  `gorm:"not null"`
	Lastname  string  `gorm:"not null"`
	ID        string  `gorm:"not null"`
	Tel       string  `gorm:"not null"`
	Email     string  `gorm:"not null"`
	Cards     []Card `gorm:"references:id"`
}

func (Customer) TableName() string {
	return "customer"
}

type Card struct {
	gorm.Model
	Serial     string `gorm:"not null;unique"`
	CustomerID int    `gorm:"fk"`
	Name       string `gorm:"name"`
	Detail     string `gorm:"detail"`
}

func (Card) TableName() string {
	return "card"
}
