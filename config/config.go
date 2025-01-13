package config

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
)

type Config struct {
	DB     *sqlx.DB
	Server *gin.Engine
}

func Init() (cfg Config) {
	Catch(cfg.InitTimezone())
	Catch(cfg.InitPostgres())
	Catch(cfg.InitService())

	return
}

func (cfg *Config) Start() (errx serror.SError) {
	cfg.Server.Run(os.Getenv("APP_PORT"))

	return
}

// func (cfg *Config) stop() {
// 	if cfg.APM != nil {
// 		cfg.APM.Shutdown()
// 	}
// }

func Catch(errx serror.SError) {
	if errx != nil {
		errx.Panic()
	}
}
