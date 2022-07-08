package postgres

import (
	"test_task/storage/internal/repo"
	"test_task/storage/models"

	"github.com/jmoiron/sqlx"
)

type phoneRepo struct {
	db *sqlx.DB
}

func NewPhone(db *sqlx.DB) repo.PhoneStorage {
	return &phoneRepo{
		db: db,
	}
}

func (p *phoneRepo) GetPhone(id int64) (*models.Response, error) {
	var (
		phone       models.Response
		phoneNumber string
	)
	phone.Phone = make(map[int64]string)
	query := `SELECT phone from phone where id = $1`

	err := p.db.QueryRow(query, id).Scan(&phoneNumber)
	if err != nil {
		return nil, err
	}
	phone.Phone[id] = phoneNumber

	return &phone, nil

}
