package db

import (
	"database/sql"
	"fmt"

	"github.com/go-ini/ini"
	_ "github.com/lib/pq"
)

func GetClient() *sql.DB {
	cfg, err := ini.Load("db/config.ini")
	if err != nil {
		panic(err)
	}

	section := cfg.Section("database")

	host := section.Key("host").String()
	port := section.Key("port").String()
	user := section.Key("user").String()
	password := section.Key("password").String()
	dbname := section.Key("dbname").String()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
