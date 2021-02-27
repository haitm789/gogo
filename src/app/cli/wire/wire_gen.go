// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire

import (
	"cli/adapter"
	"cli/infra/repo"
	"cli/usecase/mail"
)

// Injectors from wire.go:

func InitializeMailUsecase(db adapter.Database) mail.Mail {
	repoMail := repo.NewMail(db)
	mailMail := mail.NewMail(repoMail)
	return mailMail
}
