package mail

import (
	"fmt"
)

func (u *Mail) Flush() {
	fmt.Println("flush mail")

	r, _ := u.repo.GetByID(5)
	fmt.Println("r:", r)
}
