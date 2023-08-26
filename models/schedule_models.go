package models

import "time"

type ScheduleAddReq struct {
	Day     string `json:"day"`
	Time    string `json:"time"`
	Room    string `json:"room"`
	ClassID int32  `json:"class_id"`
}

type ScheduleUpdateReq struct {
	Day     string `json:"day"`
	Time    string `json:"time"`
	Room    string `json:"room"`
	ClassID int32  `json:"class_id"`
	ID      int32  `json:"id"`
}

type ScheduleResult struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Day       string    `json:"day"`
	Room      string    `json:"room"`
	Time      string    `json:"time"`
	ClassId   int       `json:"class_id"`
}
