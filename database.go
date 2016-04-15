package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var orm *xorm.Engine

// InitDatabase sets up the database and xorm
func InitDatabase() error {
	var err error

	// connect to our database
	if os.Getenv("OPENSHIFT_MYSQL_DB_HOST") != "" {
		mysql := os.Getenv("OPENSHIFT_MYSQL_DB_USERNAME")
		mysql += ":" + os.Getenv("OPENSHIFT_MYSQL_DB_PASSWORD")
		mysql += "@tcp(" + os.Getenv("OPENSHIFT_MYSQL_DB_HOST")
		mysql += ":" + os.Getenv("OPENSHIFT_MYSQL_DB_PORT")
		mysql += ")/" + os.Getenv("OPENSHIFT_MYSQL_DB_NAME")
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
		return fmt.Errorf("ClientRequest: %s", err)
	}
	err = orm.Sync(new(Domain))
	if err != nil {
		return fmt.Errorf("Domain: %s", err)
	}
	err = orm.Sync(new(Payment))
	if err != nil {
		return fmt.Errorf("Payment: %s", err)
	}

	return nil
}
