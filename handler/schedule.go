package handler

import (
	"Susun_Jadwal/models"
	"Susun_Jadwal/service"
	"Susun_Jadwal/util"
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
		util.ErrorEmptyField(c)
		return
	}
	time := c.PostForm("time")
	if time == "" {
		util.ErrorEmptyField(c)
		return
	}
	room := c.PostForm("room")
	if room == "" {
		util.ErrorEmptyField(c)
		return
	}
	classIdStr := c.PostForm("classId")
	if classIdStr == "" {
		util.ErrorEmptyField(c)
		return
	}
	classId, ok := strconv.Atoi(classIdStr)
	if ok != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return
	}

	input := models.ScheduleAddReq{
		Day:     day,
		Time:    time,
		Room:    room,
		ClassID: int32(classId),
	}

	if _, err := sh.scheduleService.CreateNewSubjectSchedules(input); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to create new schedule", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "Succes to create new schedule", input)
}

func (sh *ScheduleHandler) GetSchedules(c *gin.Context) {
	result, err := sh.scheduleService.ListAllMajorSchedules()
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to load data's", err)
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
	util.HttpSuccessResponse(c, 200, "Succes to get data's", fixResult)
}

func (sh *ScheduleHandler) GetScheduleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return
	}

	result, err := sh.scheduleService.GetSchedulesById(int32(id))
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to get data", err)
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

	util.HttpSuccessResponse(c, 200, "Data found", fixResult)
}

func (sh *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}
	newDay := c.PostForm("day")
	newTime := c.PostForm("time")
	newRoom := c.PostForm("room")
	newClassIdStr := c.PostForm("classId")

	if err := sh.scheduleService.UpdateSchedule(c, int32(id), newDay, newTime, newRoom, newClassIdStr); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to update data", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "Succes to update data", gin.H{})
}

func (sh *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}

	if err := sh.scheduleService.DeleteSchedule(c, int32(id)); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to delete a class", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "Data deleted", gin.H{})
}
