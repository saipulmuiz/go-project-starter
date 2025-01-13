package service

import (
	"context"

	"github.com/saipulmuiz/go-project-starter/models"
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
)

type UserUsecase interface {
	Register(ctx context.Context, request models.RegisterUserRequest) (errx serror.SError)
	Login(ctx context.Context, request models.LoginUser) (res models.LoginResponse, errx serror.SError)
}

type CategoryUsecase interface {
	GetCategories(ctx context.Context, req models.GetCategoryRequest) (res []models.GetCategoryResponse, errx serror.SError)
	CreateCategory(ctx context.Context, req models.CreateCategoryRequest) (res *models.GetCategoryResponse, errx serror.SError)
	UpdateCategory(ctx context.Context, req models.UpdateCategoryRequest) (res *models.GetCategoryResponse, errx serror.SError)
	DeleteCategory(ctx context.Context, categoryId int64) (errx serror.SError)
}
