package cli

import (
	"cli/adapter"
	"cli/wire"
)

func MailCommands(db adapter.Database) []Command {
	c := []Command{}
	c = append(c, Command{
		Command: "flush",
		Caption: "flush mail",
		Desc:    "flush mail",
		Func: func(args []string) {
			uc := wire.InitializeMailUsecase(db)
			uc.Flush()
		},
	})

	return c
}
