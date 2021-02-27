package mail

type (
	Mail struct {
		Sender      string
		Recipient   string
		Subject     string
		Body        string
		Attachments []string
		Cc          []string
		Bcc         []string
	}
)

func NewMail(
	sender string,
	recipient string,
	subject string,
	body string,
	attachments []string,
	cc []string,
	bcc []string,
) Mail {
	return Mail{
		Sender:      sender,
		Recipient:   recipient,
		Subject:     subject,
		Body:        body,
		Attachments: attachments,
		Cc:          cc,
		Bcc:         bcc,
	}
}
