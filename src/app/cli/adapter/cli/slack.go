package cli

import (
	"cli/usecase/slack"
	"fmt"
)

func SlackCommands() []Command {
	c := []Command{}
	c = append(c, Command{
		Command: "slack2",
		Caption: "slack intro2",
		Desc:    "slack description2",
		Func: func(args []string) {
			fmt.Println("slack2")
		},
	})

	c = append(c, Command{
		Command: "slack3",
		Caption: "slack intro3",
		Desc:    "slack description3",
		Func: func(args []string) {
			bot := slack.NewSlack()
			bot.Post("from slack bot3")
		},
	})

	return c
}
