package internal

import (
	"test_task/storage/internal/postgres"
	"test_task/storage/internal/repo"

	"github.com/jmoiron/sqlx"
)

type Storage interface {
	Phone() repo.PhoneStorage
}

type storagePg struct {
	phoneRepo repo.PhoneStorage
}

func (s *storagePg) Phone() repo.PhoneStorage {
	return s.phoneRepo
}

func NewStoragePg(db *sqlx.DB) Storage {
	return &storagePg{
		phoneRepo: postgres.NewPhone(db),
	}
}
