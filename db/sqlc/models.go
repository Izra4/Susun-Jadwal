// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"database/sql"
)

type Class struct {
	ID        int32
	Name      string
	Member    int32
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Deletedat sql.NullTime
	SubjectID int32
}

type Kr struct {
	ID        int32
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Deletedat sql.NullTime
	Totals    int32
	Userid    int32
}

type ProgramStudy struct {
	ID        int32
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Deletedat sql.NullTime
	Name      string
}

type Subject struct {
	ID         int32
	Createdat  sql.NullTime
	Updatedat  sql.NullTime
	Deletedat  sql.NullTime
	Name       string
	Curriculum string
	Sks        int32
	IDProdi    int32
}

type SubjectSchedule struct {
	ID        int32
	Day       string
	Time      string
	Room      string
	ClassID   int32
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Deletedat sql.NullTime
}

type User struct {
	ID        int32
	Createdat sql.NullTime
	Updatedat sql.NullTime
	Deletedat sql.NullTime
	Email     string
	Name      string
	Nim       string
	IDProdi   int32
}
