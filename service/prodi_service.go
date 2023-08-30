package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/repository"
	"Susun_Jadwal/util"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type ProdiService interface {
	GetAllProdi(ctx context.Context) ([]sqlc.ProgramStudy, error)
	GetProdiById(ctx context.Context, id int32) (sqlc.ProgramStudy, error)
	UpdateProdi(ctx context.Context, c *gin.Context, id int32, nameUpdate string) error
	CreateNewProdi(ctx context.Context, name string) (sql.Result, error)
	DeleteProdi(ctx context.Context, id int32, cgx *gin.Context) error
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
func (ps *prodiService) UpdateProdi(ctx context.Context, c *gin.Context, id int32, nameUpdate string) error {
	result, err := ps.prodiRepo.GetProdiById(c, id)
	if err != nil {
		return err
	}
	if nameUpdate == "" {
		nameUpdate = result.Name
	}

	var input sqlc.UpdateProdiParams
	input.ID = id
	input.Name = nameUpdate
	if err = ps.prodiRepo.UpdateProdi(context.Background(), input); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to update", err)
		return err
	}
	return ps.prodiRepo.UpdateProdi(ctx, input)
}
func (ps *prodiService) CreateNewProdi(ctx context.Context, name string) (sql.Result, error) {
	return ps.prodiRepo.CreateNewProdi(ctx, name)
}
func (ps *prodiService) DeleteProdi(ctx context.Context, id int32, cgx *gin.Context) error {
	_, err := ps.prodiRepo.GetProdiById(context.Background(), id)
	if err != nil {
		util.HttpFailOrErrorResponse(cgx, 500, "Class not found", err)
		return err
	}
	return ps.prodiRepo.DeleteProdi(ctx, id)
}
