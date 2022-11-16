package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang-freelance/helper"
	"time"
)

func NewDB() *sql.DB {
	var (
		DbHost = EnvVariable("DB_HOST")
		DbPort = EnvVariable("DB_PORT")
		DbUser = EnvVariable("DB_USER")
		DbPass = EnvVariable("DB_PASS")
		DbName = EnvVariable("DB_NAME")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DbUser, DbPass, DbHost, DbPort, DbName)
	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
