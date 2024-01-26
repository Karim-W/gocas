package usersusecase

import (
	"github.com/karim-w/gocas/pkg/domains/users"
	"github.com/karim-w/gocas/pkg/services/factory"
)

// ListUsers implements usecases.Users.
func (r *_users) ListUsers(
	ftx factory.Service,
	perPage int,
	pageState []byte,
) ([]users.Entity, []byte, error) {
	return r.usersrepo.ListUsers(ftx, perPage, pageState)
}
