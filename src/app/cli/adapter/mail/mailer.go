package mail

type Mailer interface {
	Flush(msg Mail) error
}
