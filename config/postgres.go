package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
	"github.com/saipulmuiz/go-project-starter/pkg/utils/utint"
	"github.com/saipulmuiz/go-project-starter/pkg/utils/utstring"
)

// func (cfg *Config) InitPostgres() serror.SError {
// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname =%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

// 	if err != nil {
// 		log.Fatalf("failed connect to database %+v", err)
// 		return serror.NewFromError(err)
// 	}

// 	err = db.Debug().AutoMigrate(
// 		models.User{},
// 		models.Category{},
// 	)
// 	if err != nil {
// 		log.Fatalf("failed to migrate database %+v", err)
// 		return serror.NewFromError(err)
// 	}

// 	if db.Migrator().HasTable(&models.User{}) {
// 		if err := db.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
// 			users := []models.User{
// 				{Name: "Admin 1", Email: "admin@gmail.com", Password: "admin123"},
// 			}
// 			if err := db.Create(&users).Error; err != nil {
// 				log.Printf("Error seeding users: %s", err)
// 			} else {
// 				log.Println("Users seeded successfully")
// 			}
// 		}
// 	}

// 	cfg.DB = db

// 	GlobalShutdown.RegisterGracefullyShutdown("database/postgres", func(ctx context.Context) error {
// 		return func() error {
// 			db, err := cfg.DB.DB()
// 			if err != nil {
// 				return err
// 			}
// 			return db.Close()
// 		}()
// 	})

// 	return nil
// }

func (cfg *Config) InitPostgres() serror.SError {
	sqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname =%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	db, err := sqlx.Connect("postgres", sqlConn)
	if err != nil {
		log.Fatalf("Failed connect to database %+v", err)
		return serror.NewFromError(err)
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(utint.StringToInt(utstring.Env("DB_CONNECTION_LIFETIME", "15"), 15)))
	db.SetMaxIdleConns(int(utint.StringToInt(utstring.Env("DB_CONN_MAX_IDLE", "5"), 5)))
	db.SetMaxOpenConns(int(utint.StringToInt(utstring.Env("DB_CONN_MAX_OPEN", "0"), 0)))

	cfg.DB = db

	GlobalShutdown.RegisterGracefullyShutdown("database/postgres", func(ctx context.Context) error {
		return cfg.DB.Close()
	})

	return nil

}
