package models

import "time"

type SubjectReq struct {
	Name       string `json:"name"`
	Curriculum string `json:"curriculum"`
	Sks        int    `json:"sks"`
	IdProdi    int    `json:"id_prodi"`
	Id         int    `json:"id"`
}

type SubjectResult struct {
	Id         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	Name       string    `json:"name"`
	Curriculum string    `json:"curriculum"`
	Sks        int       `json:"sks"`
	IdProdi    int       `json:"id_prodi"`
}
