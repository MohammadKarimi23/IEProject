package controllers

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/moolica/IEProject/app/models"
	"github.com/revel/revel"
	"strings"
)

func init() {
	revel.OnAppStart(InitDb)
}

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	host := getParamString("db.host", "")
	port := getParamString("db.port", "3306")
	user := getParamString("db.user", "")
	pass := getParamString("db.password", "")
	dbname := getParamString("db.name", "iranfilm")
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("dbargs", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}

var InitDb func() = func() {
	connectionString := getConnectionString()
	if db, err := sql.Open("mysql", connectionString); err != nil {
		revel.ERROR.Fatal(err)
	} else {
		Dbm := &gorp.DbMap{
			Db:      db,
			Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

		defineMovieTable(Dbm)
		defineCommentTable(Dbm)
		// Defines the table for use by GORP
		// This is a function we will create soon.
		if err := Dbm.CreateTablesIfNotExists(); err != nil {
			revel.ERROR.Fatal(err)
		}

	}
}

func defineMovieTable(dbm *gorp.DbMap) {
	// set "id" as primary key and autoincrement
	t := dbm.AddTable(models.Movie{}).SetKeys(true, "id")
	t.ColMap("rate").SetMaxSize(5) //FIXME doubt if it works for non-literals
}

func defineCommentTable(dbm *gorp.DbMap) {
	t := dbm.AddTable(models.Comment{}).SetKeys(true, "id")
	t.ColMap("movie_id").SetNotNull(true)
}
