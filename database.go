package main

import (
	"os"
	"strings"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var orm *xorm.Engine

// InitDatabase sets up the database and xorm
func InitDatabase() error {
	var err error

	// connect to our database
	if os.Getenv("OPENSHIFT_MYSQL_DB_URL") != "" {
		openshiftURL := os.Getenv("OPENSHIFT_MYSQL_DB_URL")
		mysql := strings.TrimPrefix(openshiftURL, "mysql://")
		orm, err = xorm.NewEngine("mysql", mysql)
	} else {
		orm, err = xorm.NewEngine("sqlite3", ":memory:")
	}

	if err != nil {
		return err
	}

	// Migrate our database if needed
	err = orm.Sync(new(ClientRequest))
	if err != nil {
		return err
	}

	return nil
}
