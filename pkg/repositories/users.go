package repositories

import (
	"github.com/karim-w/gocas/pkg/domains/users"
	"github.com/karim-w/gocas/pkg/services/factory"
)

type Users interface {
	ListUsers(
		ftx factory.Service,
		perPage int,
		pageState []byte,
	) ([]users.Entity, []byte, error)
	AddUser(
		ftx factory.Service,
		id string,
		email string,
		age int,
		country string,
	) error
	GenerateRandomUser(
		ftx factory.Service,
	) (*users.Entity, error)
}
