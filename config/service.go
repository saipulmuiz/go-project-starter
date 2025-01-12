package config

import (
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
	"github.com/saipulmuiz/go-project-starter/service/handler/rest"
	"github.com/saipulmuiz/go-project-starter/service/repository"
	"github.com/saipulmuiz/go-project-starter/service/usecase"
)

func (cfg *Config) InitService() (errx serror.SError) {
	userRepo := repository.NewUserRepository(cfg.DB)
	userUsecase := usecase.NewUserUsecase(userRepo)

	categoryRepo := repository.NewCategoryRepo(cfg.DB)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	route := rest.CreateHandler(
		userUsecase,
		categoryUsecase,
	)

	cfg.Server = route

	return nil
}
