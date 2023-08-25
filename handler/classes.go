package handler

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/sdk"
	"Susun_Jadwal/service"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassHandler struct {
	classService service.ClassService
}

func NewClassHandler(classService service.ClassService) *ClassHandler {
	return &ClassHandler{classService}
}

func (ch *ClassHandler) CreateClass(c *gin.Context) {
	var input sqlc.AddNewClassParams

	room := c.PostForm("name")

	memberStr := c.PostForm("member")
	member, err := strconv.Atoi(memberStr)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to convert",
		})
		return
	}

	majorIdStr := c.PostForm("majorId")
	majorId, err := strconv.Atoi(majorIdStr)
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to convert", err)
		return
	}

	//fill data
	input.Name = room
	input.Member = int32(member)
	input.MajorID = int32(majorId)
	_, err = ch.classService.CreateClass(context.Background(), input)
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to create a new class", err)
		return
	}

	c.JSON(http.StatusOK, input)
}

func (ch *ClassHandler) DeleteClass(c *gin.Context) {
	idStr := c.PostForm("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}

	name, err := ch.classService.GetClassNameById(context.Background(), int32(id))
	if err != nil {
		sdk.FailOrError(c, 500, "Class not found", err)
		return
	}

	err = ch.classService.DeleteClass(context.Background(), int32(id))
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to delete class", err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Class " + name + " deleted",
	})
}

func (ch *ClassHandler) GetAllClass(c *gin.Context) {
	result, err := ch.classService.GetListClass(context.Background())
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data's", err)
		return
	}
	sdk.Success(c, 200, "Succes to get data", result)
}
