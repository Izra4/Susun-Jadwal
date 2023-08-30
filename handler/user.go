package handler

import (
	"Susun_Jadwal/models"
	"Susun_Jadwal/service"
	"Susun_Jadwal/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		util.ErrorEmptyField(c)
		return
	}

	name := c.PostForm("name")
	if name == "" {
		util.ErrorEmptyField(c)
		return
	}

	nim := c.PostForm("nim")
	if nim == "" {
		util.ErrorEmptyField(c)
		return
	}

	prodiIDStr := c.PostForm("idProdi")
	if prodiIDStr == "" {
		util.ErrorEmptyField(c)
		return
	}
	prodiID, ok := strconv.Atoi(prodiIDStr)
	if ok != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to convert", ok)
		return
	}
	input := models.UserAddReq{
		Email:   email,
		Name:    name,
		Nim:     nim,
		IDProdi: int32(prodiID),
	}
	if _, err := uh.userService.CreateUser(input); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to create user", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "Success to create data", input)
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}
	newIdProdiStr := c.PostForm("idProdi")
	newNim := c.PostForm("nim")
	newName := c.PostForm("name")
	newEmail := c.PostForm("email")

	if err := uh.userService.UpdateUser(c, id, newEmail, newName, newNim, newIdProdiStr); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to update data", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "Success to update data", gin.H{})
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	result, err := uh.userService.GetAllUsers()
	if err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to get data", err)
		return
	}

	var fixResult []models.UserResult
	for _, datas := range result {
		fixResult = append(fixResult, models.UserResult{
			ID:        datas.ID,
			CreatedAt: datas.Createdat.Time,
			UpdatedAt: datas.Updatedat.Time,
			DeletedAt: datas.Deletedat.Time,
			Email:     datas.Email,
			Name:      datas.Name,
			Nim:       datas.Nim,
			IDProdi:   datas.IDProdi,
		})
	}
	util.HttpSuccessResponse(c, 200, "Success to get data", fixResult)
}

func (uh *UserHandler) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}
	result, err := uh.userService.GetUsersByID(int32(id))
	if err != nil {
		util.HttpFailOrErrorResponse(c, 404, "Failed to get data", err)
		return
	}

	fixResult := models.UserResult{
		ID:        result.ID,
		CreatedAt: result.Createdat.Time,
		UpdatedAt: result.Updatedat.Time,
		DeletedAt: result.Deletedat.Time,
		Email:     result.Email,
		Name:      result.Name,
		Nim:       result.Nim,
		IDProdi:   result.IDProdi,
	}

	util.HttpSuccessResponse(c, 200, "Succes to get data", fixResult)
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := util.ErrorConvertStr(idStr, c)
	if ok != nil {
		return
	}

	if err := uh.userService.DeleteUser(c, int32(id)); err != nil {
		util.HttpFailOrErrorResponse(c, 500, "Failed to delete user", err)
		return
	}
	util.HttpSuccessResponse(c, 200, "Success to delete", gin.H{})
}
