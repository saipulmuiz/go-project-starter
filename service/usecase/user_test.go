package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/saipulmuiz/go-project-starter/models"
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
	"github.com/saipulmuiz/go-project-starter/service/helper"
	"github.com/saipulmuiz/go-project-starter/service/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_UserUsecase_Register(t *testing.T) {
	type testCase struct {
		name             string
		wantError        bool
		expectedResponse serror.SError
		request          models.RegisterUserRequest
		onGetUserByEmail func(mock *mocks.MockUserRepository)
		onRegister       func(mock *mocks.MockUserRepository)
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:      "register user success",
		wantError: false,
		request: models.RegisterUserRequest{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password123",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail(gomock.Any(), "john@example.com").Return(models.User{}, nil)
		},
		onRegister: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().Register(gomock.Any(), gomock.Any()).Return(int64(1), nil)
		},
		expectedResponse: nil,
	})

	testTable = append(testTable, testCase{
		name:      "email already registered",
		wantError: true,
		request: models.RegisterUserRequest{
			Name:     "Jane Doe",
			Email:    "jane@example.com",
			Password: "password123",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail(gomock.Any(), "jane@example.com").Return(models.User{UserID: 1}, nil)
		},
		expectedResponse: serror.New("Email already registered"),
	})

	testTable = append(testTable, testCase{
		name:      "error checking user by email",
		wantError: true,
		request: models.RegisterUserRequest{
			Name:     "Error User",
			Email:    "error@example.com",
			Password: "password123",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail(gomock.Any(), "error@example.com").Return(models.User{}, serror.New("database error"))
		},
		expectedResponse: serror.New("database error"),
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			userRepo := mocks.NewMockUserRepository(mockCtrl)

			if tc.onGetUserByEmail != nil {
				tc.onGetUserByEmail(userRepo)
			}

			if tc.onRegister != nil {
				tc.onRegister(userRepo)
			}

			usecase := &UserUsecase{userRepo: userRepo}

			err := usecase.Register(context.Background(), tc.request)

			if tc.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedResponse, err)
			}
		})
	}
}

func Test_UserUsecase_Login(t *testing.T) {
	type testCase struct {
		name             string
		wantError        bool
		expectedResponse models.LoginResponse
		request          models.LoginUser
		onGetUserByEmail func(mock *mocks.MockUserRepository)
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:      "user not found",
		wantError: true,
		request: models.LoginUser{
			Email:    "john@example.com",
			Password: "password",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail(gomock.Any(), "john@example.com").Return(models.User{}, nil)
		},
	})

	testTable = append(testTable, testCase{
		name:      "password does not match",
		wantError: true,
		request: models.LoginUser{
			Email:    "john@example.com",
			Password: "wrongpassword",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mockUser := models.User{
				UserID:   1,
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: helper.HashPassword("password"),
			}
			mock.EXPECT().GetUserByEmail(gomock.Any(), "john@example.com").Return(mockUser, nil)
		},
	})

	testTable = append(testTable, testCase{
		name:      "error checking user by email",
		wantError: true,
		request: models.LoginUser{
			Email:    "error@example.com",
			Password: "password",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail(gomock.Any(), "error@example.com").Return(models.User{}, serror.New("database error"))
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			userRepo := mocks.NewMockUserRepository(mockCtrl)

			if tc.onGetUserByEmail != nil {
				tc.onGetUserByEmail(userRepo)
			}

			usecase := &UserUsecase{userRepo: userRepo}

			resp, err := usecase.Login(context.Background(), tc.request)

			if tc.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}
