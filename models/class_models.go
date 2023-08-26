package models

import "time"

type ClassAddReq struct {
	Name      string `json:"name"`
	Member    int    `json:"member"`
	SubjectId int    `json:"subject_id"`
}

type ClassUpdateReq struct {
	Name      string `json:"name"`
	Member    int32  `json:"member"`
	SubjectID int32  `json:"subject_id"`
	ID        int32  `json:"id"`
}

type ClassResult struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Name      string    `json:"name"`
	Member    int       `json:"member"`
	SubjectId int       `json:"subject_id"`
}
