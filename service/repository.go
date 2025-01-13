package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/saipulmuiz/go-project-starter/models"
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
)

type UserRepository interface {
	Register(ctx context.Context, req models.RegisterUserRequest) (userId int64, errx serror.SError)
	GetUserByID(ctx context.Context, userID string) (res models.User, errx serror.SError)
	GetUserByEmail(ctx context.Context, email string) (res models.User, errx serror.SError)
}

type CategoryRepository interface {
	CreateCategory(ctx context.Context, req models.CreateCategoryRequest) (categoryId int64, errx serror.SError)
	GetCategories(ctx context.Context, req models.GetCategoryRequest) (res []models.Category, errx serror.SError)
	GetCategoryByID(ctx context.Context, categoryId int64) (res models.Category, errx serror.SError)
	UpdateCategoryByID(ctx context.Context, tx *sqlx.DB, req models.UpdateCategoryRequest) (res models.Category, errx serror.SError)
	DeleteCategory(ctx context.Context, categoryId int64) (errx serror.SError)
}
