package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/saipulmuiz/go-project-starter/models"
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
	"github.com/saipulmuiz/go-project-starter/pkg/utils/utint"
	"github.com/saipulmuiz/go-project-starter/service/helper"
)

func (h *Handler) GetCategories(ctx *gin.Context) {
	var (
		errx serror.SError
	)

	page := utint.StringToInt(ctx.Query("page"), 1)
	limit := utint.StringToInt(ctx.Query("limit"), 10)

	data, errx := h.categoryUsecase.GetCategories(ctx, models.GetCategoryRequest{
		Page:  int(page),
		Limit: int(limit),
	})
	if errx != nil {
		handleError(ctx, errx.Code(), errx)
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Get categories successfully",
		Data:    data,
		Meta: map[string]interface{}{
			"total_data": 0,
		},
	})
}

func (h *Handler) CreateCategory(ctx *gin.Context) {
	var (
		request models.CreateCategoryRequest
		errx    serror.SError
	)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		errx = serror.NewFromErrori(http.StatusBadRequest, err)
		errx.AddComments("[handler][CreateCategory] while BodyJSONBind")
		handleError(ctx, errx.Code(), errx)
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		validationMessages := helper.BuildAndGetValidationMessage(err)
		handleValidationError(ctx, validationMessages)

		return
	}

	res, errx := h.categoryUsecase.CreateCategory(ctx, request)
	if errx != nil {
		handleError(ctx, errx.Code(), errx)
		return
	}

	ctx.JSON(http.StatusCreated, models.ResponseSuccess{
		Message: "Category successfully created",
		Data:    res,
	})
}

func (h *Handler) UpdateCategory(ctx *gin.Context) {
	var (
		request models.UpdateCategoryRequest
		errx    serror.SError
	)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		errx = serror.NewFromErrori(http.StatusBadRequest, err)
		errx.AddComments("[handler][UpdateCategory] while BodyJSONBind")
		handleError(ctx, errx.Code(), errx)
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		validationMessages := helper.BuildAndGetValidationMessage(err)
		handleValidationError(ctx, validationMessages)

		return
	}

	request.CategoryID = utint.StringToInt(ctx.Param("categoryId"), 0)
	res, errx := h.categoryUsecase.UpdateCategory(ctx, request)
	if errx != nil {
		handleError(ctx, errx.Code(), errx)
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Category successfully updated",
		Data:    res,
	})
}

func (h *Handler) DeleteCategory(ctx *gin.Context) {
	var errx serror.SError

	categoryId := utint.StringToInt(ctx.Param("categoryId"), 0)

	errx = h.categoryUsecase.DeleteCategory(ctx, categoryId)
	if errx != nil {
		handleError(ctx, errx.Code(), errx)
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Category successfully deleted",
	})
}
