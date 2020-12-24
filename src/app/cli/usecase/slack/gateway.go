package slack

import (
	"fmt"
	// repository "api/repo"
)

type Slack struct {
	// repo repository.Mail
}

func NewSlack() Slack {
	return Slack{}
}

func (u *Slack) Post(str string) {
	fmt.Println("slack gateway: " + str)
}
