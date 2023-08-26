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

func Handler(db *sql.DB) (*ClassHandler, *ProdiHandler, *SubjectHandler) {
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
	return classHandler, prodiHand, subjectHand
}

func route(r *gin.Engine, ch *ClassHandler, ph *ProdiHandler, sh *SubjectHandler) {
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
	r.GET("get-subject-by-id", sh.GetSubjectById)
	r.POST("/add-subject", sh.CreateSubject)
	r.PATCH("/update-subject/:id", sh.UpdateSubject)
	r.DELETE("/delete-subject/:id", sh.DeleteSubject)

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

	ch, ph, sh := Handler(db)

	r := gin.Default()
	route(r, ch, ph, sh)
	r.Run()
}
