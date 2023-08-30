package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/models"
	"Susun_Jadwal/repository"
	"Susun_Jadwal/util"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SubjectService interface {
	CreateNewSubject(ctx context.Context, arg models.SubjectReq) (sql.Result, error)
	DeleteSubject(ctx context.Context, c *gin.Context, id int32) error
	GetAllSubjects(ctx context.Context) ([]sqlc.Subject, error)
	GetSubjectById(ctx context.Context, id int32) (sqlc.Subject, error)
	UpdateSubject(ctx context.Context, c *gin.Context, id int, newName string, newCurriculum string, newSksStr string, newIdProdiStr string) error
}

func NewSubjectService(subjectRepository repository.SubjectRepository) SubjectService {
	return &subjectService{subjectRepository}
}

type subjectService struct {
	subjectRepository repository.SubjectRepository
}

func (s *subjectService) CreateNewSubject(ctx context.Context, arg models.SubjectReq) (sql.Result, error) {
	data := sqlc.CreateNewSubjectParams{
		Name:       arg.Name,
		Curriculum: arg.Curriculum,
		Sks:        int32(arg.Sks),
		IDProdi:    int32(arg.IdProdi),
	}
	return s.subjectRepository.CreateNewSubject(ctx, data)
}

func (s *subjectService) DeleteSubject(ctx context.Context, c *gin.Context, id int32) error {
	_, err := s.subjectRepository.GetSubjectById(context.Background(), id)
	if err != nil {
		util.HttpFailOrErrorResponse(c, 400, "Failed to get data", err)
		return err
	}
	return s.subjectRepository.DeleteSubject(ctx, id)
}

func (s *subjectService) GetAllSubjects(ctx context.Context) ([]sqlc.Subject, error) {
	return s.subjectRepository.GetAllSubjects(ctx)
}

func (s *subjectService) GetSubjectById(ctx context.Context, id int32) (sqlc.Subject, error) {
	return s.subjectRepository.GetSubjectById(ctx, id)
}

func (s *subjectService) UpdateSubject(ctx context.Context, c *gin.Context, id int, newName string, newCurriculum string, newSksStr string, newIdProdiStr string) error {
	result, err := s.subjectRepository.GetSubjectById(context.Background(), int32(id))
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to get data", err)
		return err
	}
	ok := err
	oldName := result.Name
	oldCurriculum := result.Curriculum
	oldSks := result.Sks
	oldIdProdi := result.IDProdi

	if newName == "" {
		newName = oldName
	}
	if newCurriculum == "" {
		newCurriculum = oldCurriculum
	}
	newSks := 0
	ok = err
	if newSksStr == "" {
		newSksStr = strconv.Itoa(int(oldSks))
		newSks, ok = strconv.Atoi(newSksStr)
		if ok != nil {
			util.HttpFailOrErrorResponse(c, 400, "Failed to convert", err)
			return err
		}
	} else {
		newSks, ok = util.ErrorConvertStr(newSksStr, c)
		if ok != nil {
			return err
		}
	}

	newIdProdi := 0
	if newIdProdiStr == "" {
		newIdProdiStr = strconv.Itoa(int(oldIdProdi))
		newIdProdi, ok = strconv.Atoi(newIdProdiStr)
		if ok != nil {
			util.HttpFailOrErrorResponse(c, 400, "Failed to convert", err)
			return err
		}
	} else {
		newIdProdi, ok = util.ErrorConvertStr(newIdProdiStr, c)
		if ok != nil {
			return err
		}
	}
	data := sqlc.UpdateSubjectParams{
		Name:       newName,
		Curriculum: newCurriculum,
		Sks:        int32(newSks),
		IDProdi:    int32(newIdProdi),
		ID:         int32(id),
	}
	return s.subjectRepository.UpdateSubject(ctx, data)
}
