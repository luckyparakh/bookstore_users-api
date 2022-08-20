package usersController

import (
	"fmt"
	"net/http"

	"bookstore/src/github.com/luckyparakh/bookstore_users-api/domain"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/service"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/errors"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func CreateUser(c *gin.Context) {
	fmt.Println("create user")
	var user domain.User
	// same as below
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// fmt.Printf("bytes %v\n", string(bytes))
	// if err != nil {
	// 	//TODO: raise error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//TODO: raise json error
	// 	fmt.Println(err.Error())
	// 	return
	// }
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, err := service.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
