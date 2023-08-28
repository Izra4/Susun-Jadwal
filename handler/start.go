package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func route(r *gin.Engine, ch *ClassHandler, ph *ProdiHandler, sh *SubjectHandler,
	shh *ScheduleHandler, uh *UserHandler, kh *KrsHandler) {
	r.GET("/get-classes", ch.GetAllClasses)
	r.GET("/get-class/:id", ch.GetClassById)
	r.POST("/add-class", ch.CreateClass)
	r.PATCH("/update-class/:id", ch.UpdateClass)
	r.DELETE("/delete-class/:id", ch.DeleteClass)

	r.GET("/get-class-by-id/:id", ph.GetProdiById)
	r.GET("/list-program", ph.GetAllProdi)
	r.POST("/add-program-study", ph.CreateProdi)
	r.PATCH("/update-prodi/:id", ph.UpdateProdi)
	r.DELETE("/delete-program-study/:id", ph.DeleteProdi)

	r.GET("/get-all-subjects", sh.GetAllSubjects)
	r.GET("/get-subject-by-id/:id", sh.GetSubjectById)
	r.POST("/add-subject", sh.CreateSubject)
	r.PATCH("/update-subject/:id", sh.UpdateSubject)
	r.DELETE("/delete-subject/:id", sh.DeleteSubject)

	r.GET("/get-all-schedules", shh.GetSchedules)
	r.GET("/get-schedule-by-id/:id", shh.GetScheduleByID)
	r.POST("/add-schedule", shh.CreateSchedule)
	r.PATCH("/update-schedule/:id", shh.UpdateSchedule)
	r.DELETE("/delete-schedule/:id", shh.DeleteSchedule)

	r.GET("/get-users", uh.GetAllUsers)
	r.GET("/get-user/:id", uh.GetUserById)
	r.POST("/add-user", uh.CreateUser)
	r.PATCH("/update-user/:id", uh.UpdateUser)
	r.DELETE("/delete-user/:id", uh.DeleteUser)

	r.GET("/get-krs", kh.GetAllKrs)
	r.GET("/get-krs-by-id/:id", kh.GetKrsById)
	r.GET("/get-krs-by-user-id", kh.GetKrsByUserId)
	r.POST("/add-krs", kh.AddKrs)
	r.PATCH("/update-krs", kh.UpdateKrs)
	r.DELETE("/delete-krs", kh.DeleteKrs)
}

func StartEngine(r *gin.Engine, db *sql.DB) {
	ch, ph, sh, shh, uh, kh := InitHandler(db)
	route(r, ch, ph, sh, shh, uh, kh)
}
