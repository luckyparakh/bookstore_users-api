package service

import (
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/domain"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/errors"
)

func GetUser(id int64) (*domain.User, *errors.RestErr) {
	user := &domain.User{
		Id: id,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}
func CreateUser(user domain.User) (*domain.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
