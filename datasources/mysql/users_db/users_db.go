package users_db

import (
	"database/sql"
	"fmt"
	"log"

	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/config"

	_ "github.com/go-sql-driver/mysql"
)

var Client *sql.DB

func init() {
	config, configerr := config.LoadConfig()
	if configerr != nil {
		panic(configerr.Message)
	}
	connName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB.DbUser, config.DB.DbPassword, config.DB.DbAddress, config.DB.DbPort, config.DB.DbName)
	var err error
	Client, err = sql.Open("mysql", connName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("connected to db")
}
