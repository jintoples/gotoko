package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Addresses     []Address
	FirstName     string `gorm:"size:100;no null"`
	LastName      string `gorm:"size:100;no null"`
	Email         string `gorm:"size:100;no null;uniqueIndex"`
	Password      string `gorm:"size:255;no null"`
	RememberToken string `gorm:"size:255;no null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
