package handler

import (
	"Susun_Jadwal/models"
	"Susun_Jadwal/service"
	"Susun_Jadwal/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type KrsHandler struct {
	krsService service.KrsService
}

func NewKrsHandler(krsService service.KrsService) *KrsHandler {
	return &KrsHandler{krsService}
}

func (kh *KrsHandler) AddKrs(c *gin.Context) {
	totalsStr := c.PostForm("sks")
	if totalsStr == "" {
		util.ErrorEmptyField(c)
		return
	}
	totals, ok := strconv.Atoi(totalsStr)
	if ok != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return
	}

	userIdStr := c.PostForm("userId")
	if userIdStr == "" {
		util.ErrorEmptyField(c)
		return
	}
	userId, ok := strconv.Atoi(userIdStr)
	if ok != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return
	}
	input := models.KrsAddReq{
		Totals: int32(totals),
		Userid: int32(userId),
	}
	if _, err := kh.krsService.AddKrs(input); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to add a krs", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "Success to create a new KRS", input)
}

func (kh *KrsHandler) GetAllKrs(c *gin.Context) {
	result, err := kh.krsService.GetAllKrs()
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to get data's", err)
		return
	}

	var fixResult []models.KrsResult
	for _, datas := range result {
		fixResult = append(fixResult, models.KrsResult{
			ID:        datas.ID,
			Createdat: datas.Createdat.Time,
			Updatedat: datas.Updatedat.Time,
			Deletedat: datas.Deletedat.Time,
			Totals:    datas.Totals,
			Userid:    datas.Userid,
		})
	}
	util.HttpSuccessResponse(c, 200, "Success to get data", fixResult)
}

func (kh *KrsHandler) GetKrsById(c *gin.Context) {
	idStr := c.PostForm("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}

	result, err := kh.krsService.GetKrsByID(int32(id))
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to get data", err)
		return
	}

	fixResult := models.KrsResult{
		ID:        result.ID,
		Createdat: result.Createdat.Time,
		Updatedat: result.Updatedat.Time,
		Deletedat: result.Deletedat.Time,
		Totals:    result.Totals,
		Userid:    result.Userid,
	}
	util.HttpSuccessResponse(c, 200, "Success to get data", fixResult)
}

func (kh *KrsHandler) GetKrsByUserId(c *gin.Context) {
	idStr := c.PostForm("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}
	result, err := kh.krsService.GetKrsByIDUser(int32(id))
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to get data's", err)
		return
	}
	var fixResult []models.KrsResult

	for _, datas := range result {
		fixResult = append(fixResult, models.KrsResult{
			ID:        datas.ID,
			Createdat: datas.Createdat.Time,
			Updatedat: datas.Updatedat.Time,
			Deletedat: datas.Deletedat.Time,
			Totals:    datas.Totals,
			Userid:    datas.Userid,
		})
	}

	util.HttpSuccessResponse(c, 200, "Success to get data's", fixResult)
}

func (kh *KrsHandler) UpdateKrs(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return
	}
	newTotalsStr := c.PostForm("sks")
	newUserIdStr := c.PostForm("userID")

	if err := kh.krsService.UpdateKrs(c, newTotalsStr, newUserIdStr, int32(id)); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to update", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "Success to update", gin.H{})
}

func (kh *KrsHandler) DeleteKrs(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}
	if err := kh.krsService.DeleteKrs(int32(id)); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to delete", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "Delete Success", gin.H{
		"message": "KRS Deleted",
	})
}
