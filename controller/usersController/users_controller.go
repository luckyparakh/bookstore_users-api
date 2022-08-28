package usersController

import (
	"fmt"
	"net/http"
	"strconv"

	"bookstore/src/github.com/luckyparakh/bookstore_users-api/domain"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/service"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/errors"

	"github.com/gin-gonic/gin"
)

func parseUid(uid string) (int64, *errors.RestErr) {
	user_id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("User ID should be integer.")
	}
	return user_id, nil
}
func Get(c *gin.Context) {
	userId, userErr := parseUid(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	result, err := service.UserService.GetUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result.Marshaller(c.GetHeader("X-Public") == "true"))
}

func Create(c *gin.Context) {
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
	result, err := service.UserService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result.Marshaller(c.GetHeader("X-Public") == "true"))
}

func Search(c *gin.Context) {
	status := c.Query("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("status cannot be blank"))
		return
	}
	users, err := service.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	result := make([]any, len(users))
	for _, user := range users {
		result = append(result, user.Marshaller(c.GetHeader("X-Public") == "true"))
	}
	c.JSON(http.StatusFound, result)
}

func Update(c *gin.Context) {
	user_id, userErr := parseUid(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	var inputUser domain.User
	if binderr := c.ShouldBindJSON(&inputUser); binderr != nil {
		err := errors.NewBadRequestError("Invalid JSON")
		c.JSON(err.Status, err)
		return
	}
	inputUser.Id = user_id
	partailUpdate := c.Request.Method == http.MethodPatch
	result, updateErr := service.UserService.UpdateUser(partailUpdate, inputUser)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshaller(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	user_id, userErr := parseUid(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	delErr := service.UserService.DeleteUser(user_id)
	if delErr != nil {
		c.JSON(delErr.Status, delErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "delete"})
}
