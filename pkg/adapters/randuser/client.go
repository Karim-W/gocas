package randuser

import (
	"github.com/karim-w/gocas/pkg/domains/errs"
	"github.com/karim-w/gocas/pkg/services/factory"
	"go.uber.org/zap"
)

type Client interface {
	RandomUser(
		ftx factory.Service,
	) (*User, error)
}

type _user struct {
	baseUrl string
}

func New(baseUrl string) Client {
	return &_user{baseUrl}
}

func (u *_user) RandomUser(
	ftx factory.Service,
) (*User, error) {
	l := ftx.Logger()

	l.Info("Getting random user")

	// call the random user api
	res := ftx.HttpClient(u.baseUrl + "/api").Get()

	// handle failure
	if !res.IsSuccess() {
		body := string(res.GetBody())
		code := res.GetStatusCode()

		l.Error("Failed to get random user", zap.String("body", body), zap.Int("code", code))

		return nil, errs.New(
			code,
			"Failed to get random user",
			"Failed to get random user with "+body,
			ftx.TraceParent(),
		)
	}

	l.Info("Successfully got random user")

	var result Users

	// parse the result
	err := res.SetResult(&result)
	if err != nil {
		l.Error("Failed to parse random user", zap.Error(err))

		return nil, errs.NewBadRequest(
			err.Error(),
			ftx.TraceParent(),
		)
	}

	//	check if there is at least one result
	if len(result.Results) == 0 {
		l.Error("No results found")

		return nil, errs.NewNotFound(
			"No results found",
			ftx.TraceParent(),
		)
	}

	return &result.Results[0], nil
}
