package handler

import (
	"Susun_Jadwal/models"
	"Susun_Jadwal/service"
	"Susun_Jadwal/util"
	"context"
	"github.com/gin-gonic/gin"
)

type ClassHandler struct {
	classService service.ClassService
}

func NewClassHandler(classService service.ClassService) *ClassHandler {
	return &ClassHandler{classService}
}

func (ch *ClassHandler) CreateClass(c *gin.Context) {
	name := c.PostForm("name")
	memberStr := c.PostForm("member")
	if memberStr == "" {
		util.ErrorEmptyField(c)
		return
	}
	member, ok := util.ErrorConvertStr(memberStr, c)
	if ok != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return
	}

	subjectIdStr := c.PostForm("subjectId")
	if subjectIdStr == "" {
		util.ErrorEmptyField(c)
		return
	}
	subjectId, ok := util.ErrorConvertStr(subjectIdStr, c)

	req := models.ClassAddReq{
		Name:      name,
		Member:    member,
		SubjectId: subjectId,
	}

	if _, err := ch.classService.AddNewClass(context.Background(), req); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to create a new class", err)
		return
	}

	util.HttpSuccessResponse(c, 200, "Succes to create new class", req)
}

func (ch *ClassHandler) DeleteClass(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}
	//result, err := ch.classService.GetClassById(context.Background(), int32(id))
	//if err != nil {
	//	util.HttpFailOrErrorResponse(c, 500, "Data not found", err)
	//	return
	//}

	if err := ch.classService.DeleteClass(context.Background(), int32(id)); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to delete class", err)
		return
	}

	util.HttpSuccessResponse(c, 200, "Succes to delete data", gin.H{
		"message": "class succesfully deleted",
	})
}

func (ch *ClassHandler) GetAllClasses(c *gin.Context) {
	result, err := ch.classService.ListClass(context.Background())
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to load data's", err)
		return
	}

	var fixResult []models.ClassResult

	for _, datas := range result {
		fixResult = append(fixResult, models.ClassResult{
			Id:        int(datas.ID),
			CreatedAt: datas.Createdat.Time,
			UpdatedAt: datas.Updatedat.Time,
			DeletedAt: datas.Deletedat.Time,
			Name:      datas.Name,
			Member:    int(datas.Member),
			SubjectId: int(datas.SubjectID),
		})
	}
	util.HttpSuccessResponse(c, 200, "Succes to get data", fixResult)
}

func (ch *ClassHandler) GetClassById(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return
	}
	result, err := ch.classService.GetClassById(context.Background(), int32(id))
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to get data", err)
		return
	}

	fixResult := models.ClassResult{
		Id:        int(result.ID),
		CreatedAt: result.Createdat.Time,
		UpdatedAt: result.Updatedat.Time,
		DeletedAt: result.Deletedat.Time,
		Name:      result.Name,
		Member:    int(result.Member),
		SubjectId: int(result.SubjectID),
	}

	util.HttpSuccessResponse(c, 200, "Data found", fixResult)
}

func (ch *ClassHandler) UpdateClass(c *gin.Context) {
	//	name,member,subjectId,id
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return
	}
	newName := c.PostForm("name")
	newMemberStr := c.PostForm("member")
	newSubjectIdStr := c.PostForm("subjectId")

	//result, err := ch.classService.GetClassById(context.Background(), int32(id))
	//if err != nil {
	//	util.HttpFailOrErrorResponse(c, 500, "Data not found", err)
	//	return
	//}
	//oldName := result.Name
	//oldMember := result.Member
	//oldSubjectId := result.SubjectID
	//if newName == "" {
	//	newName = oldName
	//}
	//newMember := 0
	//if newMemberStr == "" {
	//	newMemberStr = strconv.Itoa(int(oldMember))
	//	newMember, ok = strconv.Atoi(newMemberStr)
	//	if ok != nil {
	//		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", err)
	//		return
	//	}
	//} else {
	//	newMember, ok = strconv.Atoi(newMemberStr)
	//	if ok != nil {
	//		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
	//		return
	//	}
	//}
	//newSubjectId := 0
	//if newSubjectIdStr == "" {
	//	newSubjectIdStr = strconv.Itoa(int(oldSubjectId))
	//	newSubjectId, ok = strconv.Atoi(newSubjectIdStr)
	//	if ok != nil {
	//		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
	//		return
	//	}
	//} else {
	//	newSubjectId, ok = strconv.Atoi(newSubjectIdStr)
	//	if ok != nil {
	//		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
	//		return
	//	}
	//}
	//
	//input := models.ClassUpdateReq{
	//	Name:      newName,
	//	Member:    int32(newMember),
	//	SubjectID: int32(newSubjectId),
	//	ID:        int32(id),
	//}

	if err := ch.classService.UpdateClass(context.Background(), c, int32(id), newName, newMemberStr, newSubjectIdStr); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to update", err)
		return
	}

	util.HttpSuccessResponse(c, 200, "Data updated", gin.H{})
}
