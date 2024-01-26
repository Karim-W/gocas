package usecases

import (
	"github.com/karim-w/gocas/pkg/domains/users"
	"github.com/karim-w/gocas/pkg/services/factory"
)

type Users interface {
	SeedDatabase(
		ftx factory.Service,
	) error
	ListUsers(
		ftx factory.Service,
		perPage int,
		pageState []byte,
	) ([]users.Entity, []byte, error)
}
