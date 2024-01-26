package adapters

import (
	"github.com/gocql/gocql"
	"github.com/karim-w/gocas/pkg/adapters/cassie"
	"github.com/karim-w/gocas/pkg/adapters/randuser"
)

type Options struct {
	ServiceName string
	Cassandra   struct {
		Host string
		Port int
		User string
		Pass string
	}
	RandomUserBaseURL string
}

type Results struct {
	Cdb      *gocql.Session
	RandUser randuser.Client
}

// SetupAdapters initializes the adapters package.
// It is called by the main package.
// Extend this function to add your own adapters.
// Pass the adapters dependencies as parameters.
// Add your adapters to the function return type
func SetupAdapters(
	opts *Options,
) (*Results, error) {
	res := &Results{}

	var err error

	// Cassandra:
	res.Cdb, err = cassie.InitCassie(
		opts.Cassandra.User,
		opts.Cassandra.Pass,
		opts.Cassandra.Port,
		opts.ServiceName,
		opts.Cassandra.Host,
	)

	if err != nil {
		return nil, err
	}

	// RandomUser:
	res.RandUser = randuser.New(opts.RandomUserBaseURL)

	return res, nil
}
