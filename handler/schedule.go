package handler

import (
	"Susun_Jadwal/models"
	"Susun_Jadwal/sdk"
	"Susun_Jadwal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ScheduleHandler struct {
	scheduleService service.ScheduleService
}

func NewScheduleHandler(scheduleService service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{scheduleService}
}

func (sh *ScheduleHandler) CreateSchedule(c *gin.Context) {
	day := c.PostForm("day")
	if day == "" {
		sdk.FailEmptyField(c)
		return
	}
	time := c.PostForm("time")
	if time == "" {
		sdk.FailEmptyField(c)
		return
	}
	room := c.PostForm("room")
	if room == "" {
		sdk.FailEmptyField(c)
		return
	}
	classIdStr := c.PostForm("classId")
	if classIdStr == "" {
		sdk.FailEmptyField(c)
		return
	}
	classId, ok := strconv.Atoi(classIdStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}

	input := models.ScheduleAddReq{
		Day:     day,
		Time:    time,
		Room:    room,
		ClassID: int32(classId),
	}

	if _, err := sh.scheduleService.CreateNewSubjectSchedules(input); err != nil {
		sdk.FailOrError(c, 500, "Failed to create new schedule", err)
		return
	}
	sdk.Success(c, 200, "Succes to create new schedule", input)
}

func (sh *ScheduleHandler) GetSchedules(c *gin.Context) {
	result, err := sh.scheduleService.ListAllMajorSchedules()
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to load data's", err)
		return
	}

	var fixResult []models.ScheduleResult

	for _, datas := range result {
		fixResult = append(fixResult, models.ScheduleResult{
			Id:        int(datas.ID),
			CreatedAt: datas.Createdat.Time,
			UpdatedAt: datas.Updatedat.Time,
			DeletedAt: datas.Deletedat.Time,
			Day:       datas.Day,
			Room:      datas.Room,
			Time:      datas.Time,
			ClassId:   int(datas.ClassID),
		})
	}
	sdk.Success(c, 200, "Succes to get data's", fixResult)
}

func (sh *ScheduleHandler) GetScheduleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}

	result, err := sh.scheduleService.GetSchedulesById(int32(id))
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
		return
	}

	fixResult := models.ScheduleResult{
		Id:        int(result.ID),
		CreatedAt: result.Createdat.Time,
		UpdatedAt: result.Updatedat.Time,
		DeletedAt: result.Deletedat.Time,
		Day:       result.Day,
		Room:      result.Room,
		Time:      result.Time,
		ClassId:   int(result.ClassID),
	}

	sdk.Success(c, 200, "Data found", fixResult)
}

func (sh *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}
	data, err := sh.scheduleService.GetSchedulesById(int32(id))
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
		return
	}
	oldDay := data.Day
	oldTime := data.Time
	oldRoom := data.Room
	oldClassId := data.ClassID

	newDay := c.PostForm("day")
	if newDay == "" {
		newDay = oldDay
	}
	newTime := c.PostForm("time")
	if newTime == "" {
		newTime = oldTime
	}
	newRoom := c.PostForm("room")
	if newRoom == "" {
		newRoom = oldRoom
	}
	newClassIdStr := c.PostForm("classId")
	newClassId := 0

	if newClassIdStr == "" {
		newClassId = int(oldClassId)
	} else {
		newClassId, ok = strconv.Atoi(newClassIdStr)
		if ok != nil {
			sdk.FailOrError(c, 500, "Failed to convert", err)
			return
		}
	}
	input := models.ScheduleUpdateReq{
		Day:     newDay,
		Time:    newTime,
		Room:    newRoom,
		ClassID: int32(newClassId),
		ID:      int32(id),
	}

	if err = sh.scheduleService.UpdateSchedule(input); err != nil {
		sdk.FailOrError(c, 500, "Failed to update data", err)
		return
	}
	sdk.Success(c, 200, "Succes to update data", input)
}

func (sh *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}
	result, err := sh.scheduleService.GetSchedulesById(int32(id))
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
		return
	}
	name := result.Room
	time := result.Time
	classId := result.ClassID

	if err = sh.scheduleService.DeleteSchedule(int32(id)); err != nil {
		sdk.FailOrError(c, 500, "Failed to delete a class", err)
		return
	}
	sdk.Success(c, 200, "Data deleted", gin.H{
		"message": fmt.Sprintf("Class %s with time %s (class id: %d) is successfully deleted", name, time, classId),
	})
}
