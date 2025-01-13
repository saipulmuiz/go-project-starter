package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/saipulmuiz/go-project-starter/models"
	"github.com/saipulmuiz/go-project-starter/pkg/serror"
	api "github.com/saipulmuiz/go-project-starter/service"
	"github.com/saipulmuiz/go-project-starter/service/repository/queries"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) api.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Register(ctx context.Context, req models.RegisterUserRequest) (userId int64, errx serror.SError) {
	err := r.db.QueryRowxContext(ctx, queries.RegisterUser,
		req.Name,
		req.Email,
		req.Password,
		time.Now(),
		time.Now(),
	).Scan(&userId)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[repository][Register] while ExecContext queries.RegisterUser")
		return
	}

	return
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (res models.User, errx serror.SError) {
	err := r.db.QueryRowxContext(ctx, queries.GetUserByEmail, email).StructScan(&res)
	if err != nil && err != sql.ErrNoRows {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[repository][GetUserByEmail][Email: %d] while QueryRowxContext", email)
		return
	}

	return
}

func (r *userRepo) GetUserByID(ctx context.Context, userID string) (res models.User, errx serror.SError) {
	err := r.db.QueryRowxContext(ctx, queries.GetUserByID, userID).StructScan(&res)
	if err != nil && err != sql.ErrNoRows {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[repository][GetUserByEmail][UserID: %d] while QueryRowxContext", userID)
		return
	}

	return
}
