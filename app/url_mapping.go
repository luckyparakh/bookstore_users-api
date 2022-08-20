package app

import (
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/controller/pingContoller"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/controller/usersController"
)

func mapUrl() {
	router.GET("/ping", pingContoller.Ping)
	router.GET("/users/:user_id", usersController.GetUser)
	router.GET("/users/search", usersController.SearchUser)
	router.POST("/users", usersController.CreateUser)
}
