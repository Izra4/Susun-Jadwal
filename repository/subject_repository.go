package repository

import (
	"Susun_Jadwal/db/sqlc"
	"context"
	"database/sql"
)

type SubjectRepository interface {
	CreateNewSubject(ctx context.Context, arg sqlc.CreateNewSubjectParams) (sql.Result, error)
	DeleteSubject(ctx context.Context, id int32) error
	GetAllSubjects(ctx context.Context) ([]sqlc.Subject, error)
	GetSubjectById(ctx context.Context, id int32) (sqlc.Subject, error)
	UpdateSubject(ctx context.Context, arg sqlc.UpdateSubjectParams) error
}

type subjectRepository struct {
	db *sqlc.Queries
}

func NewSubjectRepository(db *sqlc.Queries) SubjectRepository {
	return &subjectRepository{db}
}

func (sr *subjectRepository) CreateNewSubject(ctx context.Context, arg sqlc.CreateNewSubjectParams) (sql.Result, error) {
	return sr.db.CreateNewSubject(ctx, arg)
}
func (sr *subjectRepository) DeleteSubject(ctx context.Context, id int32) error {
	return sr.db.DeleteSubject(ctx, id)
}
func (sr *subjectRepository) GetAllSubjects(ctx context.Context) ([]sqlc.Subject, error) {
	return sr.db.GetAllSubjects(ctx)
}
func (sr *subjectRepository) GetSubjectById(ctx context.Context, id int32) (sqlc.Subject, error) {
	return sr.db.GetSubjectById(ctx, id)
}
func (sr *subjectRepository) UpdateSubject(ctx context.Context, arg sqlc.UpdateSubjectParams) error {
	return sr.db.UpdateSubject(ctx, arg)
}
