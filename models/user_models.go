package models

import "time"

type UserAddReq struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Nim     string `json:"nim"`
	IDProdi int32  `json:"id_prodi"`
}

type UserUpdateReq struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Nim     string `json:"nim"`
	IDProdi int32  `json:"id_prodi"`
	ID      int32  `json:"id"`
}

type UserResult struct {
	ID        int32     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Nim       string    `json:"nim"`
	IDProdi   int32     `json:"id_prodi"`
}
