package repository

import (
	"Susun_Jadwal/db/sqlc"
	"context"
	"database/sql"
)

type UserRepository interface {
	CreateUser(arg sqlc.CreateUserParams) (sql.Result, error)
	DeleteUser(id int32) error
	GetAllUsers() ([]sqlc.User, error)
	GetUsersByID(id int32) (sqlc.User, error)
	UpdateUser(arg sqlc.UpdateUserParams) error
}

type userRepository struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(arg sqlc.CreateUserParams) (sql.Result, error) {
	return u.db.CreateUser(context.Background(), arg)
}

func (u *userRepository) DeleteUser(id int32) error {
	return u.db.DeleteUser(context.Background(), id)
}

func (u *userRepository) GetAllUsers() ([]sqlc.User, error) {
	return u.db.GetAllUsers(context.Background())
}

func (u *userRepository) GetUsersByID(id int32) (sqlc.User, error) {
	return u.db.GetUsersByID(context.Background(), id)
}

func (u *userRepository) UpdateUser(arg sqlc.UpdateUserParams) error {
	return u.db.UpdateUser(context.Background(), arg)
}
