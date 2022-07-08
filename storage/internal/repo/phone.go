package repo

import (
	"test_task/storage/models"
)

type PhoneStorage interface {
	GetPhone(id int64) (*models.Response, error)
}
