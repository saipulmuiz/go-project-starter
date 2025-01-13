package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/saipulmuiz/go-project-starter/models"
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
	api "github.com/saipulmuiz/go-project-starter/service"
	"github.com/saipulmuiz/go-project-starter/service/helper"
)

type CategoryUsecase struct {
	categoryRepo api.CategoryRepository
}

func NewCategoryUsecase(
	categoryRepo api.CategoryRepository,
) api.CategoryUsecase {
	return &CategoryUsecase{
		categoryRepo: categoryRepo,
	}
}

func (u *CategoryUsecase) GetCategories(ctx context.Context, req models.GetCategoryRequest) (res []models.GetCategoryResponse, errx serror.SError) {
	categories, errx := u.categoryRepo.GetCategories(ctx, req)
	if errx != nil {
		errx.AddComments("[usecase][GetCategories] Failed to get categories")
		return
	}

	for _, category := range categories {
		res = append(res, models.GetCategoryResponse{
			CategoryID:   category.CategoryID,
			CategoryName: category.CategoryName,
			CreatedAt:    helper.ParseDateTime(helper.DATE_FORMAT_YYYY_MM_DD_TIME, category.CreatedAt),
			UpdatedAt:    helper.ParseDateTime(helper.DATE_FORMAT_YYYY_MM_DD_TIME, category.UpdatedAt),
		})
	}

	return
}

func (u *CategoryUsecase) CreateCategory(ctx context.Context, req models.CreateCategoryRequest) (res *models.GetCategoryResponse, errx serror.SError) {
	category := &models.Category{
		CategoryName: req.CategoryName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	categoryId, err := u.categoryRepo.CreateCategory(ctx, req)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][CreateCategory] Failed to create category")
		return
	}

	res = &models.GetCategoryResponse{
		CategoryID:   categoryId,
		CategoryName: category.CategoryName,
		CreatedAt:    helper.ParseDateTime(helper.DATE_FORMAT_YYYY_MM_DD_TIME, category.CreatedAt),
		UpdatedAt:    helper.ParseDateTime(helper.DATE_FORMAT_YYYY_MM_DD_TIME, category.UpdatedAt),
	}

	return
}

func (u *CategoryUsecase) UpdateCategory(ctx context.Context, req models.UpdateCategoryRequest) (res *models.GetCategoryResponse, errx serror.SError) {
	checkData, errx := u.categoryRepo.GetCategoryByID(ctx, req.CategoryID)
	if errx != nil {
		errx.AddComments("[usecase][UpdateCategory] Failed to get category")
		return
	}

	if checkData.CategoryID == 0 {
		errx = serror.Newi(http.StatusNotFound, "Category not found")
		return
	}

	categoryUpdate := models.UpdateCategoryRequest{
		CategoryName: req.CategoryName,
	}

	categoryUpdated, errx := u.categoryRepo.UpdateCategoryByID(ctx, nil, categoryUpdate)
	if errx != nil {
		errx.AddComments("[usecase][UpdateCategory] Failed to update category")
		return
	}

	res = &models.GetCategoryResponse{
		CategoryID:   categoryUpdated.CategoryID,
		CategoryName: categoryUpdated.CategoryName,
		CreatedAt:    helper.ParseDateTime(helper.DATE_FORMAT_YYYY_MM_DD_TIME, categoryUpdated.CreatedAt),
		UpdatedAt:    helper.ParseDateTime(helper.DATE_FORMAT_YYYY_MM_DD_TIME, categoryUpdated.UpdatedAt),
	}

	return
}

func (u *CategoryUsecase) DeleteCategory(ctx context.Context, categoryId int64) (errx serror.SError) {
	category, errx := u.categoryRepo.GetCategoryByID(ctx, categoryId)
	if errx != nil {
		errx.AddComments("[usecase][DeleteCategory] Failed to get category")
		return
	}

	if category.CategoryID == 0 {
		errx = serror.Newi(http.StatusNotFound, "Category not found")
		return
	}

	errx = u.categoryRepo.DeleteCategory(ctx, categoryId)
	if errx != nil {
		errx.AddComments("[usecase][DeleteCategory] Failed to delete category")
		return
	}

	return
}
