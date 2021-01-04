package mail

import (
	"cli/adapter/mailer"
	"fmt"
)

func (u *Mail) Flush() {
	fmt.Println("flush mail")

	// r, _ := u.repo.Find()

	// fmt.Println(res)

	mailer := mailer.NewSmtpMailer()

	mailer.Send()
}
