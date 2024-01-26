package cassie

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

func InitCassie(
	username string,
	password string,
	port int,
	keyspace string,
	hosts string,
) (*gocql.Session, error) {
	// Connect to Cassandra cluster:
	cluster := gocql.NewCluster(
		fmt.Sprintf("%s:%d", hosts, port),
	)

	log.Println("Connecting to Cassandra cluster...")

	// Set cluster config:
	// cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10
	cluster.Keyspace = keyspace
	cluster.SslOpts = &gocql.SslOptions{
		Config: &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: true,
		},
	}
	// cluster.CQLVersion = "3.11.0"
	// cluster.DisableInitialHostLookup = true

	// Set authentication:
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}

	// Set port:
	// cluster.Port = port

	// Create session:
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println("Error creating session", err)
		return nil, err
	}

	return session, nil
}
