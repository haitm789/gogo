package v1

import (
	"api/adapter"
	"api/infra/rest/router/domain"
)

func Routes(db adapter.Database) domain.Group {
	r := []domain.Route{}
	r = append(r, mail(db)...)
	r = append(r, user()...)

	g := domain.Group{
		Prefix: "v1",
		Routes: r,
	}

	return g
}
