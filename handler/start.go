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

func Handler(db *sql.DB) (*ClassHandler, *ProdiHandler) {
	queries := sqlc.New(db)
	classRepo := repository.NewClassRepository(queries)
	classService := service.NewClassService(classRepo)
	classHandler := NewClassHandler(classService)

	prodiRepo := repository.NewProdiRepository(queries)
	prodiServ := service.NewProdiService(prodiRepo)
	prodiHand := NewProdiHandler(prodiServ)
	return classHandler, prodiHand
}

func route(r *gin.Engine, ch *ClassHandler, ph *ProdiHandler) {
	r.POST("/add-class", ch.CreateClass)
	r.GET("/get-class", ch.GetAllClass)
	r.DELETE("/delete-class", ch.DeleteClass)

	r.POST("/add-program-study", ph.CreateProdi)
	r.GET("/get-class-by-id/:id", ph.GetProdiById)
	r.GET("/list-program", ph.GetAllProdi)
	r.PATCH("/update-prodi/:id", ph.UpdateProdi)
	r.DELETE("/delete-program-study", ph.DeleteProdi)
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

	ch, ph := Handler(db)

	r := gin.Default()
	route(r, ch, ph)
	r.Run()
}
