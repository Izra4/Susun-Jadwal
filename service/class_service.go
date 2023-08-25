package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/repository"
	"context"
	"database/sql"
)

type ClassService interface {
	CreateClass(ctx context.Context, class sqlc.AddNewClassParams) (sql.Result, error)
	DeleteClass(ctx context.Context, id int32) error
	GetClassNameById(ctx context.Context, id int32) (string, error)
	GetClassById(ctx context.Context, id int32) (sqlc.Class, error)
	GetClassByName(ctx context.Context, name string) ([]sqlc.Class, error)
	GetListClass(ctx context.Context) ([]sqlc.Class, error)
}

type classService struct {
	repo repository.ClassRepository
}

func NewClassService(repo repository.ClassRepository) ClassService {
	return &classService{repo}
}

func (s *classService) CreateClass(ctx context.Context, class sqlc.AddNewClassParams) (sql.Result, error) {
	return s.repo.CreateClass(ctx, class)
}

func (s *classService) DeleteClass(ctx context.Context, id int32) error {
	return s.repo.DeleteClass(ctx, id)
}

func (s *classService) GetClassNameById(ctx context.Context, id int32) (string, error) {
	return s.repo.GetClassNameById(ctx, id)
}
func (s *classService) GetClassById(ctx context.Context, id int32) (sqlc.Class, error) {
	return s.repo.GetClassById(ctx, id)
}
func (s *classService) GetClassByName(ctx context.Context, name string) ([]sqlc.Class, error) {
	return s.repo.GetClassByName(ctx, name)
}
func (s *classService) GetListClass(ctx context.Context) ([]sqlc.Class, error) {
	return s.repo.GetListClass(ctx)
}
