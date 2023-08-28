package service

import (
	"Susun_Jadwal/db/sqlc"
	"Susun_Jadwal/models"
	"Susun_Jadwal/repository"
	"Susun_Jadwal/util"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type KrsService interface {
	AddKrs(arg models.KrsAddReq) (sql.Result, error)
	DeleteKrs(id int32) error
	GetAllKrs() ([]sqlc.Kr, error)
	GetKrsByID(id int32) (sqlc.Kr, error)
	GetKrsByIDUser(userid int32) ([]sqlc.Kr, error)
	UpdateKrs(cgx *gin.Context, newTotalsStr string, newUserIdStr string, id int32) error
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

func (k *krsService) UpdateKrs(cgx *gin.Context, newTotalsStr string, newUserIdStr string, id int32) error {
	result, err := k.krsRepository.GetKrsByID(id)
	if err != nil {
		return nil
	}

	oldTotals := result.Totals
	oldUserId := result.Userid
	ok := err
	newTotals := 0
	if newTotalsStr == "" {
		newTotals = int(oldTotals)
	} else {
		newTotals, ok = util.ErrorConvertStr(newTotalsStr, cgx)
		if ok != nil {
			return nil
		}
	}

	newUserId := 0
	if newTotalsStr == "" {
		newUserId = int(oldUserId)
	} else {
		newUserId, ok = util.ErrorConvertStr(newUserIdStr, cgx)
		if ok != nil {
			return nil
		}
	}
	input := sqlc.UpdateKrsParams{
		Totals: int32(newTotals),
		Userid: int32(newUserId),
		ID:     id,
	}
	return k.krsRepository.UpdateKrs(input)
}
