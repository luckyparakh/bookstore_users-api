package app

import (
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/controller/pingContoller"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/controller/usersController"
)

func mapUrl() {
	router.GET("/ping", pingContoller.Ping)
	router.GET("/users/:user_id", usersController.Get)
	router.GET("/users/search", usersController.Search)
	router.POST("/users", usersController.Create)
	router.PUT("/users/:user_id", usersController.Update)
	router.PATCH("/users/:user_id", usersController.Update)
	router.DELETE("/users/:user_id", usersController.Delete)
	router.GET("/internal/users/search", usersController.Search)
	router.POST("/users/login", usersController.Login)
}
