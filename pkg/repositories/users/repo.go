package usersrepository

import (
	"github.com/karim-w/gocas/pkg/adapters/randuser"
	"github.com/karim-w/gocas/pkg/repositories"
)

type _users struct {
	randomusers randuser.Client
}

func New(
	randomusers randuser.Client,
) repositories.Users {
	return &_users{randomusers}
}
