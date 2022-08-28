package service

import (
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/domain"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/crypto_utils"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/dateUtils"
	"bookstore/src/github.com/luckyparakh/bookstore_users-api/utils/errors"
)

var UserService userServiceInterface = &userService{}

type userService struct{}

type userServiceInterface interface {
	GetUser(int64) (*domain.User, *errors.RestErr)
	CreateUser(domain.User) (*domain.User, *errors.RestErr)
	UpdateUser(bool, domain.User) (*domain.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) ([]domain.User, *errors.RestErr)
}

func (u *userService) GetUser(id int64) (*domain.User, *errors.RestErr) {
	user := &domain.User{
		Id: id,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}
func (u *userService) CreateUser(user domain.User) (*domain.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = dateUtils.GetDbNowString()
	user.Status = domain.StatusActive
	user.Password = crypto_utils.GetMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) UpdateUser(partailUpdate bool, inputUser domain.User) (*domain.User, *errors.RestErr) {
	currentUser, getErr := u.GetUser(inputUser.Id)
	if getErr != nil {
		return nil, getErr
	}
	if partailUpdate {
		if inputUser.FirstName != "" {
			currentUser.FirstName = inputUser.FirstName
		}
		if inputUser.LastName != "" {
			currentUser.LastName = inputUser.LastName
		}
		if inputUser.Email != "" {
			currentUser.Email = inputUser.Email
		}
	} else {
		currentUser.FirstName = inputUser.FirstName
		currentUser.LastName = inputUser.LastName
		currentUser.Email = inputUser.Email
	}

	if upErr := currentUser.Update(); upErr != nil {
		return nil, upErr
	}
	return currentUser, nil
}

func (u *userService) DeleteUser(uid int64) *errors.RestErr {
	_, getErr := u.GetUser(uid)
	if getErr != nil {
		return getErr
	}
	user := domain.User{
		Id: uid,
	}
	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}

func (u *userService) SearchUser(status string) ([]domain.User, *errors.RestErr) {
	dao := domain.User{}
	return dao.GetUserByStatus(status)
}
