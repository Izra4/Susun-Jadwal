package handler

import (
	"Susun_Jadwal/models"
	"Susun_Jadwal/sdk"
	"Susun_Jadwal/service"
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
		sdk.FailEmptyField(c)
		return
	}
	totals, ok := strconv.Atoi(totalsStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}

	userIdStr := c.PostForm("userId")
	if userIdStr == "" {
		sdk.FailEmptyField(c)
		return
	}
	userId, ok := strconv.Atoi(userIdStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}
	input := models.KrsAddReq{
		Totals: int32(totals),
		Userid: int32(userId),
	}
	if _, err := kh.krsService.AddKrs(input); err != nil {
		sdk.FailOrError(c, 500, "Failed to add a krs", err)
		return
	}
	sdk.Success(c, 200, "Success to create a new KRS", input)
}

func (kh *KrsHandler) GetAllKrs(c *gin.Context) {
	result, err := kh.krsService.GetAllKrs()
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data's", err)
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
	sdk.Success(c, 200, "Success to get data", fixResult)
}

func (kh *KrsHandler) GetKrsById(c *gin.Context) {
	idStr := c.PostForm("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}

	result, err := kh.krsService.GetKrsByID(int32(id))
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
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
	sdk.Success(c, 200, "Success to get data", fixResult)
}

func (kh *KrsHandler) GetKrsByUserId(c *gin.Context) {
	idStr := c.PostForm("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}
	result, err := kh.krsService.GetKrsByIDUser(int32(id))
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data's", err)
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

	sdk.Success(c, 200, "Success to get data's", fixResult)
}

func (kh *KrsHandler) UpdateKrs(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}
	result, err := kh.krsService.GetKrsByID(int32(id))
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
		return
	}

	oldTotals := result.Totals
	oldUserId := result.Userid

	newTotalsStr := c.PostForm("sks")
	newTotals := 0
	if newTotalsStr == "" {
		newTotals = int(oldTotals)
	} else {
		newTotals, ok = strconv.Atoi(newTotalsStr)
		if ok != nil {
			sdk.FailOrError(c, 500, "Failed to convert", err)
			return
		}
	}

	newUserIdStr := c.PostForm("userID")
	newUserId := 0
	if newTotalsStr == "" {
		newUserId = int(oldUserId)
	} else {
		newUserId, ok = strconv.Atoi(newUserIdStr)
		if ok != nil {
			sdk.FailOrError(c, 500, "Failed to convert", err)
			return
		}
	}
	input := models.KrsUpdateReq{
		Totals: int32(newTotals),
		Userid: int32(newUserId),
		ID:     int32(id),
	}
	if err = kh.krsService.UpdateKrs(input); err != nil {
		sdk.FailOrError(c, 500, "Failed to update", err)
		return
	}
	sdk.Success(c, 200, "Success to update", input)
}

func (kh *KrsHandler) DeleteKrs(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}
	if err := kh.krsService.DeleteKrs(int32(id)); err != nil {
		sdk.FailOrError(c, 500, "Failed to delete", err)
		return
	}
	sdk.Success(c, 200, "Delete Success", gin.H{
		"message": "KRS Deleted",
	})
}
