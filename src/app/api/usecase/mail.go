package usecase

import (
	_ "context" //ctx context.Context
)

// type Usecase interface {
// 	Do(Ctx context.Context)
// }

type mail struct {}

func NewMail() *mail {
	return &mail{}
}

func (u *mail) Get() interface{} {
	r := map[string]interface{}{
		"success": 5, 
		"message": "test",
	}
	return r;
}

// interface{} 