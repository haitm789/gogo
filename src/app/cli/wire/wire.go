//+build wireinject

package wire

import (
	"cli/adapter"

	i "cli/adapter/repo"
	r "cli/infra/repo"
	uc "cli/usecase/mail"

	"cli/infra/database"
	"github.com/google/wire"
)

func InitializeMailUsecase() uc.Mail {
	wire.Build(
		database.NewMySql,
		uc.NewMail,
		r.NewMail,
		wire.Bind(new(i.Mail), new(r.Mail)),
		wire.Bind(new(adapter.Database), new(*database.Mysql)),
	)

	return uc.Mail{}
}