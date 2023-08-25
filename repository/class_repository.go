package repository

import (
	"Susun_Jadwal/db/sqlc"
	"context"
	"database/sql"
)

type ClassRepository interface {
	CreateClass(ctx context.Context, class sqlc.AddNewClassParams) (sql.Result, error)
	DeleteClass(ctx context.Context, id int32) error
	GetClassNameById(ctx context.Context, id int32) (string, error)
	GetClassById(ctx context.Context, id int32) (sqlc.Class, error)
	GetClassByName(ctx context.Context, name string) ([]sqlc.Class, error)
	GetListClass(ctx context.Context) ([]sqlc.Class, error)
}

type classRepository struct {
	db *sqlc.Queries
}

func NewClassRepository(db *sqlc.Queries) ClassRepository {
	return &classRepository{db}
}

func (r *classRepository) CreateClass(ctx context.Context, class sqlc.AddNewClassParams) (sql.Result, error) {
	return r.db.AddNewClass(ctx, class)
}

func (r *classRepository) DeleteClass(ctx context.Context, id int32) error {
	return r.db.DeleteClass(ctx, id)
}

func (r *classRepository) GetClassNameById(ctx context.Context, id int32) (string, error) {
	return r.db.GetClassNameById(ctx, id)
}

func (r *classRepository) GetClassById(ctx context.Context, id int32) (sqlc.Class, error) {
	return r.db.GetClassById(ctx, id)
}
func (r *classRepository) GetClassByName(ctx context.Context, name string) ([]sqlc.Class, error) {
	return r.db.GetClassByName(ctx, name)
}
func (r *classRepository) GetListClass(ctx context.Context) ([]sqlc.Class, error) {
	return r.db.ListClass(ctx)
}
