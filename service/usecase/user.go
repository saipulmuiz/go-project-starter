package usecase

import (
	"context"
	"net/http"

	"github.com/saipulmuiz/go-project-starter/models"
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
	api "github.com/saipulmuiz/go-project-starter/service"
	"github.com/saipulmuiz/go-project-starter/service/helper"
)

type UserUsecase struct {
	userRepo api.UserRepository
}

func NewUserUsecase(
	userRepo api.UserRepository,
) api.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Register(ctx context.Context, request models.RegisterUserRequest) (errx serror.SError) {
	userArgs := models.RegisterUserRequest{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	userCheck, err := u.userRepo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[usecase][Register] Failed to get user by email, [email: %s]", request.Email)
		return
	}

	if userCheck.UserID != 0 {
		errx = serror.Newi(http.StatusBadRequest, "Email already registered")
		return
	}

	_, err = u.userRepo.Register(ctx, userArgs)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[usecase][Register] Failed to register user, [email: %s]", request.Email)
		return
	}

	return
}

func (u *UserUsecase) Login(ctx context.Context, request models.LoginUser) (res models.LoginResponse, errx serror.SError) {
	userDB, errx := u.userRepo.GetUserByEmail(ctx, request.Email)
	if errx != nil {
		errx.AddCommentf("[usecase][Login] Failed to get user by email, [email: %s]", request.Email)
		return
	}

	if userDB.UserID == 0 {
		errx = serror.Newi(http.StatusNotFound, "User not found")
		return
	}

	accountMatch := helper.ComparePassword([]byte(userDB.Password), []byte(request.Password))
	if !accountMatch {
		errx = serror.Newi(http.StatusBadRequest, "Password does not match")
		return
	}

	token := helper.GenerateToken(userDB.UserID, userDB.Email, userDB.Name)

	res = models.LoginResponse{
		Token: token,
		User:  userDB,
	}

	return
}
