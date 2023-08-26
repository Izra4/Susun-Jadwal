package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/models"
	"Susun_Jadwal/repository"
	"database/sql"
)

type UserService interface {
	CreateUser(arg models.UserAddReq) (sql.Result, error)
	DeleteUser(id int32) error
	GetAllUsers() ([]sqlc.User, error)
	GetUsersByID(id int32) (sqlc.User, error)
	UpdateUser(arg models.UserUpdateReq) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (u userService) CreateUser(arg models.UserAddReq) (sql.Result, error) {
	input := sqlc.CreateUserParams{
		Email:   arg.Email,
		Name:    arg.Name,
		Nim:     arg.Nim,
		IDProdi: arg.IDProdi,
	}
	return u.userRepository.CreateUser(input)
}

func (u userService) DeleteUser(id int32) error {
	return u.userRepository.DeleteUser(id)
}

func (u userService) GetAllUsers() ([]sqlc.User, error) {
	return u.userRepository.GetAllUsers()
}

func (u userService) GetUsersByID(id int32) (sqlc.User, error) {
	return u.userRepository.GetUsersByID(id)
}

func (u userService) UpdateUser(arg models.UserUpdateReq) error {
	input := sqlc.UpdateUserParams{
		Email:   arg.Email,
		Name:    arg.Name,
		Nim:     arg.Nim,
		IDProdi: arg.IDProdi,
		ID:      arg.ID,
	}
	return u.userRepository.UpdateUser(input)
}
