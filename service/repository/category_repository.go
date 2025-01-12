package repository

import (
	"github.com/saipulmuiz/go-project-starter/models"
	api "github.com/saipulmuiz/go-project-starter/service"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) api.CategoryRepository {
	return &categoryRepo{db}
}

func (u *categoryRepo) GetCategories(req models.GetCategoryRequest) (*[]models.Category, int64, error) {
	var (
		categories []models.Category
		count      int64
	)

	offset := (req.Page - 1) * req.Limit

	err := u.db.
		Order("created_at DESC").
		Offset(offset).
		Limit(req.Limit).
		Find(&categories).Error

	if err != nil {
		return nil, count, err
	}

	err = u.db.
		Model(&models.Category{}).
		Count(&count).Error

	return &categories, count, err
}

func (u *categoryRepo) GetCategoryByID(categoryId int64) (*models.Category, error) {
	var category models.Category
	err := u.db.Where("category_id = ?", categoryId).First(&category).Error
	return &category, err
}

func (u *categoryRepo) CreateCategory(category *models.Category) (*models.Category, error) {
	return category, u.db.Create(&category).Error
}

func (u *categoryRepo) UpdateCategory(tx *gorm.DB, categoryId int64, categoryUpdate *models.Category) (*models.Category, error) {
	var (
		category models.Category
		result   *gorm.DB
	)
	if tx != nil {
		result = tx.Model(&category).Clauses(clause.Returning{}).Where("category_id", categoryId).Updates(categoryUpdate)
	} else {
		result = u.db.Model(&category).Clauses(clause.Returning{}).Where("category_id", categoryId).Updates(categoryUpdate)
	}

	return &category, result.Error
}

func (u *categoryRepo) DeleteCategory(categoryId int64) error {
	var category models.Category
	result := u.db.Model(&category).Where("category_id", categoryId).Delete(categoryId)
	return result.Error
}
