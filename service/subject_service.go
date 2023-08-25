package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/models"
	"Susun_Jadwal/repository"
	"context"
	"database/sql"
)

type SubjectService interface {
	CreateNewSubject(ctx context.Context, arg models.SubjectReq) (sql.Result, error)
	DeleteSubject(ctx context.Context, id int32) error
	GetAllSubjects(ctx context.Context) ([]sqlc.Subject, error)
	GetSubjectById(ctx context.Context, id int32) (sqlc.Subject, error)
	UpdateSubject(ctx context.Context, arg models.SubjectReq) error
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

func (s *subjectService) DeleteSubject(ctx context.Context, id int32) error {
	return s.subjectRepository.DeleteSubject(ctx, id)
}

func (s *subjectService) GetAllSubjects(ctx context.Context) ([]sqlc.Subject, error) {
	return s.subjectRepository.GetAllSubjects(ctx)
}

func (s *subjectService) GetSubjectById(ctx context.Context, id int32) (sqlc.Subject, error) {
	return s.subjectRepository.GetSubjectById(ctx, id)
}

func (s *subjectService) UpdateSubject(ctx context.Context, arg models.SubjectReq) error {
	data := sqlc.UpdateSubjectParams{
		Name:       arg.Name,
		Curriculum: arg.Curriculum,
		Sks:        int32(arg.Sks),
		IDProdi:    int32(arg.IdProdi),
		ID:         int32(arg.Id),
	}
	return s.subjectRepository.UpdateSubject(ctx, data)
}
