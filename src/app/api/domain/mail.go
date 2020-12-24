package domain

import "time"

type (
	Mail struct {
		id             int
		mailTemplateId int
		userId         int
		params         string
		status         MailStatus
		sendAt         time.Time
		doneAt         time.Time
	}

	MailStatus int
)

const (
	MailStatusPending MailStatus = 1
	MailStatusDone    MailStatus = 2
)

func NewMail(
	id int,
	mailTemplateId int,
	userId int,
	params string,
	status int,
	sendAt time.Time,
	doneAt time.Time,
) Mail {
	return Mail{
		id:             id,
		mailTemplateId: mailTemplateId,
		userId:         userId,
		params:         params,
		status:         MailStatus(status),
		sendAt:         sendAt,
		doneAt:         doneAt,
	}
}

func (e Mail) GetID() int {
	return e.id
}

func (e Mail) GetMailTemplateID() int {
	return e.mailTemplateId
}

func (e Mail) GetUserID() int {
	return e.userId
}

func (e Mail) GetStatus() MailStatus {
	return e.status
}

func (e Mail) GetParams() string {
	return e.params
}

func (e Mail) GetSendAt() time.Time {
	return e.sendAt
}

func (e Mail) GetDoneAt() time.Time {
	return e.doneAt
}
