package usecase

import (
	"api/domain"
)

type MailRepo interface {
	GetByID(id int) (domain.Mail, error)
}
