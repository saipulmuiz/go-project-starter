package service

import (
	"github.com/saipulmuiz/go-project-starter/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (user *models.User, err error)
}

type CategoryRepository interface {
	GetCategories(req models.GetCategoryRequest) (*[]models.Category, int64, error)
	GetCategoryByID(productId int64) (*models.Category, error)
	CreateCategory(product *models.Category) (*models.Category, error)
	UpdateCategory(tx *gorm.DB, productId int64, productUpdate *models.Category) (*models.Category, error)
	DeleteCategory(productId int64) error
}
