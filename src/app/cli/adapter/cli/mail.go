package cli

import (
	"cli/wire"
)

func MailCommands() []Command {
	c := []Command{}
	c = append(c, Command{
		Command: "flush",
		Caption: "flush mail",
		Desc:    "flush mail",
		Func: func(args []string) {
			uc := wire.InitializeMailUsecase()
			uc.Flush()
		},
	})

	return c
}
