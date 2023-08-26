package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/models"
	"Susun_Jadwal/repository"
	"context"
	"database/sql"
)

type ClassService interface {
	AddNewClass(ctx context.Context, arg models.ClassAddReq) (sql.Result, error)
	DeleteClass(ctx context.Context, id int32) error
	GetClassById(ctx context.Context, id int32) (sqlc.Class, error)
	ListClass(ctx context.Context) ([]sqlc.Class, error)
	UpdateClass(ctx context.Context, arg models.ClassUpdateReq) error
}

type classService struct {
	repo repository.ClassRepository
}

func NewClassService(repo repository.ClassRepository) ClassService {
	return &classService{repo}
}

func (c *classService) AddNewClass(ctx context.Context, arg models.ClassAddReq) (sql.Result, error) {
	req := sqlc.AddNewClassParams{
		Name:      arg.Name,
		Member:    int32(arg.Member),
		SubjectID: int32(arg.SubjectId),
	}

	return c.repo.AddNewClass(ctx, req)
}

func (c *classService) DeleteClass(ctx context.Context, id int32) error {
	return c.repo.DeleteClass(ctx, id)
}

func (c *classService) GetClassById(ctx context.Context, id int32) (sqlc.Class, error) {
	return c.repo.GetClassById(ctx, id)
}

func (c *classService) ListClass(ctx context.Context) ([]sqlc.Class, error) {
	return c.repo.ListClass(ctx)
}

func (c *classService) UpdateClass(ctx context.Context, arg models.ClassUpdateReq) error {
	req := sqlc.UpdateClassParams{
		Name:      arg.Name,
		Member:    arg.Member,
		SubjectID: arg.SubjectID,
		ID:        arg.ID,
	}
	return c.repo.UpdateClass(ctx, req)
}
