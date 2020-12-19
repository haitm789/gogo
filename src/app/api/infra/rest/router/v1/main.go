package v1

import (
	"api/infra/rest/router/domain"
)

func Routes() domain.Group {
	r := []domain.Route{}
	r = append(r, mail()...)
	r = append(r, user()...)

	g := domain.Group{
		Prefix: "v1",
		Routes: r,
	}

	return g
}
