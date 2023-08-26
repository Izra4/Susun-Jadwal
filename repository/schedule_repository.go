package repository

import (
	"Susun_Jadwal/db/sqlc"
	"context"
	"database/sql"
)

type SchedulesRepository interface {
	CreateNewSubjectSchedules(arg sqlc.CreateNewSubjectSchedulesParams) (sql.Result, error)
	DeleteSchedule(id int32) error
	GetSchedulesById(id int32) (sqlc.SubjectSchedule, error)
	ListAllMajorSchedules() ([]sqlc.SubjectSchedule, error)
	UpdateSchedule(arg sqlc.UpdateScheduleParams) error
}

type scheduleRepository struct {
	db *sqlc.Queries
}

func NewScheduleRepository(db *sqlc.Queries) SchedulesRepository {
	return &scheduleRepository{db}
}

func (s *scheduleRepository) CreateNewSubjectSchedules(arg sqlc.CreateNewSubjectSchedulesParams) (sql.Result, error) {
	return s.db.CreateNewSubjectSchedules(context.Background(), arg)
}

func (s *scheduleRepository) DeleteSchedule(id int32) error {
	return s.db.DeleteSchedule(context.Background(), id)
}

func (s *scheduleRepository) GetSchedulesById(id int32) (sqlc.SubjectSchedule, error) {
	return s.db.GetSchedulesById(context.Background(), id)
}

func (s *scheduleRepository) ListAllMajorSchedules() ([]sqlc.SubjectSchedule, error) {
	return s.db.ListAllMajorSchedules(context.Background())
}

func (s *scheduleRepository) UpdateSchedule(arg sqlc.UpdateScheduleParams) error {
	return s.db.UpdateSchedule(context.Background(), arg)
}
