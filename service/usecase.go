package service

import (
	"github.com/saipulmuiz/go-project-starter/models"
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
)

type UserUsecase interface {
	Register(request *models.RegisterUser) (user *models.User, errx serror.SError)
	Login(request *models.LoginUser) (res models.LoginResponse, errx serror.SError)
}

type CategoryUsecase interface {
	GetCategories(req models.GetCategoryRequest) (res []models.GetCategoryResponse, totalData int64, errx serror.SError)
	CreateCategory(request models.CreateCategoryRequest) (res *models.GetCategoryResponse, errx serror.SError)
	UpdateCategory(productId int64, request models.UpdateCategoryRequest) (res *models.GetCategoryResponse, errx serror.SError)
	DeleteCategory(productId int64) (errx serror.SError)
}
