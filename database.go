package main

import (
	"os"

	beegoorm "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var orm beegoorm.Ormer

// InitDatabase sets up the database and beegoorm
func InitDatabase() error {
	var err error
	beegoorm.Debug = true

	// Migrate our database if needed
	beegoorm.RegisterModel(new(ClientRequest))
	beegoorm.RegisterModel(new(Domain))
	beegoorm.RegisterModel(new(Payment))

	// connect to our database
	if os.Getenv("OPENSHIFT_MYSQL_DB_HOST") != "" {
		mysql := os.Getenv("OPENSHIFT_MYSQL_DB_USERNAME")
		mysql += ":" + os.Getenv("OPENSHIFT_MYSQL_DB_PASSWORD")
		mysql += "@tcp(" + os.Getenv("OPENSHIFT_MYSQL_DB_HOST")
		mysql += ":" + os.Getenv("OPENSHIFT_MYSQL_DB_PORT")
		mysql += ")/" + os.Getenv("OPENSHIFT_MYSQL_DB_NAME")
		err = beegoorm.RegisterDataBase("default", "mysql", mysql, 30)
	} else {
		err = beegoorm.RegisterDataBase("default", "sqlite3", ":memory:", 30)
	}

	err = beegoorm.RunSyncdb("default", false, true)

	if err != nil {
		return err
	}

	orm = beegoorm.NewOrm()

	return nil
}
