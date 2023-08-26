package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/models"
	"Susun_Jadwal/repository"
	"database/sql"
)

type KrsService interface {
	AddKrs(arg models.KrsAddReq) (sql.Result, error)
	DeleteKrs(id int32) error
	GetAllKrs() ([]sqlc.Kr, error)
	GetKrsByID(id int32) (sqlc.Kr, error)
	GetKrsByIDUser(userid int32) ([]sqlc.Kr, error)
	UpdateKrs(arg models.KrsUpdateReq) error
}

type krsService struct {
	krsRepository repository.KrsRepository
}

func NewKrsService(krsRepository repository.KrsRepository) KrsService {
	return &krsService{krsRepository}
}

func (k *krsService) AddKrs(arg models.KrsAddReq) (sql.Result, error) {
	input := sqlc.AddKrsParams{
		Totals: arg.Totals,
		Userid: arg.Userid,
	}
	return k.krsRepository.AddKrs(input)
}

func (k *krsService) DeleteKrs(id int32) error {
	return k.krsRepository.DeleteKrs(id)
}

func (k *krsService) GetAllKrs() ([]sqlc.Kr, error) {
	return k.krsRepository.GetAllKrs()
}

func (k *krsService) GetKrsByID(id int32) (sqlc.Kr, error) {
	return k.krsRepository.GetKrsByID(id)
}

func (k *krsService) GetKrsByIDUser(userid int32) ([]sqlc.Kr, error) {
	return k.krsRepository.GetKrsByIDUser(userid)
}

func (k *krsService) UpdateKrs(arg models.KrsUpdateReq) error {
	input := sqlc.UpdateKrsParams{
		Totals: arg.Totals,
		Userid: arg.Userid,
		ID:     arg.ID,
	}
	return k.krsRepository.UpdateKrs(input)
}
