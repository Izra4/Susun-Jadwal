package handler

import (
	"Susun_Jadwal/models"
	"Susun_Jadwal/sdk"
	"Susun_Jadwal/service"
	"fmt"
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
		sdk.FailEmptyField(c)
		return
	}

	name := c.PostForm("name")
	if name == "" {
		sdk.FailEmptyField(c)
		return
	}

	nim := c.PostForm("nim")
	if nim == "" {
		sdk.FailEmptyField(c)
		return
	}

	prodiIDStr := c.PostForm("idProdi")
	if prodiIDStr == "" {
		sdk.FailEmptyField(c)
		return
	}
	prodiID, ok := strconv.Atoi(prodiIDStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}
	input := models.UserAddReq{
		Email:   email,
		Name:    name,
		Nim:     nim,
		IDProdi: int32(prodiID),
	}
	if _, err := uh.userService.CreateUser(input); err != nil {
		sdk.FailOrError(c, 500, "Failed to create user", err)
		return
	}
	sdk.Success(c, 200, "Success to create data", input)
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Faileed to convert", ok)
		return
	}
	result, err := uh.userService.GetUsersByID(int32(id))
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
		return
	}
	oldEmail := result.Email
	oldName := result.Name
	oldNim := result.Nim
	oldIdProdi := result.IDProdi

	newEmail := c.PostForm("email")
	if newEmail == "" {
		newEmail = oldEmail
	}
	newName := c.PostForm("name")
	if newName == "" {
		newName = oldName
	}
	newNim := c.PostForm("nim")
	if newNim == "" {
		newNim = oldNim
	}
	newIdProdiStr := c.PostForm("idProdi")
	newIdProdi := 0
	if newIdProdiStr == "" {
		newIdProdi = int(oldIdProdi)
	} else {
		newIdProdi, ok = strconv.Atoi(newIdProdiStr)
		if ok != nil {
			sdk.FailOrError(c, 500, "Failed to convert", err)
			return
		}
	}
	input := models.UserUpdateReq{
		Email:   newEmail,
		Name:    newName,
		Nim:     newNim,
		IDProdi: int32(newIdProdi),
		ID:      int32(id),
	}
	if err = uh.userService.UpdateUser(input); err != nil {
		sdk.FailOrError(c, 500, "Failed to update data", err)
		return
	}
	sdk.Success(c, 200, "Success to update data", input)
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	result, err := uh.userService.GetAllUsers()
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
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
	sdk.Success(c, 200, "Success to get data", fixResult)
}

func (uh *UserHandler) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}
	result, err := uh.userService.GetUsersByID(int32(id))
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
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

	sdk.Success(c, 200, "Succes to get data", fixResult)
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, ok := strconv.Atoi(idStr)
	if ok != nil {
		sdk.FailOrError(c, 500, "Failed to convert", ok)
		return
	}
	result, err := uh.userService.GetUsersByID(int32(id))
	if err != nil {
		sdk.FailOrError(c, 500, "Failed to get data", err)
		return
	}
	email := result.Email
	if err = uh.userService.DeleteUser(int32(id)); err != nil {
		sdk.FailOrError(c, 500, "Failed to delete user", err)
		return
	}
	sdk.Success(c, 200, "Success to delete", gin.H{
		"message": fmt.Sprintf("User %s successfully deleted", email),
	})
}
