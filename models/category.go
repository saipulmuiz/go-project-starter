package models

import "time"

type Category struct {
	CategoryID   int64     `json:"category_id" db:"category_id"`
	CategoryName string    `json:"category_name" db:"category_name"`
	CreatedAt    time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type GetCategoryRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type GetCategoryResponse struct {
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name" validate:"required"`
}

type UpdateCategoryRequest struct {
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name" validate:"required"`
}

type CategoryResponse struct {
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type CategoryProduct struct {
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
	ProductID    int64  `json:"product_id"`
}
