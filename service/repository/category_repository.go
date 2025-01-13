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

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) api.CategoryRepository {
	return &categoryRepo{db}
}

func (r *categoryRepo) CreateCategory(ctx context.Context, req models.CreateCategoryRequest) (categoryId int64, errx serror.SError) {
	err := r.db.QueryRowxContext(ctx, queries.InsertCategory,
		req.CategoryName,
		time.Now(),
		time.Now(),
	).Scan(&categoryId)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[repository][CreateCategory] while QueryRowxContext")
		return
	}

	return
}

func (r *categoryRepo) GetCategories(ctx context.Context, req models.GetCategoryRequest) (res []models.Category, errx serror.SError) {
	var (
		rows *sqlx.Rows
		err  error
	)
	rows, err = r.db.QueryxContext(ctx, queries.GetCategories)
	if err != nil && err != sql.ErrNoRows {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[repository][GetCategories] while QueryxContext")
		return
	}

	defer rows.Close()

	for rows.Next() {
		var tmp models.Category
		err = rows.StructScan(&tmp)
		if err != nil {
			errx = serror.NewFromError(err)
			errx.AddCommentf("[repository][GetCategories] while StructScan")
			return
		}

		res = append(res, tmp)
	}

	return
}

func (r *categoryRepo) GetCategoryByID(ctx context.Context, categoryId int64) (res models.Category, errx serror.SError) {
	err := r.db.QueryRowxContext(ctx, queries.GetCategoryByID, categoryId).StructScan(&res)
	if err != nil && err != sql.ErrNoRows {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[repository][GetCategoryByID][CategoryID: %d] while QueryRowxContext queries.GetCategoryByID", categoryId)
		return
	}

	return
}

func (r *categoryRepo) UpdateCategoryByID(ctx context.Context, tx *sqlx.DB, req models.UpdateCategoryRequest) (res models.Category, errx serror.SError) {
	var err error
	if tx == nil {
		_, err = r.db.ExecContext(ctx, queries.UpdateCategoryByID,
			req.CategoryID,
			req.CategoryName,
			time.Now(),
		)
	} else {
		_, err = tx.ExecContext(ctx, queries.UpdateCategoryByID,
			req.CategoryID,
			req.CategoryName,
			time.Now(),
		)
	}
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[repository][UpdateCategoryByID][CategoryID: %d] while ExecContext queries.UpdateCategoryByID", req.CategoryID)
		return
	}

	err = r.db.QueryRowxContext(ctx, queries.GetCategoryByID, req.CategoryID).StructScan(&res)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[repository][UpdateCategoryByID][CategoryID: %d] while QueryRowxContext queries.GetCategoryByID", req.CategoryID)
		return
	}

	return
}

func (r *categoryRepo) DeleteCategory(ctx context.Context, categoryId int64) (errx serror.SError) {
	var err error
	_, err = r.db.ExecContext(ctx, queries.DeleteCategoryByID,
		categoryId,
	)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddCommentf("[repository][DeleteCategory][CategoryID: %d] while ExecContext queries.DeleteCategory", categoryId)
		return
	}

	return
}
