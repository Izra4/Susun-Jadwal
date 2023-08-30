package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/models"
	"Susun_Jadwal/repository"
	"Susun_Jadwal/util"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(arg models.UserAddReq) (sql.Result, error)
	DeleteUser(c *gin.Context, id int32) error
	GetAllUsers() ([]sqlc.User, error)
	GetUsersByID(id int32) (sqlc.User, error)
	UpdateUser(c *gin.Context, id int, newEmail string, newName string, newNim string, newIdProdiStr string) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (u *userService) CreateUser(arg models.UserAddReq) (sql.Result, error) {
	input := sqlc.CreateUserParams{
		Email:   arg.Email,
		Name:    arg.Name,
		Nim:     arg.Nim,
		IDProdi: arg.IDProdi,
	}
	return u.userRepository.CreateUser(input)
}

func (u *userService) DeleteUser(c *gin.Context, id int32) error {
	_, err := u.userRepository.GetUsersByID(int32(id))
	if err != nil {
		util.HttpFailOrErrorResponse(c, 404, "Failed to get data", err)
		return err
	}
	return u.userRepository.DeleteUser(id)
}

func (u *userService) GetAllUsers() ([]sqlc.User, error) {
	return u.userRepository.GetAllUsers()
}

func (u *userService) GetUsersByID(id int32) (sqlc.User, error) {
	return u.userRepository.GetUsersByID(id)
}

func (u *userService) UpdateUser(c *gin.Context, id int, newEmail string, newName string, newNim string, newIdProdiStr string) error {
	result, err := u.userRepository.GetUsersByID(int32(id))
	if err != nil {
		util.HttpFailOrErrorResponse(c, 404, "Failed to get data", err)
		return err
	}
	oldEmail := result.Email
	oldName := result.Name
	oldNim := result.Nim
	oldIdProdi := result.IDProdi
	ok := err

	if newEmail == "" {
		newEmail = oldEmail
	}
	if newName == "" {
		newName = oldName
	}
	if newNim == "" {
		newNim = oldNim
	}
	newIdProdi := 0
	if newIdProdiStr == "" {
		newIdProdi = int(oldIdProdi)
	} else {
		newIdProdi, ok = util.ErrorConvertStr(newIdProdiStr, c)
		if ok != nil {
			return err
		}
	}
	input := sqlc.UpdateUserParams{
		Email:   newEmail,
		Name:    newName,
		Nim:     newNim,
		IDProdi: int32(newIdProdi),
		ID:      int32(id),
	}
	return u.userRepository.UpdateUser(input)
}
