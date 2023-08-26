package handler

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/initializers"
	"Susun_Jadwal/repository"
	"Susun_Jadwal/service"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Handler(db *sql.DB) (*ClassHandler, *ProdiHandler, *SubjectHandler, *ScheduleHandler) {
	queries := sqlc.New(db)
	classRepo := repository.NewClassRepository(queries)
	classService := service.NewClassService(classRepo)
	classHandler := NewClassHandler(classService)

	prodiRepo := repository.NewProdiRepository(queries)
	prodiServ := service.NewProdiService(prodiRepo)
	prodiHand := NewProdiHandler(prodiServ)

	subjectRepo := repository.NewSubjectRepository(queries)
	subjectServ := service.NewSubjectService(subjectRepo)
	subjectHand := NewSubjectHandler(subjectServ)

	scheduleRepo := repository.NewScheduleRepository(queries)
	scheduleServ := service.NewScheduleService(scheduleRepo)
	scheduleHand := NewScheduleHandler(scheduleServ)

	return classHandler, prodiHand, subjectHand, scheduleHand
}

func route(r *gin.Engine, ch *ClassHandler, ph *ProdiHandler, sh *SubjectHandler, shh *ScheduleHandler) {
	r.GET("/get-classes", ch.GetAllClasses)
	r.GET("/get-class/:id", ch.GetClassById)
	r.POST("/add-class", ch.CreateClass)
	r.PATCH("/update-class/:id", ch.UpdateClass)
	r.DELETE("/delete-class/:id", ch.DeleteClass)

	r.GET("/get-class-by-id/:id", ph.GetProdiById)
	r.GET("/list-program", ph.GetAllProdi)
	r.POST("/add-program-study", ph.CreateProdi)
	r.PATCH("/update-prodi/:id", ph.UpdateProdi)
	r.DELETE("/delete-program-study", ph.DeleteProdi)

	r.GET("/get-all-subjects", sh.GetAllSubjects)
	r.GET("/get-subject-by-id/:id", sh.GetSubjectById)
	r.POST("/add-subject", sh.CreateSubject)
	r.PATCH("/update-subject/:id", sh.UpdateSubject)
	r.DELETE("/delete-subject/:id", sh.DeleteSubject)

	r.GET("/get-all-schedules", shh.GetSchedules)
	r.GET("/get-schedule-by-id/:id", shh.GetScheduleByID)
	r.POST("/add-schedule", shh.CreateSchedule)
	r.PATCH("/update-schedule", shh.UpdateSchedule)
	r.DELETE("/delete-schedule", shh.DeleteSchedule)

}

func StartEngine() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := initializers.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ch, ph, sh, shh := Handler(db)

	r := gin.Default()
	route(r, ch, ph, sh, shh)
	r.Run()
}
