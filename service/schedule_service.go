package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/models"
	"Susun_Jadwal/repository"
	"database/sql"
)

type ScheduleService interface {
	CreateNewSubjectSchedules(arg models.ScheduleAddReq) (sql.Result, error)
	DeleteSchedule(id int32) error
	GetSchedulesById(id int32) (sqlc.SubjectSchedule, error)
	ListAllMajorSchedules() ([]sqlc.SubjectSchedule, error)
	UpdateSchedule(arg models.ScheduleUpdateReq) error
}

type scheduleService struct {
	scheduleRepository repository.SchedulesRepository
}

func NewScheduleService(schedulesRepository repository.SchedulesRepository) ScheduleService {
	return &scheduleService{schedulesRepository}
}

func (s scheduleService) CreateNewSubjectSchedules(arg models.ScheduleAddReq) (sql.Result, error) {
	input := sqlc.CreateNewSubjectSchedulesParams{
		Day:     arg.Day,
		Time:    arg.Time,
		Room:    arg.Room,
		ClassID: arg.ClassID,
	}
	return s.scheduleRepository.CreateNewSubjectSchedules(input)
}

func (s scheduleService) DeleteSchedule(id int32) error {
	return s.scheduleRepository.DeleteSchedule(id)
}

func (s scheduleService) GetSchedulesById(id int32) (sqlc.SubjectSchedule, error) {
	return s.scheduleRepository.GetSchedulesById(id)
}

func (s scheduleService) ListAllMajorSchedules() ([]sqlc.SubjectSchedule, error) {
	return s.scheduleRepository.ListAllMajorSchedules()
}

func (s scheduleService) UpdateSchedule(arg models.ScheduleUpdateReq) error {
	input := sqlc.UpdateScheduleParams{
		Day:     arg.Day,
		Time:    arg.Time,
		Room:    arg.Room,
		ClassID: arg.ClassID,
		ID:      arg.ID,
	}
	return s.scheduleRepository.UpdateSchedule(input)
}
