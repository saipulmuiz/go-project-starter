package usecase

import (
	"net/http"
	"time"

	"github.com/saipulmuiz/go-project-starter/models"
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
	api "github.com/saipulmuiz/go-project-starter/service"
	"github.com/saipulmuiz/go-project-starter/service/helper"
	"gorm.io/gorm"
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

func (u *CategoryUsecase) GetCategories(req models.GetCategoryRequest) (res []models.GetCategoryResponse, totalData int64, errx serror.SError) {
	var (
		err        error
		categories *[]models.Category
	)

	categories, totalData, err = u.categoryRepo.GetCategories(req)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][GetCategories] Failed to get categories")
		return
	}

	for _, category := range *categories {
		res = append(res, models.GetCategoryResponse{
			CategoryID:   category.CategoryID,
			CategoryName: category.CategoryName,
			CreatedAt:    helper.ParseDateTime(helper.DATE_FORMAT_YYYY_MM_DD_TIME, category.CreatedAt),
			UpdatedAt:    helper.ParseDateTime(helper.DATE_FORMAT_YYYY_MM_DD_TIME, category.UpdatedAt),
		})
	}

	return
}

func (u *CategoryUsecase) CreateCategory(request models.CreateCategoryRequest) (res *models.GetCategoryResponse, errx serror.SError) {
	category := &models.Category{
		CategoryName: request.CategoryName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	category, err := u.categoryRepo.CreateCategory(category)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][CreateCategory] Failed to create category")
		return
	}

	res = &models.GetCategoryResponse{
		CategoryID:   category.CategoryID,
		CategoryName: category.CategoryName,
		CreatedAt:    helper.ParseDateTime(helper.DATE_FORMAT_YYYY_MM_DD_TIME, category.CreatedAt),
		UpdatedAt:    helper.ParseDateTime(helper.DATE_FORMAT_YYYY_MM_DD_TIME, category.UpdatedAt),
	}

	return
}

func (u *CategoryUsecase) UpdateCategory(categoryId int64, request models.UpdateCategoryRequest) (res *models.GetCategoryResponse, errx serror.SError) {
	checkData, err := u.categoryRepo.GetCategoryByID(categoryId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errx = serror.Newi(http.StatusNotFound, "Category not found")
			return
		}

		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][UpdateCategory] Failed to get category")
		return
	}

	category := models.Category{
		CategoryName: request.CategoryName,
		UpdatedAt:    time.Now(),
	}

	categoryUpdated, err := u.categoryRepo.UpdateCategory(nil, checkData.CategoryID, &category)
	if err != nil {
		errx = serror.NewFromError(err)
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

func (u *CategoryUsecase) DeleteCategory(categoryId int64) (errx serror.SError) {
	_, err := u.categoryRepo.GetCategoryByID(categoryId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errx = serror.Newi(http.StatusNotFound, "Category not found")
			return
		}

		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][DeleteCategory] Failed to get category")
		return
	}

	err = u.categoryRepo.DeleteCategory(categoryId)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][DeleteCategory] Failed to delete category")
		return
	}

	return
}
