//+build wireinject

package wire

import (
	"api/adapter"

	"api/adapter/repo"
	"api/adapter/rest/handler"

	repositroy "api/repo"
	"api/usecase"

	"github.com/google/wire"
)

func InitializeMailHandler(db adapter.Database) handler.Mail {
	wire.Build(
		handler.NewMail,
		repo.NewMail,
		usecase.NewMail,
		wire.Bind(new(repositroy.Mail), new(repo.Mail)),
	)

	return handler.Mail{}
}
