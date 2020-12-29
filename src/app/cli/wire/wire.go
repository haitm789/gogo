//+build wireinject

package wire

import (
	"cli/adapter"

	"cli/adapter/repo"

	i "cli/repo"
	usecase "cli/usecase/mail"

	"github.com/google/wire"
)

func InitializeMailUsecase(db adapter.Database) usecase.Mail {
	wire.Build(
		usecase.NewMail,
		repo.NewMail,
		wire.Bind(new(i.Mail), new(repo.Mail)),
	)

	return usecase.Mail{}
}
