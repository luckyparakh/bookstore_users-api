package domain

import (
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/dateUtils"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/errors"
	"fmt"
)

var usersDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User with id %d not found.", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated
	user.Email = result.Email
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("User with email %s already present.", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User with id %d already present.", user.Id))
	}
	user.DateCreated = dateUtils.GetNowString()
	usersDB[user.Id] = user
	return nil
}
