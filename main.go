package main

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/handler"
	"Susun_Jadwal/initializers"
	"Susun_Jadwal/repository"
	"Susun_Jadwal/service"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitHandler(db *sql.DB) (*handler.ClassHandler, *handler.ProdiHandler, *handler.SubjectHandler,
	*handler.ScheduleHandler, *handler.UserHandler, *handler.KrsHandler) {
	queries := sqlc.New(db)
	classRepo := repository.NewClassRepository(queries)
	classService := service.NewClassService(classRepo)
	classHandler := handler.NewClassHandler(classService)

	prodiRepo := repository.NewProdiRepository(queries)
	prodiServ := service.NewProdiService(prodiRepo)
	prodiHand := handler.NewProdiHandler(prodiServ)

	subjectRepo := repository.NewSubjectRepository(queries)
	subjectServ := service.NewSubjectService(subjectRepo)
	subjectHand := handler.NewSubjectHandler(subjectServ)

	scheduleRepo := repository.NewScheduleRepository(queries)
	scheduleServ := service.NewScheduleService(scheduleRepo)
	scheduleHand := handler.NewScheduleHandler(scheduleServ)

	userRepo := repository.NewUserRepository(queries)
	userServ := service.NewUserService(userRepo)
	userHand := handler.NewUserHandler(userServ)

	krsRepo := repository.NewKrsRepository(queries)
	krsServ := service.NewKrsService(krsRepo)
	krsHand := handler.NewKrsHandler(krsServ)
	return classHandler, prodiHand, subjectHand, scheduleHand, userHand, krsHand
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := initializers.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()
	handler.StartEngine(r, db)
	r.Run()
}
