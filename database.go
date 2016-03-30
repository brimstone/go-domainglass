package main

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var orm *xorm.Engine

// InitDatabase sets up the database and xorm
func InitDatabase() error {
	var err error

	// connect to our database
	orm, err = xorm.NewEngine("sqlite3", "/tmp/requests.db")

	if err != nil {
		return err
	}

	err = orm.Sync(new(ClientRequest))
	if err != nil {
		return err
	}

	return nil
}
