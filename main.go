package main

import (
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/app"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/logger"
)

func main() {
	logger.Info("about to start app")
	app.StartApplication()
}
