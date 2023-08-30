package handler

import (
	"Susun_Jadwal/models"
	"Susun_Jadwal/service"
	"Susun_Jadwal/util"
	"context"
	"github.com/gin-gonic/gin"
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
	newName := c.PostForm("name")
	newCurriculum := c.PostForm("curr")
	newSksStr := c.PostForm("sks")
	newIdProdiStr := c.PostForm("idProdi")

	if err := sh.subjectService.UpdateSubject(context.Background(), c, id, newName, newCurriculum, newSksStr, newIdProdiStr); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to update", err)
		return
	}

	util.HttpSuccessResponse(c, 200, "Data updated", gin.H{})
}

func (sh *SubjectHandler) DeleteSubject(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}
	if err := sh.subjectService.DeleteSubject(context.Background(), c, int32(id)); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to delete", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "Data deleted", gin.H{
		"message": "program study successfully deleted",
	})
}
