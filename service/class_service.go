package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/models"
	"Susun_Jadwal/repository"
	"Susun_Jadwal/util"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type ClassService interface {
	AddNewClass(ctx context.Context, arg models.ClassAddReq) (sql.Result, error)
	DeleteClass(ctx context.Context, id int32) error
	GetClassById(ctx context.Context, id int32) (sqlc.Class, error)
	ListClass(ctx context.Context) ([]sqlc.Class, error)
	UpdateClass(ctx context.Context, cgx *gin.Context, id int32, newName string, newMemberStr string,
		newSubjectIdStr string) error
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
	_, err := c.repo.GetClassById(context.Background(), id)
	if err != nil {
		return nil
	}
	return c.repo.DeleteClass(ctx, id)
}

func (c *classService) GetClassById(ctx context.Context, id int32) (sqlc.Class, error) {
	return c.repo.GetClassById(ctx, id)
}

func (c *classService) ListClass(ctx context.Context) ([]sqlc.Class, error) {
	return c.repo.ListClass(ctx)
}

func (c *classService) UpdateClass(ctx context.Context, cgx *gin.Context, id int32, newName string, newMemberStr string,
	newSubjectIdStr string) error {
	result, err := c.repo.GetClassById(context.Background(), id)
	if err != nil {
		util.HttpFailOrErrorResponse(cgx, 500, "Data not found", err)
		return nil
	}
	ok := err
	oldName := result.Name
	oldMember := result.Member
	oldSubjectId := result.SubjectID
	if newName == "" {
		newName = oldName
	}
	newMember := 0
	if newMemberStr == "" {
		newMember = int(oldMember)
	} else {
		newMember, ok = util.ErrorConvertStr(newMemberStr, cgx)
		if ok != nil {
			util.HttpFailOrErrorResponse(cgx, 400, "Failed to convert", ok)
			return nil
		}
	}
	newSubjectId := 0
	if newSubjectIdStr == "" {
		newSubjectId = int(oldSubjectId)
	} else {
		newSubjectId, ok = util.ErrorConvertStr(newSubjectIdStr, cgx)
		if ok != nil {
			return nil
		}
	}
	req := sqlc.UpdateClassParams{
		Name:      newName,
		Member:    int32(newMember),
		SubjectID: int32(newSubjectId),
		ID:        id,
	}
	return c.repo.UpdateClass(ctx, req)
}
