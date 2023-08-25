package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/repository"
	"context"
	"database/sql"
)

type ProdiService interface {
	GetAllProdi(ctx context.Context) ([]sqlc.ProgramStudy, error)
	GetProdiById(ctx context.Context, id int32) (sqlc.ProgramStudy, error)
	UpdateProdi(ctx context.Context, arg sqlc.UpdateProdiParams) error
	CreateNewProdi(ctx context.Context, name string) (sql.Result, error)
	DeleteProdi(ctx context.Context, id int32) error
}

type prodiService struct {
	prodiRepo repository.ProdiRepo
}

func NewProdiService(prodiRepo repository.ProdiRepo) ProdiService {
	return &prodiService{prodiRepo}
}

func (ps *prodiService) GetAllProdi(ctx context.Context) ([]sqlc.ProgramStudy, error) {
	return ps.prodiRepo.GetAllProdi(ctx)
}
func (ps *prodiService) GetProdiById(ctx context.Context, id int32) (sqlc.ProgramStudy, error) {
	return ps.prodiRepo.GetProdiById(ctx, id)
}
func (ps *prodiService) UpdateProdi(ctx context.Context, arg sqlc.UpdateProdiParams) error {
	return ps.prodiRepo.UpdateProdi(ctx, arg)
}
func (ps *prodiService) CreateNewProdi(ctx context.Context, name string) (sql.Result, error) {
	return ps.prodiRepo.CreateNewProdi(ctx, name)
}
func (ps *prodiService) DeleteProdi(ctx context.Context, id int32) error {
	return ps.prodiRepo.DeleteProdi(ctx, id)
}
