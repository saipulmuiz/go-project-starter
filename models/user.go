package models

import (
	"time"

	"github.com/saipulmuiz/go-project-starter/service/helper"
	"gorm.io/gorm"
)

type User struct {
	UserID    int       `gorm:"not null;uniqueIndex;primaryKey;" json:"user_id"`
	Name      string    `gorm:"not null;size:256" json:"name"`
	Email     string    `gorm:"not null;" json:"email"`
	Password  string    `gorm:"not null;" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserInfo struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helper.HashPassword(u.Password)
	return
}