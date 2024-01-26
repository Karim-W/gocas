package usersrepository

import (
	"github.com/karim-w/gocas/pkg/domains/users"
	"github.com/karim-w/gocas/pkg/services/factory"
	"go.uber.org/zap"
)

// GenerateRandomUser implements repositories.Users.
func (r *_users) GenerateRandomUser(ftx factory.Service) (*users.Entity, error) {
	l := ftx.Logger()

	l.Info("Generating random user")

	user, err := r.randomusers.RandomUser(ftx)
	if err != nil {
		l.Error("Failed to generate random user", zap.Error(err))

		return nil, err
	}

	l.Info("Successfully generated random user")

	// map the user
	return &users.Entity{
		Id:      user.Login.UUID,
		Email:   user.Email,
		Age:     int(user.Dob.Age),
		Country: user.Location.Country,
	}, nil
}
