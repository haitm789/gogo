//+build wireinject

package wire

import (
	"cli/adapter"

	i "cli/adapter/repo"
	r "cli/infra/repo"
	uc "cli/usecase/mail"

	"github.com/google/wire"
)

func InitializeMailUsecase(db adapter.Database) uc.Mail {
	wire.Build(
		uc.NewMail,
		r.NewMail,
		wire.Bind(new(i.Mail), new(r.Mail)),
	)

	return uc.Mail{}
}
