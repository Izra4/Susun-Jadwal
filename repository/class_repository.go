package repository

import (
	"Susun_Jadwal/db/sqlc"
	"context"
	"database/sql"
)

type ClassRepository interface {
	AddNewClass(ctx context.Context, arg sqlc.AddNewClassParams) (sql.Result, error)
	DeleteClass(ctx context.Context, id int32) error
	GetClassById(ctx context.Context, id int32) (sqlc.Class, error)
	ListClass(ctx context.Context) ([]sqlc.Class, error)
	UpdateClass(ctx context.Context, arg sqlc.UpdateClassParams) error
}

type classRepository struct {
	db *sqlc.Queries
}

func NewClassRepository(db *sqlc.Queries) ClassRepository {
	return &classRepository{db}
}

func (c *classRepository) AddNewClass(ctx context.Context, arg sqlc.AddNewClassParams) (sql.Result, error) {
	return c.db.AddNewClass(ctx, arg)
}

func (c *classRepository) DeleteClass(ctx context.Context, id int32) error {
	return c.db.DeleteClass(ctx, id)
}

func (c *classRepository) GetClassById(ctx context.Context, id int32) (sqlc.Class, error) {
	return c.db.GetClassById(ctx, id)
}

func (c *classRepository) ListClass(ctx context.Context) ([]sqlc.Class, error) {
	return c.db.ListClass(ctx)
}

func (c *classRepository) UpdateClass(ctx context.Context, arg sqlc.UpdateClassParams) error {
	return c.db.UpdateClass(ctx, arg)
}
