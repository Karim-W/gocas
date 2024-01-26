package services

import (
	"github.com/gocql/gocql"
	"github.com/karim-w/gocas/pkg/infra/tracing"
	"github.com/karim-w/gocas/pkg/services/factory"
)

// SetupServices initializes the services package.
// It is called by the main package.
// Extendable by adding more to the functions parameter list.
// and adding the return type to the return statement.

type Options struct {
	Trx tracing.Tracer
	Cdb *gocql.Session
}

func SetupServices(
	opts *Options,
) error {
	factory.SetUpDependencies(opts.Trx, opts.Cdb)
	return nil
}
