package handler

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/models"
	"Susun_Jadwal/service"
	"Susun_Jadwal/util"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProdiHandler struct {
	prodiService service.ProdiService
}

func NewProdiHandler(prodiService service.ProdiService) *ProdiHandler {
	return &ProdiHandler{prodiService}
}

func (ph *ProdiHandler) CreateProdi(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		util.ErrorEmptyField(c)
		return
	}

	if _, err := ph.prodiService.CreateNewProdi(context.Background(), name); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to create new class", err)
		return
	}

	util.HttpSuccessResponse(c, 200, "New class created", gin.H{
		"class": name,
	})

}

func (ph *ProdiHandler) DeleteProdi(c *gin.Context) {
	idStr := c.PostForm("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return
	}
	result, err := ph.prodiService.GetProdiById(context.Background(), int32(id))
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Class not found", err)
		return
	}

	name := result.Name
	if err := ph.prodiService.DeleteProdi(context.Background(), int32(id)); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to delete program study", err)
		return
	}

	util.HttpSuccessResponse(c, 200, "Class deleted", gin.H{
		"message": "Class " + name + " deleted successfully",
	})
}

func (ph *ProdiHandler) GetProdiById(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return
	}
	result, err := ph.prodiService.GetProdiById(context.Background(), int32(id))
	fmt.Println(result.Name + " ini debug")

	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Class not found", err)
		return
	}

	fixResult := models.ProdiResponse{
		Id:        int(result.ID),
		Createdat: result.Createdat.Time,
		Updatedat: result.Updatedat.Time,
		Deletedat: result.Deletedat.Time,
		Name:      result.Name,
	}

	util.HttpSuccessResponse(c, 200, "Class found", fixResult)
}

func (ph *ProdiHandler) GetAllProdi(c *gin.Context) {
	result, err := ph.prodiService.GetAllProdi(context.TODO())
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to get data's", err)
		return
	}
	var fixResult []models.ProdiResponse
	for _, prodi := range result {
		fixResult = append(fixResult, models.ProdiResponse{
			Id:        int(prodi.ID),
			Createdat: prodi.Createdat.Time,
			Updatedat: prodi.Updatedat.Time,
			Deletedat: prodi.Deletedat.Time,
			Name:      prodi.Name,
		})
	}

	util.HttpSuccessResponse(c, 200, "Data Loaded", fixResult)
}

func (ph *ProdiHandler) UpdateProdi(c *gin.Context) {
	idStr := c.Param("id")
	nameUpdate := c.PostForm("name")
	if nameUpdate == "" {
		util.ErrorEmptyField(c)
		return
	}
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return
	}
	var input sqlc.UpdateProdiParams
	input.ID = int32(id)
	input.Name = nameUpdate
	if err := ph.prodiService.UpdateProdi(context.Background(), input); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to update", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "Name updated", gin.H{
		"new name": input.Name,
	})

}
