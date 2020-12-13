package v1

import (
  "api/infra/server/router/domain"
)

func Routes() domain.Group {
	r := []domain.Route{}
	r = append(r, MailRoutes()...)
	r = append(r, PostRoutes()...)

	g := domain.Group{
		Prefix: "v1",
		Routes: r,
	}

	return g
}