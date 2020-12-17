package v1

import (
<<<<<<< HEAD:src/app/api/infra/rest/router/v1/main.go
	"api/infra/rest/router/domain"
=======
	"api/infra/server/router/domain"
>>>>>>> b66ef7e55506fad641e3673c69db9b36f9ea0492:src/app/api/infra/rest/router/v1/v1.go
)

func Routes() domain.Group {
	r := []domain.Route{}
	r = append(r, MailRoutes()...)

	g := domain.Group{
		Prefix: "v1",
		Routes: r,
	}

	return g
}
