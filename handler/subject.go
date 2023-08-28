package handler

import (
	"Susun_Jadwal/models"
	"Susun_Jadwal/service"
	"Susun_Jadwal/util"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SubjectHandler struct {
	subjectService service.SubjectService
}

func NewSubjectHandler(subjectService service.SubjectService) *SubjectHandler {
	return &SubjectHandler{subjectService}
}

func (sh *SubjectHandler) CreateSubject(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		util.ErrorEmptyField(c)
		return
	}
	curriculum := c.PostForm("curr")
	if curriculum == "" {
		util.ErrorEmptyField(c)
		return
	}
	sksStr := c.PostForm("sks")
	if sksStr == "" {
		util.ErrorEmptyField(c)
		return
	}
	sks, ok := util.ErrorConvertStr(sksStr, c)
	if ok != nil {
		return
	}

	idProdiStr := c.PostForm("idProdi")
	if idProdiStr == "" {
		util.ErrorEmptyField(c)
		return
	}
	idProdi, ok := util.ErrorConvertStr(idProdiStr, c)
	if ok != nil {
		return
	}
	req := models.SubjectReq{
		Name:       name,
		Curriculum: curriculum,
		Sks:        sks,
		IdProdi:    idProdi,
	}
	if _, err := sh.subjectService.CreateNewSubject(context.Background(), req); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to create subject", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "New subject created", req)
}

func (sh *SubjectHandler) GetAllSubjects(c *gin.Context) {
	result, err := sh.subjectService.GetAllSubjects(context.Background())
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to get data's", err)
		return
	}
	var fixResult []models.SubjectResult

	for _, subjects := range result {
		fixResult = append(fixResult, models.SubjectResult{
			Id:         int(subjects.ID),
			CreatedAt:  subjects.Createdat.Time,
			UpdatedAt:  subjects.Updatedat.Time,
			DeletedAt:  subjects.Deletedat.Time,
			Name:       subjects.Name,
			Curriculum: subjects.Curriculum,
			Sks:        int(subjects.Sks),
			IdProdi:    int(subjects.IDProdi),
		})
	}
	util.HttpSuccessResponse(c, 200, "Succes to get data's", fixResult)
}

func (sh *SubjectHandler) GetSubjectById(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}
	result, err := sh.subjectService.GetSubjectById(context.Background(), int32(id))
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to get data", err)
		return
	}

	fixResult := models.SubjectResult{
		Id:         int(result.ID),
		CreatedAt:  result.Createdat.Time,
		UpdatedAt:  result.Updatedat.Time,
		DeletedAt:  result.Deletedat.Time,
		Name:       result.Name,
		Curriculum: result.Curriculum,
		Sks:        int(result.Sks),
		IdProdi:    int(result.IDProdi),
	}
	util.HttpSuccessResponse(c, 200, "Data found", fixResult)
}

func (sh *SubjectHandler) UpdateSubject(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}
	result, err := sh.subjectService.GetSubjectById(context.Background(), int32(id))
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to get data", err)
		return
	}
	oldName := result.Name
	oldCurriculum := result.Curriculum
	oldSks := result.Sks
	oldIdProdi := result.IDProdi

	newName := c.PostForm("name")
	if newName == "" {
		newName = oldName
	}
	newCurriculum := c.PostForm("curr")
	if newCurriculum == "" {
		newCurriculum = oldCurriculum
	}
	newSksStr := c.PostForm("sks")
	newSks := 0
	ok = err
	if newSksStr == "" {
		newSksStr = strconv.Itoa(int(oldSks))
		newSks, ok = strconv.Atoi(newSksStr)
		if ok != nil {
			util.HttpFailOrErrorResponse(c, 500, "Failed to convert", err)
			return
		}
	} else {
		newSks, ok = util.ErrorConvertStr(newSksStr, c)
		if ok != nil {
			return
		}
	}

	newIdProdiStr := c.PostForm("idProdi")
	newIdProdi := 0
	if newIdProdiStr == "" {
		newIdProdiStr = strconv.Itoa(int(oldIdProdi))
		newIdProdi, ok = strconv.Atoi(newIdProdiStr)
		if ok != nil {
			util.HttpFailOrErrorResponse(c, 500, "Failed to convert", err)
			return
		}
	} else {
		newIdProdi, ok = util.ErrorConvertStr(newIdProdiStr, c)
		if ok != nil {
			return
		}
	}

	req := models.SubjectReq{
		Name:       newName,
		Curriculum: newCurriculum,
		Sks:        newSks,
		IdProdi:    newIdProdi,
		Id:         id,
	}

	if err = sh.subjectService.UpdateSubject(context.Background(), req); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to update", err)
		return
	}

	util.HttpSuccessResponse(c, 200, "Data updated", req)
}

func (sh *SubjectHandler) DeleteSubject(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}
	result, err := sh.subjectService.GetSubjectById(context.Background(), int32(id))
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to get data", err)
		return
	}
	if err = sh.subjectService.DeleteSubject(context.Background(), int32(id)); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to delete", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "Data deleted", gin.H{
		"message": "program study: " + result.Name + " successfully deleted",
	})
}
