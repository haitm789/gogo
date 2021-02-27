package repo

import (
	"cli/domain"
)

type (
	Mail interface {
		Find() ([]domain.Mail, error)
	}
)
