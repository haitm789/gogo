//+build wireinject

//https://itnext.io/mastering-wire-f1226717bbac
package wire

import (
	"api/adapter"

	"api/adapter/repo"
	"api/adapter/rest/handler"

	irepo "api/repo"
	"api/usecase"

	"github.com/google/wire"
)

func InitializeMailHandler(db adapter.Database) handler.Mail {
	wire.Build(
		handler.NewMail,
		repo.NewMail,
		usecase.NewMail,
		wire.Bind(new(irepo.Mail), new(repo.Mail)),
	)

	return handler.Mail{}
}
