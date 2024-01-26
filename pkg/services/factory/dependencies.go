package factory

import (
	"github.com/gocql/gocql"
	"github.com/karim-w/gocas/pkg/infra/tracing"
)

type depenencies struct {
	trx tracing.Tracer
	cdb *gocql.Session
}

var deps *depenencies

// not thread safe
func SetUpDependencies(
	trx tracing.Tracer,
	cdb *gocql.Session,
) {
	if deps != nil {
		return
	}

	deps = &depenencies{trx, cdb}
}
