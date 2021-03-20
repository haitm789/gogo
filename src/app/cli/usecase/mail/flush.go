package mail

import (
	mailer "cli/infra/mailer"
	"cli/lib"

	"fmt"
)

func (u *Mail) Flush() {
	fmt.Println("uc: flush mail")

	r, _ := u.repo.Find()

	fmt.Println(r)

	agent := lib.NewMailer(
		mailer.NewSmtpMailer(),
	)

	msg := NewMailMessage(
		"j3ns789@gmail.com",
		"haitm789@gmail.com",
		"mail flush",
		"this is a mail flush.\n\npandog.",
		[]string{},
		[]string{},
		[]string{},
	)

	agent.Flush(msg)
}
