package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/karim-w/gocas/cmd/apps"
	"github.com/karim-w/gocas/internal/config"
	"github.com/karim-w/gocas/internal/constants"
	"github.com/karim-w/gocas/pkg/adapters"
	"github.com/karim-w/gocas/pkg/domains/errs"
	"github.com/karim-w/gocas/pkg/infra"
	usersrepository "github.com/karim-w/gocas/pkg/repositories/users"
	"github.com/karim-w/gocas/pkg/services"
	"github.com/karim-w/gocas/pkg/services/factory"
	usersusecase "github.com/karim-w/gocas/pkg/usecases/users"

	"go.uber.org/zap"
)

type cliApp struct {
	// TODO: add fields to close resources here
}

func CliApp() apps.Application {
	return &cliApp{}
}

func (a *cliApp) Setup() error {
	config.InitCliConfigOrDie()

	host, port, user, pass, err := config.GetCassandraConfig()
	if err != nil {
		return err
	}

	randomUserBaseURL, err := config.GetRandomUserBaseURL()
	if err != nil {
		return err
	}

	// ========= SetupAdapters =========
	adpts, err := adapters.SetupAdapters(&adapters.Options{
		ServiceName: constants.SERVICE_NAME,
		Cassandra: struct {
			Host string
			Port int
			User string
			Pass string
		}{
			Host: host,
			Port: port,
			User: user,
			Pass: pass,
		},
		RandomUserBaseURL: randomUserBaseURL,
	})
	if err != nil {
		return err
	}

	// ========= SetupInfra =========
	infras, err := infra.SetupInfra(&infra.Options{})
	if err != nil {
		return err
	}

	// ========= Setup Repositories =========
	ur := usersrepository.New(adpts.RandUser)
	// ========= Setup Services =========
	err = services.SetupServices(&services.Options{
		Trx: infras.Trx,
		Cdb: adpts.Cdb,
	})
	// ========= Setup Usecases =========
	uuc := usersusecase.New(ur)

	// ========= Setup Apps =========
	for i := 0; i < 100; i++ {
		err = uuc.SeedDatabase(factory.NewFactory(context.TODO()))
		if err != nil {
			return err
		}
	}

	users, _, err := uuc.ListUsers(factory.NewFactory(context.TODO()), 100, nil)
	if err != nil {
		return err
	}

	if len(users) == 0 {
		zap.L().Error("No users found")
		return errs.NewNotFound(
			"no users",
			"",
		)
	}

	if len(users) != 100 {
		return errors.New(fmt.Sprintf("Expected 100 users, got %d", len(users)))
	}
	return nil
}

func (a *cliApp) Close() {
	zap.L().Info("Closing cli app")
	zap.L().Sync()
	// TODO: close resources here
}
