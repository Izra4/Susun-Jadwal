package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/models"
	"Susun_Jadwal/repository"
	"Susun_Jadwal/util"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ScheduleService interface {
	CreateNewSubjectSchedules(arg models.ScheduleAddReq) (sql.Result, error)
	DeleteSchedule(c *gin.Context, id int32) error
	GetSchedulesById(id int32) (sqlc.SubjectSchedule, error)
	ListAllMajorSchedules() ([]sqlc.SubjectSchedule, error)
	UpdateSchedule(c *gin.Context, id int32, newDay string, newTime string, newRoom string, newClassIdStr string) error
}

type scheduleService struct {
	scheduleRepository repository.SchedulesRepository
}

func NewScheduleService(schedulesRepository repository.SchedulesRepository) ScheduleService {
	return &scheduleService{schedulesRepository}
}

func (s *scheduleService) CreateNewSubjectSchedules(arg models.ScheduleAddReq) (sql.Result, error) {
	input := sqlc.CreateNewSubjectSchedulesParams{
		Day:     arg.Day,
		Time:    arg.Time,
		Room:    arg.Room,
		ClassID: arg.ClassID,
	}
	return s.scheduleRepository.CreateNewSubjectSchedules(input)
}

func (s *scheduleService) DeleteSchedule(c *gin.Context, id int32) error {
	_, err := s.scheduleRepository.GetSchedulesById(id)
	if err != nil {
		util.HttpFailOrErrorResponse(c, 404, "Data not found", err)
		return err
	}
	return s.scheduleRepository.DeleteSchedule(id)
}

func (s *scheduleService) GetSchedulesById(id int32) (sqlc.SubjectSchedule, error) {
	return s.scheduleRepository.GetSchedulesById(id)
}

func (s *scheduleService) ListAllMajorSchedules() ([]sqlc.SubjectSchedule, error) {
	return s.scheduleRepository.ListAllMajorSchedules()
}

func (s *scheduleService) UpdateSchedule(c *gin.Context, id int32, newDay string, newTime string, newRoom string, newClassIdStr string) error {
	data, err := s.scheduleRepository.GetSchedulesById(id)
	if err != nil {
		util.HttpFailOrErrorResponse(c, 404, "Failed to get data", err)
		return err
	}
	oldDay := data.Day
	oldTime := data.Time
	oldRoom := data.Room
	oldClassId := data.ClassID

	if newDay == "" {
		newDay = oldDay
	}
	if newTime == "" {
		newTime = oldTime
	}
	if newRoom == "" {
		newRoom = oldRoom
	}
	newClassId := 0
	ok := err
	if newClassIdStr == "" {
		newClassId = int(oldClassId)
	} else {
		newClassId, ok = strconv.Atoi(newClassIdStr)
		if ok != nil {
			util.HttpFailOrErrorResponse(c, 500, "Failed to convert", err)
			return err
		}
	}
	input := sqlc.UpdateScheduleParams{
		Day:     newDay,
		Time:    newTime,
		Room:    newRoom,
		ClassID: int32(newClassId),
		ID:      id,
	}
	return s.scheduleRepository.UpdateSchedule(input)
}
