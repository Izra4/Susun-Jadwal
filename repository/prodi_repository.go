package repository

import (
	"Susun_Jadwal/db/sqlc"
	"context"
	"database/sql"
)

type ProdiRepo interface {
	GetAllProdi(ctx context.Context) ([]sqlc.ProgramStudy, error)
	GetProdiById(ctx context.Context, id int32) (sqlc.ProgramStudy, error)
	UpdateProdi(ctx context.Context, arg sqlc.UpdateProdiParams) error
	CreateNewProdi(ctx context.Context, name string) (sql.Result, error)
	DeleteProdi(ctx context.Context, id int32) error
}

type prodiRepo struct {
	db *sqlc.Queries
}

func NewProdiRepository(db *sqlc.Queries) ProdiRepo {
	return &prodiRepo{db}
}

func (pr *prodiRepo) GetAllProdi(ctx context.Context) ([]sqlc.ProgramStudy, error) {
	return pr.db.GetAllProdi(ctx)
}
func (pr *prodiRepo) GetProdiById(ctx context.Context, id int32) (sqlc.ProgramStudy, error) {
	return pr.db.GetProdiById(ctx, id)
}
func (pr *prodiRepo) UpdateProdi(ctx context.Context, arg sqlc.UpdateProdiParams) error {
	return pr.db.UpdateProdi(ctx, arg)
}
func (pr *prodiRepo) CreateNewProdi(ctx context.Context, name string) (sql.Result, error) {
	return pr.db.CreateNewProdi(ctx, name)
}
func (pr *prodiRepo) DeleteProdi(ctx context.Context, id int32) error {
	return pr.db.DeleteProdi(ctx, id)
}
