package usecase

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/saipulmuiz/go-project-starter/models"
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
	"github.com/saipulmuiz/go-project-starter/service/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_CategoryUsecase_GetCategories(t *testing.T) {
	type testCase struct {
		name             string
		wantError        bool
		page, size       int
		onGetCategories  func(mock *mocks.MockCategoryRepository)
		expectedResponse []models.GetCategoryResponse
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:      "get categories success",
		wantError: false,
		page:      1,
		size:      10,
		onGetCategories: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().GetCategories(gomock.Any(), models.GetCategoryRequest{
				Page:  1,
				Limit: 10,
			}).Return([]models.Category{
				{CategoryID: 1, CategoryName: "Category 1"},
			}, nil)
		},
		expectedResponse: []models.GetCategoryResponse{
			{
				CategoryID:   1,
				CategoryName: "Category 1",
				CreatedAt:    "0001-01-01 00:00:00",
				UpdatedAt:    "0001-01-01 00:00:00",
			},
		},
	})

	testTable = append(testTable, testCase{
		name:      "error getting categories",
		wantError: true,
		page:      1,
		size:      10,
		onGetCategories: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().GetCategories(gomock.Any(), models.GetCategoryRequest{
				Page:  1,
				Limit: 10,
			}).Return(nil, serror.New("some error"))
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			categoryRepo := mocks.NewMockCategoryRepository(mockCtrl)

			if tc.onGetCategories != nil {
				tc.onGetCategories(categoryRepo)
			}

			service := &CategoryUsecase{
				categoryRepo: categoryRepo,
			}

			resp, err := service.GetCategories(context.Background(), models.GetCategoryRequest{
				Page:  tc.page,
				Limit: tc.size,
			})

			if tc.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}

func Test_CategoryUsecase_CreateCategory(t *testing.T) {
	type testCase struct {
		name             string
		request          models.CreateCategoryRequest
		wantError        bool
		onCreateCategory func(mock *mocks.MockCategoryRepository)
		expectedResponse *models.GetCategoryResponse
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name: "success create new category",
		request: models.CreateCategoryRequest{
			CategoryName: "New Category",
		},
		wantError: false,
		onCreateCategory: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(int64(1), nil)
		},
		expectedResponse: &models.GetCategoryResponse{
			CategoryID:   1,
			CategoryName: "New Category",
			CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		},
	})

	testTable = append(testTable, testCase{
		name: "error creating category",
		request: models.CreateCategoryRequest{
			CategoryName: "New Category",
		},
		wantError: true,
		onCreateCategory: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(int64(0), serror.New("some error"))
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			categoryRepo := mocks.NewMockCategoryRepository(mockCtrl)

			if tc.onCreateCategory != nil {
				tc.onCreateCategory(categoryRepo)
			}

			service := &CategoryUsecase{
				categoryRepo: categoryRepo,
			}

			resp, err := service.CreateCategory(context.Background(), tc.request)

			if tc.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}

func Test_CategoryUsecase_UpdateCategory(t *testing.T) {
	type testCase struct {
		name              string
		categoryId        int64
		request           models.UpdateCategoryRequest
		wantError         bool
		onGetCategoryByID func(mock *mocks.MockCategoryRepository)
		onUpdateCategory  func(mock *mocks.MockCategoryRepository)
		expectedResponse  *models.GetCategoryResponse
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:       "success update category",
		categoryId: 1,
		request: models.UpdateCategoryRequest{
			CategoryName: "Updated Category",
		},
		wantError: false,
		onGetCategoryByID: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(models.Category{CategoryID: 1}, nil)
		},
		onUpdateCategory: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().UpdateCategoryByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Category{CategoryID: 1, CategoryName: "Updated Category"}, nil)
		},
		expectedResponse: &models.GetCategoryResponse{
			CategoryID:   int64(1),
			CategoryName: "Updated Category",
			CreatedAt:    "0001-01-01 00:00:00",
			UpdatedAt:    "0001-01-01 00:00:00",
		},
	})

	testTable = append(testTable, testCase{
		name:       "category not found",
		categoryId: 1,
		request: models.UpdateCategoryRequest{
			CategoryName: "Updated Category",
		},
		wantError: true,
		onGetCategoryByID: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(models.Category{CategoryID: 0}, nil)
		},
	})

	testTable = append(testTable, testCase{
		name:       "error getting category",
		categoryId: 1,
		request: models.UpdateCategoryRequest{
			CategoryName: "Updated Category",
		},
		wantError: true,
		onGetCategoryByID: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().GetCategoryByID(gomock.Any(), gomock.Any()).Return(models.Category{}, serror.New("some error"))
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			categoryRepo := mocks.NewMockCategoryRepository(mockCtrl)

			if tc.onGetCategoryByID != nil {
				tc.onGetCategoryByID(categoryRepo)
			}
			if tc.onUpdateCategory != nil {
				tc.onUpdateCategory(categoryRepo)
			}

			service := &CategoryUsecase{
				categoryRepo: categoryRepo,
			}

			resp, err := service.UpdateCategory(context.Background(), tc.request)

			if tc.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}

func Test_CategoryUsecase_DeleteCategory(t *testing.T) {
	type testCase struct {
		name              string
		categoryId        int64
		wantError         bool
		onGetCategoryByID func(mock *mocks.MockCategoryRepository)
		onDeleteCategory  func(mock *mocks.MockCategoryRepository)
		expectedError     serror.SError
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:       "success delete category",
		categoryId: 1,
		wantError:  false,
		onGetCategoryByID: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().GetCategoryByID(gomock.Any(), int64(1)).Return(models.Category{CategoryID: int64(1)}, nil)
		},
		onDeleteCategory: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().DeleteCategory(gomock.Any(), int64(1)).Return(nil)
		},
	})

	testTable = append(testTable, testCase{
		name:       "category not found",
		categoryId: 1,
		wantError:  true,
		onGetCategoryByID: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().GetCategoryByID(gomock.Any(), int64(1)).Return(models.Category{}, nil)
		},
		expectedError: serror.Newi(http.StatusNotFound, "Category not found"),
	})

	testTable = append(testTable, testCase{
		name:       "error getting category",
		categoryId: 1,
		wantError:  true,
		onGetCategoryByID: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().GetCategoryByID(gomock.Any(), int64(1)).Return(models.Category{}, serror.New("some error"))
		},
		expectedError: serror.NewFromError(serror.New("some error")),
	})

	testTable = append(testTable, testCase{
		name:       "error deleting category",
		categoryId: 1,
		wantError:  true,
		onGetCategoryByID: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().GetCategoryByID(gomock.Any(), int64(1)).Return(models.Category{CategoryID: int64(1)}, nil)
		},
		onDeleteCategory: func(mock *mocks.MockCategoryRepository) {
			mock.EXPECT().DeleteCategory(gomock.Any(), int64(1)).Return(serror.New("some error"))
		},
		expectedError: serror.NewFromError(serror.New("some error")),
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			categoryRepo := mocks.NewMockCategoryRepository(mockCtrl)

			if tc.onGetCategoryByID != nil {
				tc.onGetCategoryByID(categoryRepo)
			}
			if tc.onDeleteCategory != nil {
				tc.onDeleteCategory(categoryRepo)
			}

			service := &CategoryUsecase{
				categoryRepo: categoryRepo,
			}

			err := service.DeleteCategory(context.Background(), tc.categoryId)

			if tc.wantError {
				assert.NotNil(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
