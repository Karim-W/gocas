package usersrepository

import (
	"github.com/karim-w/gocas/pkg/domains/errs"
	"github.com/karim-w/gocas/pkg/domains/users"
	"github.com/karim-w/gocas/pkg/services/factory"
	"go.uber.org/zap"
)

// ListUsers implements repositories.Users.
func (r *_users) ListUsers(
	ftx factory.Service,
	perPage int,
	pageState []byte,
) ([]users.Entity, []byte, error) {
	l := ftx.Logger()

	l.Info("Listing users")

	session := ftx.CDB()

	// get the users
	iter := session.Query(
		"SELECT user_id, email, age, country FROM users",
	).PageSize(perPage).PageState(pageState).WithContext(ftx.Context()).Iter()

	// nil check the iterator
	if iter == nil {
		l.Error("Failed to list users")

		return nil, nil, errs.NewBadRequest(
			"Failed to list users",
			ftx.TraceParent(),
		)
	}

	// check if no users were found
	if iter.NumRows() == 0 {
		l.Error("No users found")

		return nil, nil, errs.NewNotFound(
			"No users found",
			ftx.TraceParent(),
		)
	}

	var result []users.Entity
	var user users.Entity

	// parse the result
	for iter.Scan(
		&user.Id,
		&user.Email,
		&user.Age,
		&user.Country,
	) {
		result = append(result, user)
	}

	// handle failure
	if err := iter.Close(); err != nil {
		l.Error("Failed to list users", zap.Error(err))

		return nil, nil, errs.NewBadRequest(
			err.Error(),
			ftx.TraceParent(),
		)
	}

	l.Info("Successfully listed users")

	return result, iter.PageState(), nil
}
