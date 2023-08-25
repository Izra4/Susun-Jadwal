package models

import "time"

type ProdiResponse struct {
	Id        int       `json:"id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
	Deletedat time.Time `json:"deleted_at"`
	Name      string    `json:"name"`
}
