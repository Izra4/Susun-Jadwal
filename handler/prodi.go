package handler

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/sdk"
	"Susun_Jadwal/service"
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
		sdk.Fail(c, 400, "Please insert name of study program")
		return
	}

	if _, err := ph.prodiService.CreateNewProdi(context.Background(), name); err != nil {
		sdk.FailOrError(c, 500, "Failed to create new class", err)
		return
	}

	sdk.Success(c, 200, "New class created", gin.H{
		"class": name,
	})

}

func (ph *ProdiHandler) DeleteProdi(c *gin.Context) {
	idStr := c.PostForm("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}
	result, err := ph.prodiService.GetProdiById(context.Background(), int32(id))
	if err != nil {
		sdk.FailOrError(c, 500, "Class not found", err)
		return
	}

	name := result.Name
	if err := ph.prodiService.DeleteProdi(context.Background(), int32(id)); err != nil {
		sdk.FailOrError(c, 500, "Failed to delete program study", err)
		return
	}

	sdk.Success(c, 200, "Class deleted", gin.H{
		"message": "Class " + name + " deleted successfully",
	})
}

func (ph *ProdiHandler) GetProdiById(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}
	result, err := ph.prodiService.GetProdiById(context.Background(), int32(id))
	fmt.Println(result.Name + " ini debug")

	if err != nil {
		sdk.FailOrError(c, 500, "Class not found", err)
		return
	}
	sdk.Success(c, 200, "Class found", result)
}

func (ph *ProdiHandler) GetAllProdi(c *gin.Context) {
	result, err := ph.prodiService.GetAllProdi(context.TODO())
	if err != nil {
		// sdk.FailOrError(c, 500, "Failed to get data's", err)
		return
	}
	sdk.Success(c, 200, "Succes to get data", result)
}

func (ph *ProdiHandler) UpdateProdi(c *gin.Context) {
	idStr := c.Param("id")
	nameUpdate := c.PostForm("name")
	if nameUpdate == "" {
		sdk.Fail(c, 400, "Fill the name")
		return
	}
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}
	var input sqlc.UpdateProdiParams
	input.ID = int32(id)
	input.Name = nameUpdate
	if err := ph.prodiService.UpdateProdi(context.Background(), input); err != nil {
		sdk.FailOrError(c, 500, "Failed to update", err)
		return
	}
	sdk.Success(c, 200, "Name updated", gin.H{
		"new name": input.Name,
	})

}
