package repository

import (
	"Susun_Jadwal/db/sqlc"
	"context"
	"database/sql"
)

type KrsRepository interface {
	AddKrs(arg sqlc.AddKrsParams) (sql.Result, error)
	DeleteKrs(id int32) error
	GetAllKrs() ([]sqlc.Kr, error)
	GetKrsByID(id int32) (sqlc.Kr, error)
	GetKrsByIDUser(userid int32) ([]sqlc.Kr, error)
	UpdateKrs(arg sqlc.UpdateKrsParams) error
}

type krsRepository struct {
	db *sqlc.Queries
}

func NewKrsRepository(db *sqlc.Queries) KrsRepository {
	return &krsRepository{db}
}

func (k *krsRepository) AddKrs(arg sqlc.AddKrsParams) (sql.Result, error) {
	return k.db.AddKrs(context.Background(), arg)
}

func (k *krsRepository) DeleteKrs(id int32) error {
	return k.db.DeleteKrs(context.Background(), id)
}

func (k *krsRepository) GetAllKrs() ([]sqlc.Kr, error) {
	return k.db.GetAllKrs(context.Background())
}

func (k *krsRepository) GetKrsByID(id int32) (sqlc.Kr, error) {
	return k.db.GetKrsByID(context.Background(), id)
}

func (k *krsRepository) GetKrsByIDUser(userid int32) ([]sqlc.Kr, error) {
	return k.db.GetKrsByIDUser(context.Background(), userid)
}

func (k *krsRepository) UpdateKrs(arg sqlc.UpdateKrsParams) error {
	return k.db.UpdateKrs(context.Background(), arg)
}
