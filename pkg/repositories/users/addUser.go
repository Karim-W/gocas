package usersrepository

import (
	"github.com/karim-w/gocas/pkg/domains/errs"
	"github.com/karim-w/gocas/pkg/services/factory"
	"go.uber.org/zap"
)

// AddUser implements repositories.Users.
func (r *_users) AddUser(
	ftx factory.Service,
	id string,
	email string,
	age int,
	country string,
) error {
	l := ftx.Logger()

	l.Info("Adding user")

	session := ftx.CDB()

	// create the user
	err := session.Query(
		"INSERT INTO users (user_id, email, age, country) VALUES (?, ?, ?, ?)",
		id,
		email,
		age,
		country,
	).WithContext(ftx.Context()).Exec()
	// handle failure
	if err != nil {
		l.Error("Failed to add user", zap.Error(err))

		return errs.NewBadRequest(
			err.Error(),
			ftx.TraceParent(),
		)
	}

	l.Info("Successfully added user")

	return nil
}
