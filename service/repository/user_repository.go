package repository

import (
	"github.com/saipulmuiz/go-project-starter/models"
	api "github.com/saipulmuiz/go-project-starter/service"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) api.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Register(user *models.User) (*models.User, error) {
	return user, u.db.Create(&user).Error
}

func (u *userRepo) GetUserByEmail(email string) (user *models.User, err error) {
	return user, u.db.Where("email = ?", email).First(&user).Error
}