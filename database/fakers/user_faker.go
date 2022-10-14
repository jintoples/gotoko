package fakers

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/jintoples/gotoko/app/models"
	"gorm.io/gorm"
)

func UserFaker(*gorm.DB) *models.User {
	return &models.User{
		ID:            uuid.New().String(),
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		Email:         faker.Email(),
		Password:      "$2a$12$2iPRLW00R.hpkjfjjxa0N.LjR7ULyCDrw0jd7LjLRzZtAUOOYvPhy",
		RememberToken: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
}
