package config

import (
	"github.com/karim-w/cafe"
)

// config is the app config scoped to this package.
// It is initialized by InitOrDie.
// refer to https://github.com/karim-w/cafe for more details on how to use it.
var config *cafe.Cafe

// Cafe is my personal package for config management and can be swapped out
// for any other config management package since the getters are the only
// thing that is used in the code.

// InitRestConfigOrDie initializes the config and panics if it fails.
func InitCliConfigOrDie() {
	var err error
	config, err = cafe.New(cafe.Schema{
		"CASSANDRA_HOST":       cafe.String("CASSANDRA_HOST").Require(),
		"CASSANDRA_PORT":       cafe.Int("CASSANDRA_PORT").Default("9042"),
		"CASSANDRA_USER":       cafe.String("CASSANDRA_USER").Require(),
		"CASSANDRA_PASS":       cafe.String("CASSANDRA_PASS").Require(),
		"RANDOM_USER_BASE_URL": cafe.String("RANDOM_USER_BASE_URL").Require(),
	})
	if err != nil {
		panic(err)
	}
}

// GetServerPort returns the port the server will listen on
// It is used to initialize the server adapter
func GetServerPort() (string, error) {
	port, err := config.GetString("SERVER_PORT")
	if err != nil {
		return "", err
	}
	return ":" + port, nil
}

// TODO: Add your config getters here
// func GetFoo() (string,error) {
// 	return Config.GetString("foo")
// }

// GetCassandraConfig returns the cassandra config
func GetCassandraConfig() (string, int, string, string, error) {
	host, err := config.GetString("CASSANDRA_HOST")
	if err != nil {
		return "", 0, "", "", err
	}
	port, err := config.GetInt("CASSANDRA_PORT")
	if err != nil {
		return "", 0, "", "", err
	}
	user, err := config.GetString("CASSANDRA_USER")
	if err != nil {
		return "", 0, "", "", err
	}
	pass, err := config.GetString("CASSANDRA_PASS")
	if err != nil {
		return "", 0, "", "", err
	}
	return host, port, user, pass, nil
}

// GetRandomUserBaseURL returns the random user base url
func GetRandomUserBaseURL() (string, error) {
	url, err := config.GetString("RANDOM_USER_BASE_URL")
	if err != nil {
		return "", err
	}
	return url, nil
}
