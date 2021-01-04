package repo

import (
	"cli/domain"
)

type (
	Mail interface {
		// GetByID(id int) (domain.Mail, error)
		Find() ([]domain.Mail, error)
	}
)
