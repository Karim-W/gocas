package usersusecase

import (
	"github.com/karim-w/gocas/pkg/repositories"
	"github.com/karim-w/gocas/pkg/usecases"
)

type _users struct {
	usersrepo repositories.Users
}

func New(
	usersrepo repositories.Users,
) usecases.Users {
	return &_users{usersrepo}
}
