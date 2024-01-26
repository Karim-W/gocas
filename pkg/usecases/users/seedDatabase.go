package usersusecase

import "github.com/karim-w/gocas/pkg/services/factory"

// SeedDatabase implements usecases.Users.
func (r *_users) SeedDatabase(ftx factory.Service) error {
	user, err := r.usersrepo.GenerateRandomUser(ftx)
	if err != nil {
		return err
	}

	return r.usersrepo.AddUser(ftx, user.Id, user.Email, user.Age, user.Country)
}
