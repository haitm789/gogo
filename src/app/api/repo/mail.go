package repo

import (
	"api/domain"
)

type (
	Mail interface {
		GetByID(id int) (domain.Mail, error)
	}
)
